package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

// GetBaseTypeName returns a string containing a type's name with all package name prefixes removed
func GetBaseTypeName(i interface{}) string {
	t := reflect.TypeOf(i).String()
	ts := strings.Split(t, ".")
	return ts[len(ts)-1]
}

// StructToMap converts a struct type to a map via JSON marshalling/unmarshalling
func StructToMap(i interface{}) map[string]interface{} {
	newMap := map[string]interface{}{}
	inrec, _ := json.Marshal(i)
	json.Unmarshal(inrec, &newMap)
	return newMap
}
