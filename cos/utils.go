package cos

import (
	"fmt"
	"strings"

	"encoding/json"
)

func formatMapByKeys(keys []string, srcMap map[string]string) map[string]string {
	newMap := map[string]string{}
	if srcMap == nil {
		return nil
	}

	for _, k := range keys {
		if v, ok := srcMap[k]; ok {
			newMap[k] = v
		}
	}

	return newMap
}

func mapToStr(m map[string]string, sep string) string {
	if m == nil {
		return ""
	}

	l := []string{}
	for k := range m {
		l = append(l, fmt.Sprintf("%s=%s", k, m[k]))
	}

	return strings.Join(l, sep)
}

func printPretty(result interface{}) {
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}

func structToMap(v interface{}) map[string]string {
	if v == nil {
		return nil
	}
	var r map[string]string
	b, e := json.Marshal(v)
	if e != nil {
		fmt.Println("e:", e)
	}
	json.Unmarshal(b, &r)
	return r
}
