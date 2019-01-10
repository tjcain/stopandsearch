// Grab information from ukpolice database and store in our Storage solution.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"stopandsearch/pkg/Storage/postgres"
	"sync"
	"time"

	"github.com/tjcain/ukpolice"
	"upper.io/db.v3/postgresql"
)

// This should come from some sort of config
const (
	host     = "localhost"
	port     = 5432
	password = ""
	user     = "postgres"
	database = "stopandsearch_test"
)

var records, errs, retries int
var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func main() {
	start := time.Now()
	// set up Storage
	var settings = postgresql.ConnectionURL{
		Database: `stopandsearch_test`,
		Host:     `localhost`,
		User:     `postgres`,
	}
	db, err := postgres.NewPostgresDB(settings)
	if err != nil {
		log.Fatalf("Could not connect to db: %s", err)
	}
	defer db.Close()
	// reset db
	// log.Println("RESETTING DATABASE")
	// err = db.DestructiveReset("wm_test_ss")
	// if err != nil {
	// 	log.Fatalf("could not reset DB %v", err)
	// }

	// defined clients
	customClient := http.Client{Timeout: time.Second * 120}
	client := ukpolice.NewClient(&customClient)

	// Grab avaliable data
	avaliable, _, err := getAvaliable(client)
	if err != nil {
		log.Fatal(err)
	}

	fetch(avaliable, client, db)
	fmt.Println("waiting on go routines...")

	wg.Wait()
	fmt.Printf("%.2fs elapsed\t records stored: %d\t errors: %d retries: %d\n",
		time.Since(start).Seconds(), records, errs, retries)

}

func getAvaliable(client *ukpolice.Client) ([]ukpolice.AvailabilityInfo, *ukpolice.Response, error) {
	return client.Availability.GetAvailabilityInfo(context.Background())
}

type Storage interface {
	//store methods
	StoreSearch(search ukpolice.Search) error
}

type search struct {
	ukpolice.StopAndSearchService
	force string
}

func fetch(available []ukpolice.AvailabilityInfo, client *ukpolice.Client, db Storage) {
	for _, month := range available {
		for _, force := range month.StopAndSearch {
			if force == "west-midlands" {
				wg.Add(1)
				go fetchAndStore(client, month.Date, force, db)
			}
		}
	}
}

func fetchAndStore(client *ukpolice.Client, date, force string, db Storage) {
	searches, _, err := client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
		ukpolice.WithDate(date), ukpolice.WithForce(force))

	if err != nil {
		mutex.Lock()
		retries++
		errs++
		mutex.Unlock()
		fmt.Printf("ERROR FROM %s-%s: %s... RETRYING...\n", date, force, err)
		go retryError(client, date, force, db)
		return
	}

	for _, search := range searches {
		// mutex.Lock()
		search.Force = force
		db.StoreSearch(search)
		// mutex.Unlock()
	}

	mutex.Lock()
	records += len(searches)
	mutex.Unlock()
	wg.Done()
}

func retryError(client *ukpolice.Client, date, force string, db Storage) {
	searches, _, err := client.StopAndSearch.GetStopAndSearchesByForce(context.Background(),
		ukpolice.WithDate(date), ukpolice.WithForce(force))

	if err != nil {
		fmt.Printf("ERROR ON RETRY FROM %s-%s: %s\n", date, force, err)
		wg.Done()
		return
	}

	for _, search := range searches {
		search.Force = force
		db.StoreSearch(search)
	}

	fmt.Printf("RETRY OF %s-%s: SUCCESSFUL\n", date, force)

	mutex.Lock()
	records += len(searches)
	errs--
	mutex.Unlock()
	wg.Done()
}
