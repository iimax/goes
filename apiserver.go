package main

import (
	"context"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net/http"
	//"log"
	"github.com/olivere/elastic/v7"
)

type DeviceLog struct {
	DeviceId  string
	Time      string
	LocalTime string
	Product   string
	Msg       string
	Level     string
	Exception string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/api/log/devices", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// 从es获取数据
		client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
		if err != nil {
			panic(err)
		}
		// https://godoc.org/github.com/olivere/elastic
		// Search with a term query
		//termQuery := elastic.NewTermQuery("level", "DEBUG")
		// boolQuery := elastic.NewBoolQuery()
		// boolQuery = boolQuery.Should(elastic.NewTermQuery("level", "D"))
		matchQuery := elastic.NewMatchQuery("level", "DEBUG")
		src, err := matchQuery.Source()
		if err != nil {
			panic(err)
		}
		data, err := json.MarshalIndent(src, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
		searchResult, err := client.Search().
			Index("logs-*").
			Query(matchQuery).
			From(0).Size(10).
			Do(ctx) // execute
		fmt.Println("search executed")
		if err != nil {
			// Handle error
			fmt.Fprintf(w, "error! ")
			panic(err)
		}

		// searchResult is of type SearchResult and returns hits, suggestions,
		// and all kinds of other information from Elasticsearch.
		fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
		// 将es查询结果 json输出
		//jsonBytes, err := json.Marshal(searchResult.Hits)
		//fmt.Fprintf(w, jsonBytes)
		json.NewEncoder(w).Encode(searchResult)
	})

	fmt.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
