package main

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tjcain/ukpolice"
)

const dateFormat = "2006-01"

func main() {

}

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

// func (pg *pgStore) GetDate() string {
// 	// do stuff
// 	return ""
// }
