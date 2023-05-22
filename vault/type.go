package vault

import "time"

type Vault struct {
	Time time.Time `json:"time"`
	Type string    `json:"type"`
	Auth struct {
		DisplayName   string   `json:"display_name"`
		Policies      []string `json:"policies"`
		TokenPolicies []string `json:"token_policies"`
		PolicyResults struct {
			Allowed bool `json:"allowed"`
		} `json:"policy_results"`
		TokenIssueTime time.Time `json:"token_issue_time"`
	} `json:"auth"`
	Request struct {
		ID                  string `json:"id"`
		ClientID            string `json:"client_id"`
		Operation           string `json:"operation"`
		MountType           string `json:"mount_type"`
		MountAccessor       string `json:"mount_accessor"`
		ClientToken         string `json:"client_token"`
		ClientTokenAccessor string `json:"client_token_accessor"`
		Namespace           struct {
			ID string `json:"id"`
		} `json:"namespace"`
		Path          string `json:"path"`
		RemoteAddress string `json:"remote_address"`
		RemotePort    int    `json:"remote_port"`
	} `json:"request"`
	Response struct {
		MountType     string `json:"mount_type"`
		MountAccessor string `json:"mount_accessor"`
	} `json:"response"`
}
