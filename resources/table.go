package resources

import (
	"context"

	"github.com/alyragab/cq-source-vault/vault"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

// Transform Vault audit data
func VaultTable() *schema.Table {
	return &schema.Table{
		Name:     "vault",
		Resolver: fetchVaultTable,
		Columns: []schema.Column{
			{
				Name: "audit",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchVaultTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	vault.WatchFile()
	return nil
}
