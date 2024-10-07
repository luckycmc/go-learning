package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
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
		Index("test1").
		Id("1").
		BodyJson(user1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed %s to %s, type %s\n", user1.FirstName, put1.Id, put1.Type)

	// 使用字符串
	user2 := `{"first_name":"ergou","last_name":"liu","age":30,"about":"hello gagaga","interests":["pingpang","basketball"]}`
	put2, err := client.Index().
		Index("test1").
		Id("2").
		BodyJson(user2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
}
func deleteDoc() {
	res, err := client.Delete().Index("test1").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("Deleted %s from %s\n", res.Id, res.Index)
}

func update() {
	res, err := client.Update().
		Index("maisi").
		Id("2").
		Doc(map[string]interface{}{"age": 30}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("Updated %s to %s\n", res.Id, res.Index)
}

func gets() {
	// 通过id查找
	get1, err := client.Get().Index("maisi").Id("2").Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	if get1.Found {
		fmt.Printf("got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

func query() {
	var res *elastic.SearchResult
	var err error
	// 取所有
	res, err = client.Search("maisi").Do(context.Background())
	printUser(res, err)

	// 字段相等
	q := elastic.NewQueryStringQuery("last_name:liu")
	res, err = client.Search("maisi").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printUser(res, err)
	// 条件查询
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "ergou"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(12))
	res, err = client.Search("maisi").Query(boolQ).Do(context.Background())
	printUser(res, err)

	// 短语搜索
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "hello")
	res, err = client.Search("maisi").Query(matchPhraseQuery).Do(context.Background())
	printUser(res, err)

	// 分析interest
	aggs := elastic.NewTermsAggregation().Field("interests.keyword") // field中需要加入keyword，因为在es中只有keyword类型的字符串可以使用termsAggregation
	res, err = client.Search("maisi").Aggregation("all_interests", aggs).Do(context.Background())
	printUser(res, err)
}

// 分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("maisi").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printUser(res, err)
}

// 打印数据
func printUser(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ User
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(User)
		fmt.Printf("%#v\n", t)
	}
}

func main() {
	create()
	deleteDoc()
	update()
	gets()
	query()
	list(1, 3)
}
