package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func loadDocs(source string) error {

	fmt.Println("loading data from", source)

	store, err := getVectorStore()

	if err != nil {
		return err
	}
	docs, err := getDocs(source)

	if err != nil {
		return err
	}

	fmt.Println("no. of documents to be loaded", len(docs))

	_, err = store.AddDocuments(context.Background(), docs)

	if err != nil {
		return err
	}

	fmt.Println("data successfully loaded into vector store")

	return nil
}

func getDocs(source string) ([]schema.Document, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	docs, err := documentloaders.NewHTML(resp.Body).LoadAndSplit(context.Background(), textsplitter.NewRecursiveCharacter())

	if err != nil {
		return nil, err
	}

	return docs, nil
}
