// Sample logging-quickstart writes a log entry to Cloud Logging.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/logadmin"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "stanke-sandbox-2020-10"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name of the log to write to.
	logName := "stdout"

	// Selects the log to write to.
	logger := client.Logger(logName)

	// Sets the data to log.
	text := "Hello, world!"

	// Adds an entry to the log buffer.
	logger.Log(logging.Entry{Payload: text})

	// Closes the client and flushes the buffer to the Cloud Logging
	// service.
	if err := client.Close(); err != nil {
		log.Fatalf("Failed to close client: %v", err)
	}

	fmt.Printf("Logged: %v\n", text)
}

func fetchLogs(ctx context.Context, adminClient *logging.Client) ([]*logging.Entry, error) {
	var entries []*logging.Entry
	const name = "log-example"
	lastHour := time.Now().Add(-1 * time.Hour).Format(time.RFC3339)

	iter := adminClient.Entries(ctx,
		// Only get entries from the "log-example" log within the last hour.
		logadmin.Filter(fmt.Sprintf(`logName = "projects/%s/logs/%s" AND timestamp > "%s"`, projID, name, lastHour)),
		// Get most recent entries first.
		logadmin.NewestFirst(),
	)

	// Fetch the most recent 20 entries.
	for len(entries) < 20 {
		entry, err := iter.Next()
		if err == iterator.Done {
			return entries, nil
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
