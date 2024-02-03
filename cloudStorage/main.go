package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP Storage Bucket
		bucket, err := storage.NewBucket(ctx, "my-storage-bucket", &storage.BucketArgs{
			Location: pulumi.String("us-central1"), // Change the region as needed
		})
		if err != nil {
			return err
		}

		// Upload a file to the Storage Bucket
		object, err := storage.NewBucketObject(ctx, "my-file-object", &storage.BucketObjectArgs{
			Bucket: bucket.Name,
			Source: pulumi.NewFileAsset("path/to/your/local/file.txt"), // Change the path accordingly
		})
		if err != nil {
			return err
		}

		// Export the bucket name and object self-link
		ctx.Export("bucketName", bucket.Name)
		ctx.Export("objectSelfLink", object.SelfLink)

		return nil
	})
}
