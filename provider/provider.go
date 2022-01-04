package provider

import (
	"context"
	"fmt"
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var secretManagerPrefix = "gcp:secretmanager:"

func Get(key string) string {
	if isSecretManager(key) {
		secretKey := getKey(key)
		value, err := getSecretManagerValue(secretKey)
		if err != nil {
			println(fmt.Sprintf("Error while envsubst from secret manager: %v", err))
		}
		return value
	}
	return os.Getenv(key)
}

func isSecretManager(key string) bool {
	return strings.HasPrefix(key, secretManagerPrefix)
}

func getKey(key string) string {
	if isSecretManager(key) {
		return strings.ReplaceAll(key, secretManagerPrefix, "")
	}
	return key
}

func getSecretManagerValue(key string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: key,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}

	return string(result.Payload.Data), nil
}
