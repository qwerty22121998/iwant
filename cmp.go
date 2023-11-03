package iwant

import "reflect"

func IsZero(v interface{}) bool {
	return v == nil || reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
