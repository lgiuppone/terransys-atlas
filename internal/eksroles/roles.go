package eksroles

import (
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateRoles crea los roles necesarios para EKS
func CreateRoles(ctx *pulumi.Context) error {
    _, err := iam.NewRole(ctx, "eksRole", &iam.RoleArgs{
        AssumeRolePolicy: pulumi.String(`{
            "Version": "2012-10-17",
            "Statement": [{
                "Effect": "Allow",
                "Principal": {
                    "Service": "ec2.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }]
        }`),
        Tags: pulumi.StringMap{
            "Name": pulumi.String("eksRole"),
        },
    })
    return err
}
