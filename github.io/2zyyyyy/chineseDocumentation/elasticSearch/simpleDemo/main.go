package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	// handle error
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success!")
	p1 := Person{Name: "2zyyyyy", Age: 25, Married: false}
	put, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())
	// handle error
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}
