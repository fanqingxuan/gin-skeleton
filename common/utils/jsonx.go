package utils

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(&v)
}

func Decode(str []byte, v interface{}) error {
	return json.Unmarshal(str, &v)
}
