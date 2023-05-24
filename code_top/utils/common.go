package utils

import jsoniter "github.com/json-iterator/go"

func Stringify(obj interface{}) string {
	ret, _ := jsoniter.MarshalToString(obj)
	return ret
}
