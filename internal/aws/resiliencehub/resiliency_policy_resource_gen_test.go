// Code generated by generators/resource/main.go; DO NOT EDIT.

package resiliencehub_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSResilienceHubResiliencyPolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ResilienceHub::ResiliencyPolicy", "awscc_resiliencehub_resiliency_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
