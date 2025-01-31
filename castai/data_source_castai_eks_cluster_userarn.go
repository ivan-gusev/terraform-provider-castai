package castai

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/castai/terraform-provider-castai/castai/sdk"
)

const (
	EKSClusterUserARNFieldClusterID = "cluster_id"
	EKSClusterUserARNFieldARN = "arn"
)

func dataSourceCastaiEKSClusterUserARN() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCastaiEKSUserARN,
		Schema: map[string]*schema.Schema{
			EKSClusterUserARNFieldClusterID: {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
			},
			EKSClusterUserARNFieldARN: {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCastaiEKSUserARN(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ProviderConfig).api

	clusterID := data.Get(EKSClusterUserARNFieldClusterID).(string)

	resp, err := client.ExternalClusterAPIGetAssumeRoleUserWithResponse(ctx, clusterID)
	if checkErr := sdk.CheckOKResponse(resp, err); checkErr != nil {
		return diag.FromErr(checkErr)
	}

	arn := *resp.JSON200.Arn

	data.SetId(arn)
	data.Set(EKSClusterUserARNFieldARN, arn)

	return nil
}