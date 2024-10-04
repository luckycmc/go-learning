package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.72.130:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to elasticsearch success")
	p1 := Person{
		Name:    "kevin",
		Age:     23,
		Married: true,
	}
	put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
