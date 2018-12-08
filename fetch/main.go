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

}

// UpdateRequired returns true if the current dataset is out of date.
func UpdateRequired(d DataStore, c *ukpolice.Client) bool {
	const dateFormat = "2006-01"

	storedDate := d.GetDate()
	t, err := time.Parse(dateFormat, storedDate)
	if err != nil {
		log.Fatalln(err)
	}

	avaliable, _, err := getAvaliable(c)
	if err != nil {
		log.Fatalln("getAvaliable()", err)
	}

	for _, dates := range avaliable {
		dt, err := time.Parse(dateFormat, dates.Date)
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
