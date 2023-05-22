package resources

import (
	"context"

	"github.com/alyragab/cq-source-vault/client"
	"github.com/alyragab/cq-source-vault/vault"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

// Transform Vault audit data
func VaultTable() *schema.Table {
	return &schema.Table{
		Name:      "vault",
		Resolver:  fetchVaultAudit,
		Transform: transformers.TransformWithStruct(&vault.Vault{}),
		// Columns: []schema.Column{
		// 	{
		// 		Name: "audit",
		// 		Type: schema.TypeString,
		// 	},
		// },
	}
}

func fetchVaultAudit(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	c.ID()
	res <- vault.CH
	return nil
}
