package prayertexter

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func generateID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		slog.Error("failed to generate random bytes")
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func removeItem[T comparable](items *[]T, target T) {
	slice := *items
	var newItems []T

	for _, v := range slice {
		if v != target {
			newItems = append(newItems, v)
		}
	}

	*items = newItems
}

func getAwsConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}
