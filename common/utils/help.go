/*
 * @Author: 尤_Ta
 * @Date:  23:47
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:47
 */

package utils

import (
	"encoding/json"
	"strings"
)

// AnyToMap interface 转 map
func AnyToMap(v any) (map[string]any, error) {
	dataJson, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var MapData map[string]any
	err = json.Unmarshal(dataJson, &MapData)
	if err != nil {
		return nil, err
	}
	return MapData, nil
}

// ParamsFilter 过滤指定数组中的key
func ParamsFilter(isFilterStr string, params map[string]any) map[string]any {
	var data = make(map[string]any)
	for key, value := range params {
		if value != "" {
			find := strings.Contains(isFilterStr, key)
			if !find {
				data[key] = value
			}
		}
	}
	return data
}
