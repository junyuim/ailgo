package core_utils

import (
	"encoding/json"
	"github.com/bytedance/sonic"
)

func MarshalJson(val interface{}) ([]byte, error) {
	return sonic.Marshal(val)
}

func MarshalJsonString(val interface{}) (string, error) {
	return sonic.MarshalString(val)
}

func MarshalPrettyJson(val interface{}) ([]byte, error) {
	return json.MarshalIndent(val, "", "  ")
}

func UnmarshalJson(buf []byte, val interface{}, path ...interface{}) error {
	if len(path) < 1 {
		return sonic.Unmarshal(buf, val)
	}

	node, err := sonic.Get(buf, path...)

	if err != nil {
		return err
	}

	data, err := node.MarshalJSON()

	if err != nil {
		return err
	}

	if err = sonic.Unmarshal(data, val); err != nil {
		return err
	}

	return nil
}

//func StringJoin(sep string, items ...string) string {
//	var itemsarray []string
//
//	for _, item := range items {
//		if len(item) > 0 {
//			itemsarray = append(itemsarray, item)
//		}
//	}
//
//	return strings.Join(itemsarray, sep)
//}
