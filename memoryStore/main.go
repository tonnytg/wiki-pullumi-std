package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/redis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Define the region and tier
		region := "us-central1"
		tier := "BASIC"

		// Create a new Redis instance
		_, err := redis.NewInstance(ctx, "myRedisInstance", &redis.InstanceArgs{
			Tier:        pulumi.String(tier),
			MemorySizeGb: pulumi.Int(1),
			Region:      pulumi.String(region),
			RedisVersion: pulumi.String("REDIS_6_X"),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

