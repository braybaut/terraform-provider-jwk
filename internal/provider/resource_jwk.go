package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-jwk/internal/jwk"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceJwk() *schema.Resource {
	return &schema.Resource{
		CreateContext: createResourceJwk,
		DeleteContext: deleteResourceJwk,
		ReadContext:   readResourceJwk,

		Schema: map[string]*schema.Schema{
			"public_key_pem": {
				Type:        schema.TypeString,
				Description: "Key name to be used to create the jwk document",
				Required:    true,
				ForceNew:    true,
			},
			"jwk_document": {
				Type:        schema.TypeString,
				Description: "jwk document generated",
				Computed:    true,
			},
		},
	}
}

func createResourceJwk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	keypair := d.Get("public_key_pem")
	keypairstring := fmt.Sprintf("%v", keypair)

	value, err := jwk.CreateJwk(keypairstring)
	if err != nil {
		return diag.Errorf("failed to create JWK document %s", err)
	}

	jwkvalues := jwkValues{}
	err = json.Unmarshal(value, &jwkvalues)
	if err != nil {
		return diag.Errorf("error to Unmarshal struct %s", err)
	}
	d.SetId(jwkvalues.Keys[0].Kid)
	tflog.Trace(ctx, "Create a resource")

	if err := d.Set("jwk_document", string(value)); err != nil {
		return diag.Errorf("Failed to define jwk_document attribute %s", err)
	}
	return diags
}

func deleteResourceJwk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
func readResourceJwk(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
