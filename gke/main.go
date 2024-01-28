package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Configurar o cluster GKE
		cluster, err := container.NewCluster(ctx, "my-cluster", &container.ClusterArgs{
			InitialNodeCount: pulumi.Int(2),
			MinMasterVersion: pulumi.String("latest"),
			NodeConfig: &container.ClusterNodeConfigArgs{
				MachineType: pulumi.String("n1-standard-1"),
				OauthScopes: pulumi.StringArray{
					pulumi.String("https://www.googleapis.com/auth/compute"),
					pulumi.String("https://www.googleapis.com/auth/devstorage.read_only"),
					pulumi.String("https://www.googleapis.com/auth/logging.write"),
					pulumi.String("https://www.googleapis.com/auth/monitoring"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("clusterName", cluster.Name)
		ctx.Export("clusterEndpoint", cluster.Endpoint)
		ctx.Export("clusterMasterVersion", cluster.MasterVersion)
		return nil
	})
}
