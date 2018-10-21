package lib

import (
	"reflect"
)

func UpdateBuilder(old, new interface{}) interface{} {
	oldVal := reflect.ValueOf(old).Elem()
	newVal := reflect.ValueOf(new).Elem()

	for i := 0; i < oldVal.NumField(); i++ {
		for j := 0; j < newVal.NumField(); j++ {
			if oldVal.Type().Field(i).Name == newVal.Type().Field(j).Name {
				if newVal.Field(j).Interface() != nil {
					oldVal.Field(i).Set(newVal.Field(j))
				}
			}
		}
	}
	return oldVal.Interface()
}
