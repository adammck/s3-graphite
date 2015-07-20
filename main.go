package main

import (
	"github.com/Sirupsen/logrus"
	"os"
	"time"
)

func main() {
	name := "s3-graphite"
	logrus.Printf("Starting %s...", name)

	// TODO: Read these from somewhere more sensible
	s3_bucket := os.Getenv("S3_BUCKET")
	s3_prefix := os.Getenv("S3_PREFIX")
	graphite_address := os.Getenv("GRAPHITE_ADDRESS")
	graphite_prefix := os.Getenv("GRAPHITE_PREFIX")

	s3, err := NewS3(s3_bucket, s3_prefix)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	graphite, err := NewGraphite(graphite_address, graphite_prefix)
	defer graphite.Close()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	c := make(chan int, 100)
	go FetchCounts(s3, c, 10*time.Second)
	go SendCounts(graphite, c)

	// Block forever.
	select {}
}

func FetchCounts(s3 *S3, objectCounts chan int, interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		c, err := s3.Count()
		if err == nil {
			objectCounts <- c
		} else {
			logrus.Error(err)
		}
	}
}

func SendCounts(graphite *Graphite, objectsCount chan int) {
	for c := range objectsCount {
		graphite.Send(c)
	}
}
