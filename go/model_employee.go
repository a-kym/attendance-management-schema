/*
 * 勤怠管理API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Employee struct {
	// 姓と名の間には半角スペースを設ける 
	FullName string `json:"fullName"`
	// 性別
	Gender string `json:"gender"`
	// 誕生日   ISO 8601のYYYY-MM-DD形式 
	Birthday string `json:"birthday"`
	// 役割
	Position string `json:"position"`
}
