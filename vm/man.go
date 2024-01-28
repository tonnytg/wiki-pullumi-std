package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Criar uma nova instância VM
		instance, err := compute.NewInstance(ctx, "my-instance", &compute.InstanceArgs{
			MachineType: pulumi.String("f1-micro"),
			Zone:        pulumi.String("us-central1-a"),
			BootDisk: &compute.InstanceBootDiskArgs{
				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String("projects/debian-cloud/global/images/family/debian-10"),
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

		ctx.Export("instanceName", instance.Name)
		ctx.Export("instanceIP", instance.NetworkInterfaces.Index(pulumi.Int(0)).AccessConfigs.Index(pulumi.Int(0)).NatIp)
		return nil
	})
}
