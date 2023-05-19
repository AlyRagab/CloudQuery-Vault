package resources

import (
	"context"

	"github.com/alyragab/cq-source-vault/vault"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func VaultTable() *schema.Table {
	return &schema.Table{
		Name:     "vault_table",
		Resolver: fetchVaultTable,
		Columns: []schema.Column{
			{
				Name: "column",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchVaultTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) {
	vault.WatchFile()
}
