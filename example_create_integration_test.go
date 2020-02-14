package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"log"
)

func Example_createIntegration() {
	client, err := rockset.NewClient(rockset.FromEnv())
	if err != nil {
		log.Fatal(err)
	}

	// create integration
	req := models.CreateIntegrationRequest{
		Name:        "my-first-integration",
		Description: "my first integration",
		Kinesis: &models.KinesisIntegration{
			AwsAccessKey: &models.AwsAccessKey{
				AwsAccessKeyId:     "...",
				AwsSecretAccessKey: "...",
			},
			AwsRole: &models.AwsRole{
				AwsRoleArn: "arn:aws:iam::<account-id>:role/<role-name>",
			},
		},
	}

	res, _, err := client.Integration.Create(req)
	if err != nil {
		log.Fatal(err)
	}

	res.PrintResponse()
}
