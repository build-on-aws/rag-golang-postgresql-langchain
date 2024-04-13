package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/bedrock"
	"github.com/tmc/langchaingo/vectorstores"
)

func ragSearch(question string, numOfResults int) error {

	store, err := getVectorStore()

	if err != nil {
		return err
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}
	brc := bedrockruntime.NewFromConfig(cfg)

	llm, err := bedrock.New(bedrock.WithClient(brc), bedrock.WithModel(bedrock.ModelAnthropicClaudeV3Sonnet))
	//llm.CallbacksHandler = callbacks.LogHandler{}

	if err != nil {
		return err
	}

	result, err := chains.Run(
		context.Background(),
		chains.NewRetrievalQAFromLLM(
			llm,
			vectorstores.ToRetriever(store, numOfResults),
		),
		question,
		chains.WithMaxTokens(2048),
	)
	if err != nil {
		return err
	}

	fmt.Println("====final answer====\n", result)

	return nil

}

func semanticSearch(searchQuery string, maxResults int) error {

	store, err := getVectorStore()
	if err != nil {
		return err
	}

	searchResults, err := store.SimilaritySearch(context.Background(), searchQuery, maxResults)

	if err != nil {
		return err
	}

	fmt.Println("============== similarity search results ==============")

	for _, doc := range searchResults {
		fmt.Println("similarity search info -", doc.PageContent)
		fmt.Println("similarity search score -", doc.Score)
		fmt.Println("============================")

	}

	return nil

}
