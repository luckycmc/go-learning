package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
)

var client *elastic.Client
var host = "http://192.168.72.130:9200"

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {
	errlog := log.New(os.Stdout, "App", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

// create
func create() {
	// 使用结构体
	user1 := User{
		"kevin1",
		"chen",
		30,
		"hello",
		[]string{"pingpang, music"},
	}
	put1, err := client.Index().
		Index("maisi").
		Id("2").
		BodyJson(user1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed %s to %s, type %s\n", user1.FirstName, put1.Id, put1.Type)
}

func main() {
	create()
}
