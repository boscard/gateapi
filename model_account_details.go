/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AccountDetails struct {
	CloudProvider string `json:"cloudProvider,omitempty"`
	RequiredGroupMembership []string `json:"requiredGroupMembership,omitempty"`
	PrimaryAccount bool `json:"primaryAccount,omitempty"`
	Type_ string `json:"type,omitempty"`
	Environment string `json:"environment,omitempty"`
	Name string `json:"name,omitempty"`
	AccountId string `json:"accountId,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	Permissions map[string][]string `json:"permissions,omitempty"`
	ChallengeDestructiveActions bool `json:"challengeDestructiveActions,omitempty"`
}