package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tjcain/ukpolice"
)

const dateFormat = "2006-01"

type Fetch struct {
	db DataStore
	// cfg Config
}

func main() {
	// c := time.Tick(1 * time.Minute)
	// for now := range c {
	// 	fmt.Printf("%v %s\n", now, statusUpdate())
	// }
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

// New returns a Fetch with the sql.DB set with the postgres
// DB connection string in the configuration
// func New() (*Fetch, error) {
// 	cfg, err := LoadConfig()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var fetch Fetch

// 	db, err := sql.Open("postgres", fmt.Sprintf(
// 		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
// 		cfg.Database.User, cfg.Database.Password, cfg.Database.Database,
// 		cfg.Database.Host, cfg.Database.Port))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Fetch{db: db}, nil
// }

func statusUpdate() string { return "stats" }

// GetUpdateSlice returns a slice of avaliability information newer than that
// currently stored in the DataStore. Nil is returned if the DataStore is up to date.
func GetUpdateSlice(d DataStore, c *ukpolice.Client) ([]ukpolice.AvailabilityInfo, error) {
	storedDate := d.GetDate()
	t, err := time.Parse(dateFormat, storedDate)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	avaliable, _, err := c.Availability.GetAvailabilityInfo(context.Background())
	if err != nil {
		log.Println("getAvaliable()", err)
		return nil, err
	}

	var avaliabilities []ukpolice.AvailabilityInfo

	for _, dates := range avaliable {
		dt, err := time.Parse(dateFormat, dates.Date)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if t.Before(dt) {
			avaliabilities = append(avaliabilities, dates)
		}
	}

	return avaliabilities, nil
}

// DataStore describes a database
type DataStore interface {
	GetDate() string
}

type pgStore struct {
	db *sqlx.DB
}

func (pg pgStore) GetDate() string {
	// do stuff
	return "2017-02"
}
