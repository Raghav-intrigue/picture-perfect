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

type FilterOptions2 struct {

	MovieId int64 `json:"movieId,omitempty"`

	UserId int64 `json:"userId,omitempty"`

	StartDate string `json:"start_date,omitempty"`

	EndDate string `json:"end_date,omitempty"`
}