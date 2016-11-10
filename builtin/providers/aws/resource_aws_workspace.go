package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/hashicorp/terraform/helper/resource"
)

func resourceAwsWorkSpace() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsWorkSpaceCreate,
		Read:   resourceAwsWorkSpaceRead,
		Update: resourceAwsWorkSpaceUpdate,
		Delete: resourceAwsWorkSpaceDelete,

		Schema: map[string]*schema.Schema{
			"bundle_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"directory_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"root_volume_encryption_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_volume_encryption_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"volume_encryption_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"workspace_properties": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true, // TODO: check this
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"running_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "AUTO_STOP",
							ForceNew: true, // TODO: check this
							ValidateFunc: func(v interface{}, k string) (ws []string, es []error) {
								validTypes := []string{"AUTO_STOP", "ALWAYS_ON"}
								value := v.(string)
								for validType, _ := range validTypes 
								// TODO: valid values "AUTO_STOP", "ALWAYS_ON"
								return
							},
						},
						"auto_stop_timeout_minutes": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true, // TODO: check this
							// TODO: ValidateFunc:
							// must be multiple of 60
						},
					},
				},
			},
		},
	}
}

// func resourceAwsWorkSpaceCreate {
// }
//
// func resourceAwsWorkSpaceRead {
// }
//
// func resourceAwsWorkSpaceUpdate {
// }
//
// func resourceAwsWorkSpaceDelete {
// }
