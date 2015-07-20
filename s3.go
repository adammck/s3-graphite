package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	maxPages = 10
)

type S3 struct {
	client *s3.S3
	params *s3.ListObjectsInput
}

func NewS3(bucket string, path string) (*S3, error) {
	return &S3{
		s3.New(nil),
		&s3.ListObjectsInput{
			Bucket:    aws.String(bucket),
			Prefix:    aws.String(path),
			Delimiter: aws.String("/"),
		},
	}, nil
}

func (s *S3) Count() (int, error) {
	pageNum := 0
	objectCount := 0

	err := s.client.ListObjectsPages(s.params, func(page *s3.ListObjectsOutput, isLastPage bool) bool {
		pageNum++

		// Abort if we've fetched more than the page number limit. Each page
		// contains 1000 objects by default, so this is mostly in case I screwed up
		// and we're stuck in an infinite loop.
		if pageNum > maxPages {
			return false
		}

		objectCount += len(page.CommonPrefixes)
		return true
	})

	if err != nil {
		return 0, err
	}

	logrus.Debugf("Object count: %d", objectCount)
	return objectCount, nil
}
