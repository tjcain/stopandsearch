package main

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tjcain/ukpolice"
)

func main() {
	// customClient := http.Client{Timeout: time.Second * 120}
	// client := ukpolice.NewClient(&customClient)
	// db := &pgStore{db: nil}

	// UpdateRequired(db, client)

}

// UpdateRequired returns true if the current dataset is out of date.
func UpdateRequired(d DataStore, c *ukpolice.Client) bool {
	const shortDate = "2006-01"

	storedDate := d.GetDate()
	t, err := time.Parse(shortDate, storedDate)
	if err != nil {
		log.Fatalln(err)
	}

	avaliable, _, err := getAvaliable(c)
	if err != nil {
		log.Fatal(err)
	}

	for _, dates := range avaliable {
		dt, err := time.Parse(shortDate, dates.Date)
		if err != nil {
			log.Fatalln(err)
		}

		if t.Before(dt) {
			return true
		}
	}

	return false
}

// DataStore describes a database
type DataStore interface {
	GetDate() string
}

type pgStore struct {
	db *sqlx.DB
}

func (pg *pgStore) GetDate() string {
	// do stuff
	return ""
}

func getAvaliable(client *ukpolice.Client) ([]ukpolice.AvailabilityInfo, *ukpolice.Response, error) {
	return client.Availability.GetAvailabilityInfo(context.Background())
}
