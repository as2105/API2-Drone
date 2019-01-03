package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

func GetBaseTypeName(i interface{}) string {
	t := reflect.TypeOf(i).String()
	ts := strings.Split(t, ".")
	return ts[len(ts)-1]
}

func StructToMap(i interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	inrec, _ := json.Marshal(i)
	json.Unmarshal(inrec, &newMap)
	return newMap
}
