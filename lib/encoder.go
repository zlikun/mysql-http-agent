package lib

import (
	"encoding/json"
	"strings"
)

// 转换数据为JSON格式（字符串）
func EncodeJson(lst *[]map[string]interface{}) ([]byte, error) {
	// 对JSON结果作格式化（格式化非必要）
	return json.MarshalIndent(lst, "", strings.Repeat(" ", 2))
}
