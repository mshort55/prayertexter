package utility

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("generateID: %w", err)
	}

	return hex.EncodeToString(bytes), nil
}

func RemoveItem[T comparable](items *[]T, target T) {
	slice := *items
	var newItems []T

	for _, v := range slice {
		if v != target {
			newItems = append(newItems, v)
		}
	}

	*items = newItems
}
