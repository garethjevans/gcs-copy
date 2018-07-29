package gcsCopy

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

// Run executes a copy from one file in a bucket to another.
func Run(bucketName string, copyFrom string, copyTo string, googleApplicationCredentials string) {

	log.Println("--bucket-name set to: " + bucketName)
	log.Println("--copy-from set to: " + copyFrom)
	log.Println("--copy-to set to: " + copyTo)
	log.Println("--google-application-credentials set to: " + googleApplicationCredentials)

	log.Println("Beginning copy")

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	src := client.Bucket(bucketName).Object(copyFrom)
	dst := client.Bucket(bucketName).Object(copyTo)

	_, err = dst.CopierFrom(src).Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Complete")
}
