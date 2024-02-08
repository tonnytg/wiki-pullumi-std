package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new GCP project
		project, err := compute.NewProject(ctx, "my-project", &compute.ProjectArgs{
			ProjectId: pulumi.String("my-project-id"),
		})
		if err != nil {
			return err
		}

		// Enable the Compute Engine API for the project
		_, err = compute.NewService(ctx, "compute-service", &compute.ServiceArgs{
			Project: project.ProjectId,
			Service: pulumi.String("compute.googleapis.com"),
		})
		if err != nil {
			return err
		}

		// Create a new Compute Engine instance
		_, err = compute.NewInstance(ctx, "my-instance", &compute.InstanceArgs{
			Zone:        pulumi.String("us-central1-a"),
			MachineType: pulumi.String("e2-micro"),
			BootDisk: &compute.InstanceBootDiskArgs{
				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String("debian-cloud/debian-9"),
				},
			},
			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
				&compute.InstanceNetworkInterfaceArgs{
					Network: pulumi.String("default"),
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
