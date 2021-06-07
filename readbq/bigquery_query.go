// Not ready for production!!
package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func main() {
	//err := queryBasic(os.Stdout, "stanke-sandbox-2020-10")
	//fmt.Println("err =", err)

	err2 := writeSomething(context.Background(),
		"stanke-sandbox-2020-10",
		"log_sink",
		"stderr_20210607",
		[]*Item{
			{LogName: "Valentin 7", Timestamp: time.Now()},
			{LogName: "Valentin 8", Timestamp: time.Now()},
		})
	fmt.Println("err2 =", err2)
}

func writeSomething(ctx context.Context, projectID, dataset, table string, items []*Item) error {
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	ins := client.Dataset(dataset).Table(table).Inserter()

	return ins.Put(ctx, items)
}

type Item struct {
	LogName   string
	Timestamp time.Time
}

// Save implements the ValueSaver interface.
// This example disables best-effort de-duplication, which allows for higher throughput.
func (i *Item) Save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"logName":   i.LogName,
		"timestamp": i.Timestamp,
	}, "insertIDconstant", nil
}

// queryBasic demonstrates issuing a query and reading results.
func queryBasic(w io.Writer, projectID string) error {
	// projectID := "my-project-id"
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	q := client.Query(
		"SELECT * FROM `log_sink.stderr_20210607` " +
			"LIMIT 100")
	// Location must match that of the dataset(s) referenced in the query.
	q.Location = "US"
	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}
	if err := status.Err(); err != nil {
		return err
	}
	it, err := job.Read(ctx)
	if err != nil {
		return err
	}
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, row)
	}
	return nil
}

// [END bigquery_query]
