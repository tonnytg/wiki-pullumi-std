package main

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/secretmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new GCP Secret Manager secret
		secret, err := secretmanager.NewSecret(ctx, "my-secret", &secretmanager.SecretArgs{
			SecretId: pulumi.String("my-secret"),
		})
		if err != nil {
			return err
		}

		// Create a new GCP Secret Manager secret version with a user and password
		_, err = secretmanager.NewSecretVersion(ctx, "my-secret-version", &secretmanager.SecretVersionArgs{
			SecretId: pulumi.String(secret.SecretId),
			SecretData: pulumi.StringMap{
				"username": pulumi.String("my-username"),
				"password": pulumi.String("my-password"),
			},
		})
		if err != nil {
			return err
		}

		// Export the secret ID
		ctx.Export("secretId", secret.SecretId)

		// Retrieve the secret value
		secretValue, err := secretmanager.GetSecretSecretVersion(ctx, &secretmanager.GetSecretSecretVersionArgs{
			SecretId:  pulumi.String(secret.SecretId),
			VersionId: pulumi.String("latest"), // Using the latest version of the secret
		})
		if err != nil {
			return err
		}

		// Use the secret value
		username, ok := secretValue.SecretData["username"]
		if !ok {
			return fmt.Errorf("username not found in secret data")
		}
		password, ok := secretValue.SecretData["password"]
		if !ok {
			return fmt.Errorf("password not found in secret data")
		}
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Password: %s\n", password)

		return nil
	})
}
