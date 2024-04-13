package main

import (
	"flag"
	"log"

	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

var store pgvector.Store

const (
	loadAction           = "load"
	ragSearchAction      = "rag_search"
	semanticSearchAction = "semantic_search"
)

func main() {
	action := flag.String("action", "", "valid options: load, search, rag")
	source := flag.String("source", "", "enter a public link")
	query := flag.String("query", "", "enter the search query")
	maxResults := flag.Int("maxResults", 0, "enter maximum number of search results")

	flag.Parse()

	if *action == loadAction {
		err := loadDocs(*source)

		if err != nil {
			log.Fatal(err)
		}
	} else if *action == ragSearchAction {
		err := ragSearch(*query, *maxResults)

		if err != nil {
			log.Fatal(err)
		}
	} else if *action == semanticSearchAction {
		err := semanticSearch(*query, *maxResults)

		if err != nil {
			log.Fatal(err)
		}
	}

}
