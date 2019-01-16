package fetch

import (
	"context"
	"log"

	"github.com/tjcain/ukpolice"
)

// Repository provides access to storage.
type Repository interface {
	// StoreSearch stores a search
	StoreSearch(ukpolice.Search) error
	//CreateTables creates tables, it will delete any tables that exist.
	CreateTables()
}

// Service provides fetch options.
type Service interface {
	UpdateData() error
	CreateTables()
}

type service struct {
	r Repository
	c *ukpolice.Client
}

// NewService creates a fetch service with the necessary dependencies.
func NewService(r Repository, c *ukpolice.Client) Service {
	return &service{r, c}
}

func (s *service) CreateTables() {
	s.r.CreateTables()
}

func (s *service) UpdateData() error {
	avaliable, _, err := s.c.Availability.GetAvailabilityInfo(context.Background())
	if err != nil {
		return err
	}

	maxWorkers := 5
	maxRetries := 2

	requests := make(chan requestInfo, 100)
	store := make(chan ukpolice.Search)

	// start workers
	for i := 0; i <= maxWorkers; i++ {
		go reqWorker(requests, store, s.c, maxRetries)
		go storeWorker(store, s.r)
	}

	buildRequests(requests, avaliable)
	return nil
}

type requestInfo struct {
	force   string
	date    string
	err     error
	retries int
}

func buildRequests(r chan<- requestInfo, avaliable []ukpolice.AvailabilityInfo) {
	for _, month := range avaliable {
		for _, force := range month.StopAndSearch {
			if force == "west-midlands" { //@TODO: Remove hard coded force.
				r <- requestInfo{
					force: force,
					date:  month.Date,
				}
			}
		}
	}
	close(r)
}

func reqWorker(r chan requestInfo, s chan<- ukpolice.Search, c *ukpolice.Client, nMax int) {
	for req := range r {
		searches, _, err := c.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
			ukpolice.WithDate(req.date), ukpolice.WithForce(req.force))
		if err != nil {
			req.err = err
			req.retries++
			if req.retries < nMax {
				r <- req
			}
		}

		for _, search := range searches {
			search.Force = req.force
			s <- search
		}
	}
}

func storeWorker(s <-chan ukpolice.Search, r Repository) {
	for search := range s {
		if err := r.StoreSearch(search); err != nil {
			log.Println(err) // @TODO: logging solution.
		}
	}
}
