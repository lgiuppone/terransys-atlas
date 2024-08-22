package main

import (
	"github.com/terransys/terransys-atlas/internal/eksroles"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		eksroles.CreateRoles(ctx)
		return nil
	})
}
