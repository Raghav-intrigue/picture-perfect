/*
 * Picture Perfect
 *
 * Sample API description. You can find more at.
 *
 * API version: 0.1.0
 * Contact: 1raghavmahajan@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Address struct {

	StreetAdresses string `json:"streetAdresses,omitempty"`

	Pincode int32 `json:"pincode,omitempty"`

	City string `json:"city,omitempty"`
}
