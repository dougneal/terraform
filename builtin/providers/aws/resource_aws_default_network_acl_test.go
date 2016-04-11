package aws

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSDefaultNetworkAcl_basic(t *testing.T) {
	var networkAcl ec2.NetworkAcl

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSDefaultNetworkAclDestroy,
		Steps: []resource.TestStep{
			// Tests that a default_network_acl will show a non-empty plan if no rules are
			// given, indicating that it wants to destroy the default rules
			resource.TestStep{
				Config: testAccAWSDefaultNetworkConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccGetWSDefaultNetworkAcl("aws_default_network_acl.default", &networkAcl),
				),
				ExpectNonEmptyPlan: true,
			},
			// Add default ACL rules and veryify plan is empty
			resource.TestStep{
				Config: testAccAWSDefaultNetworkConfig_basicDefaultRules,
				Check: resource.ComposeTestCheckFunc(
					testAccGetWSDefaultNetworkAcl("aws_default_network_acl.default", &networkAcl),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccCheckAWSDefaultNetworkAclDestroy(s *terraform.State) error {
	// conn := testAccProvider.Meta().(*AWSClient).ec2conn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_default_network_acl" {
			continue
		}

	}

	return nil
}

func testAccGetWSDefaultNetworkAcl(n string, networkAcl *ec2.NetworkAcl) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Network ACL is set")
		}
		conn := testAccProvider.Meta().(*AWSClient).ec2conn

		resp, err := conn.DescribeNetworkAcls(&ec2.DescribeNetworkAclsInput{
			NetworkAclIds: []*string{aws.String(rs.Primary.ID)},
		})
		if err != nil {
			return err
		}

		if len(resp.NetworkAcls) > 0 && *resp.NetworkAcls[0].NetworkAclId == rs.Primary.ID {
			*networkAcl = *resp.NetworkAcls[0]
			return nil
		}

		return fmt.Errorf("Network Acls not found")
	}
}

const testAccAWSDefaultNetworkConfig_basic = `
resource "aws_vpc" "tftestvpc" {
  cidr_block = "10.1.0.0/16"

  tags {
    Name = "TestAccAWSDefaultNetworkAcl_basic"
  }
}

resource "aws_default_network_acl" "default" {
  default_network_acl_id = "${aws_vpc.tftestvpc.default_network_acl_id}"

  tags {
    Name = "TestAccAWSDefaultNetworkAcl_basic"
  }
}
`

const testAccAWSDefaultNetworkConfig_basicDefaultRules = `
resource "aws_vpc" "tftestvpc" {
  cidr_block = "10.1.0.0/16"

  tags {
    Name = "TestAccAWSDefaultNetworkAcl_basic"
  }
}

resource "aws_default_network_acl" "default" {
  default_network_acl_id = "${aws_vpc.tftestvpc.default_network_acl_id}"

  ingress {
    protocol   = -1
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }

  egress {
    protocol   = -1
    rule_no    = 100
    action     = "allow"
    cidr_block = "0.0.0.0/0"
    from_port  = 0
    to_port    = 0
  }

  tags {
    Name = "TestAccAWSDefaultNetworkAcl_basic"
  }
}
`
