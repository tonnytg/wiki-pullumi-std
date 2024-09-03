package main

import (
    "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Nome do bucket de armazenamento
        bucketName := "hello-world-bucket"

        // Criação do bucket no Google Cloud Storage
        bucket, err := storage.NewBucket(ctx, bucketName, &storage.BucketArgs{
            Website: storage.BucketWebsiteArgs{
                MainPageSuffix: pulumi.String("index.html"),
            },
        })
        if err != nil {
            return err
        }

        // Criação do arquivo index.html com conteúdo "Hello World"
        indexContent := `<html><body><h1>Hello, World!</h1></body></html>`
        _, err = storage.NewBucketObject(ctx, "index.html", &storage.BucketObjectArgs{
            Bucket: bucket.Name,
            Content: pulumi.String(indexContent),
            Name:    pulumi.String("index.html"),
        })
        if err != nil {
            return err
        }

        // Exporta a URL do site
        ctx.Export("bucketName", bucket.Url)

        return nil
    })
}

