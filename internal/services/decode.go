package services

import (
	"encoding/json"
	"reflect"
	"strings"
)

// decodeJSONLines 解析 wslc 的 --format json 输出为 []T。
// wslc 可能输出一个 JSON 数组，也可能每行一个 JSON 对象（JSONL），两者都兼容。
// 字段匹配大小写不敏感，因此模型上的 camelCase json 标签可绑定到 wslc 的 PascalCase 键。
func decodeJSONLines[T any](out string) ([]T, error) {
	out = strings.TrimSpace(out)
	if out == "" {
		return nil, nil
	}

	// 先尝试整体作为 JSON 数组解析
	var arr []T
	if err := json.Unmarshal([]byte(out), &arr); err == nil && len(arr) > 0 {
		return arr, nil
	}

	var res []T
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var raw map[string]json.RawMessage
		if err := json.Unmarshal([]byte(line), &raw); err != nil {
			continue
		}
		norm := make(map[string]json.RawMessage, len(raw))
		for k, v := range raw {
			norm[strings.ToLower(k)] = v
		}

		var item T
		rv := reflect.ValueOf(&item).Elem()
		rt := rv.Type()
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			tag := f.Tag.Get("json")
			if tag == "" || tag == "-" {
				continue
			}
			key := strings.Split(tag, ",")[0]
			key = strings.TrimSpace(key)
			if r, ok := norm[strings.ToLower(key)]; ok {
				_ = json.Unmarshal(r, rv.Field(i).Addr().Interface())
			}
		}
		res = append(res, item)
	}
	return res, nil
}
