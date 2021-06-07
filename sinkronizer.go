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

	// // Creates a client.
	// client, err := logging.NewClient(ctx, projectID)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	// Sets the name of the log to write to.
	logName := "stdout"

	// Selects the log to write to.
	// logger := client.Logger(logName)

	// Sets the data to log.
	// text := "Hello, world!"

	// Adds an entry to the log buffer.
	// logger.Log(logging.Entry{Payload: text})

	// Closes the client and flushes the buffer to the Cloud Logging
	// service.
	// if err := client.Close(); err != nil {
	// 	log.Fatalf("Failed to close client: %v", err)
	// }

	// fmt.Printf("Logged: %v\n", text)

	entries, err := fetchLogs(ctx, projectID, logName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entries)

	for _, entry := range entries {
		fmt.Println(entry)
	}

}

func fetchLogs(ctx context.Context, projID string, logName string) ([]*logging.Entry, error) {
	var entries []*logging.Entry
	lastHour := time.Now().Add(-1 * time.Hour).Format(time.RFC3339)

	adminClient, err := logadmin.NewClient(ctx, projID)
	if err != nil {
		log.Fatalf("Failed to create logadmin client: %v", err)
	}
	defer adminClient.Close()

	iter := adminClient.Entries(ctx,
		// Only get entries from the "log-example" log within the last hour.
		logadmin.Filter(fmt.Sprintf(`logName = "projects/%s/logs/%s" AND timestamp > "%s"`, projID, logName, lastHour)),
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
