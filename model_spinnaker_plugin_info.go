/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SpinnakerPluginInfo struct {
	Provider string `json:"provider,omitempty"`
	Id string `json:"id,omitempty"`
	Releases []SpinnakerPluginRelease `json:"releases"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	RepositoryId string `json:"repositoryId,omitempty"`
	ProjectUrl string `json:"projectUrl,omitempty"`
}
