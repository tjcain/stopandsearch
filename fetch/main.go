package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"github.com/tjcain/ukpolice"
)

const dateFormat = "2006-01"

func main() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	config := LoadConfig()
	fmt.Println(config)

	// db := pgStore{db: nil}
	// client := ukpolice.NewClient(&http.Client{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := cron.New()
	c.AddFunc("@every 2s", func() { fmt.Println("Every second") })
	c.Start()

	// u, err := GetUpdateSlice(db, client)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// for _, f := range u {
	// 	fmt.Println(f.Date)
	// }
	sig := <-gracefulStop
	fmt.Printf("caught sig: %+v", sig)
	fmt.Println("Wait for 2 second to finish processing")
	time.Sleep(2 * time.Second)
	os.Exit(0)
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

func (pg pgStore) GetDate() string {
	// do stuff
	return "2017-02"
}
