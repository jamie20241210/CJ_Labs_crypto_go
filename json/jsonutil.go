package json

import "encoding/json"

// ToJSON 将对象转换为 JSON 字符串
func ToJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToJSONString 将对象转换为 JSON 字符串，如果出错则返回空字符串
func ToJSONString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// FromJSON 将 JSON 字符串转换为对象
func FromJSON(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
