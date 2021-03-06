/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Url struct {
	Path string `json:"path,omitempty"`
	DeserializedFields *UrlStreamHandler `json:"deserializedFields,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	UserInfo string `json:"userInfo,omitempty"`
	Authority string `json:"authority,omitempty"`
	File string `json:"file,omitempty"`
	Host string `json:"host,omitempty"`
	Content *interface{} `json:"content,omitempty"`
	SerializedHashCode int32 `json:"serializedHashCode,omitempty"`
	DefaultPort int32 `json:"defaultPort,omitempty"`
	Port int32 `json:"port,omitempty"`
	Query string `json:"query,omitempty"`
	Ref string `json:"ref,omitempty"`
}
