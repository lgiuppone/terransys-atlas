package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/your-org/terransys-atlas/internal/aws/ekscluster"
	"github.com/your-org/terransys-atlas/internal/aws/eksinit"
	"github.com/your-org/terransys-atlas/internal/aws/eksnet"
	"github.com/your-org/terransys-atlas/internal/aws/eksroles"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Crear IAM roles en AWS
		roles, err := eksroles.CreateRoles(ctx)
		if err != nil {
			return err
		}

		// Configurar la red en AWS
		network, err := eksnet.SetupNetwork(ctx)
		if err != nil {
			return err
		}

		// Desplegar el clúster EKS en AWS
		cluster, err := ekscluster.DeployCluster(ctx, roles, network)
		if err != nil {
			return err
		}

		// Inicializar el clúster con configuraciones necesarias en AWS
		err = eksinit.InitializeCluster(ctx, cluster)
		if err != nil {
			return err
		}

		return nil
	})
}
