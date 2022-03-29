package easycondition

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	cList = []string{"=", ">", "<"}
)

const (
	biggest  = ">"
	smallest = "<"
	equal    = "="
)

func First(source interface{}, condition string) (interface{}, error) {
	s := reflect.Indirect(reflect.ValueOf(source))
	if s.Kind() != reflect.Slice {
		panic("source' type is not slice")
	}

	c := reflect.Indirect(reflect.ValueOf(condition))
	if c.Kind() != reflect.String || c.IsZero() == true {
		panic("condition' type is not string or const value empty")
	}

	crt := criterion(condition)

	for i := 0; i < s.Len(); i++ {
		item := s.Index(i)
		if item.Kind() != reflect.Struct {
			panic("item' type is not struct")
		}

		v := reflect.Indirect(item)
		for j := 0; j < v.NumField(); j++ {
			cds := strings.Split(fmt.Sprintf("%v", condition), crt)
			condKey := cds[0]
			condVal := cds[1]

			key := v.Type().Field(j).Name
			val := v.Field(j).Interface()

			if key == condKey {
				switch crt {
				case equal:
					if fmt.Sprintf("%v", val) == condVal {
						return item.Interface(), nil
					}
					break
				case biggest:
					v, _ := strconv.Atoi(fmt.Sprintf("%v", val))
					cv, _ := strconv.Atoi(condVal)
					if v > cv {
						return item.Interface(), nil
					}
					break
				case smallest:
					v, _ := strconv.Atoi(fmt.Sprintf("%v", val))
					cv, _ := strconv.Atoi(condVal)
					if v < cv {
						return item.Interface(), nil
					}
					break
				}
			}
		}
	}

	return nil, fmt.Errorf("cannot find data")
}

func criterion(c string) string {
	for _, item := range cList {
		if strings.Contains(c, item) {
			return item
		}
	}
	return ""
}
