package easycondition

import (
	"fmt"
	"reflect"
	"strings"
)

func FirstOrDefault(source interface{}, condition interface{}) (interface{}, error) {
	s := reflect.Indirect(reflect.ValueOf(source))
	if s.Kind() != reflect.Slice {
		panic("source' type is not slice")
	}

	c := reflect.Indirect(reflect.ValueOf(condition))
	if c.Kind() != reflect.String || c.IsZero() == true {
		panic("condition' type is not string or const value empty")
	}

	for i := 0; i < s.Len(); i++ {
		item := s.Index(i)
		if item.Kind() != reflect.Struct {
			panic("item' type is not struct")
		}

		v := reflect.Indirect(item)
		for j := 0; j < v.NumField(); j++ {
			cd := fmt.Sprintf("%v", condition)
			cds := strings.Split(cd, "=")
			key := v.Type().Field(j).Name
			condKey := cds[0]

			if key == condKey {
				val := v.Field(j).Interface()
				condVal := cds[1]

				if fmt.Sprintf("%v", val) == condVal {
					return item.Interface(), nil
				}
			}
		}
	}

	return nil, fmt.Errorf("cannot find data")
}
