package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/tmc/langchaingo/embeddings/bedrock"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func getVectorStore() (vectorstores.VectorStore, error) {

	host := os.Getenv("PG_HOST")
	if host == "" {
		log.Fatal("missing PG_HOST")
	}

	user := os.Getenv("PG_USER")
	if user == "" {
		log.Fatal("missing PG_USER")
	}

	password := os.Getenv("PG_PASSWORD")
	if password == "" {
		log.Fatal("missing PG_PASSWORD")
	}

	dbName := os.Getenv("PG_DB")
	if dbName == "" {
		log.Fatal("missing PG_DB")
	}

	connURLFormat := "postgres://%s:%s@%s:5432/%s?sslmode=disable"

	pgConnURL := fmt.Sprintf(connURLFormat, user, url.QueryEscape(password), host, dbName)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	brc := bedrockruntime.NewFromConfig(cfg)

	embeddingModel, err := bedrock.NewBedrock(bedrock.WithClient(brc), bedrock.WithModel(bedrock.ModelTitanEmbedG1))

	if err != nil {
		return nil, err
	}
	store, err = pgvector.New(
		context.Background(),
		//pgvector.WithPreDeleteCollection(true),
		pgvector.WithConnectionURL(pgConnURL),
		pgvector.WithEmbedder(embeddingModel),
	)
	if err != nil {
		return nil, err
	}

	log.Println("vector store ready")

	return store, nil
}
