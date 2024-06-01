package array

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"reflect"
)

func ColumnAny[T any, V any](slice []T, column string) []V {
	// 创建一个空 map 用于存储去重后的值
	uniqueMap := make(map[interface{}]bool)

	s := make([]V, 0)
	for _, item := range slice {
		v := reflect.ValueOf(item)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		field := v.FieldByName(column)
		if field.IsValid() {
			uniqueMap[field.Interface().(V)] = true
			//s = append(s, field.Interface().(V))
		}
	}

	for key := range uniqueMap {
		s = append(s, key.(V))
	}

	return s
}

func ContainsValue[T comparable](list []T, target T) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}

func ArrayUnique(arr interface{}) interface{} {
	// 创建一个空 map 用于存储去重后的值
	uniqueMap := make(map[interface{}]bool)

	// 获取数组的类型和值
	arrType := reflect.TypeOf(arr)
	arrValue := reflect.ValueOf(arr)

	// 判断数组类型是否为切片
	if arrType.Kind() != reflect.Slice {
		return nil
	}

	// 遍历数组元素并存入 map 中
	for i := 0; i < arrValue.Len(); i++ {
		elemValue := arrValue.Index(i).Interface()
		uniqueMap[elemValue] = true
	}

	// 将 map 中的值存入新的切片中
	resultValue := reflect.MakeSlice(arrType, 0, len(uniqueMap))
	for key := range uniqueMap {
		resultValue = reflect.Append(resultValue, reflect.ValueOf(key))
	}

	return resultValue.Interface()
}

// DeleteSlice deletes the specified value from the slice.
func DeleteSlice(slice, value interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice {
		panic("deleteValue: slice parameter must be a slice")
	}

	/*
		valueValue := reflect.ValueOf(value)
		elemType := sliceValue.Type().Elem()
		if valueValue.Type() != elemType {
			panic("deleteValue: value parameter must have the same type as the elements in the slice")
		}
	*/

	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i)
		if reflect.DeepEqual(elem.Interface(), value) {
			// Shift the elements after the deletion point
			if i < sliceValue.Len()-1 {
				reflect.Copy(sliceValue.Slice(i, sliceValue.Len()), sliceValue.Slice(i+1, sliceValue.Len()))
			}

			// Truncate the slice by one element
			newLen := sliceValue.Len() - 1
			newSlice := reflect.MakeSlice(sliceValue.Type(), newLen, newLen)
			reflect.Copy(newSlice, sliceValue.Slice(0, newLen))
			return newSlice.Interface()
		}
	}

	// Value not found, return the original slice
	return slice
}

func ArraySearchStr(arr []string, searchValue string) int {
	for i, v := range arr {
		if v == searchValue {
			return i
		}
	}

	return -1
}

func ArraySearch(arr []uint, searchValue uint) int {
	for i, v := range arr {
		if v == searchValue {
			return i
		}
	}

	return -1
}

func Disjoint(nums1 []uint64, nums2 []uint64) bool {
	// 创建一个map用于存储数组1中的元素
	nums1Map := make(map[uint64]bool)

	// 将数组1中的元素存储到map中
	for _, num := range nums1 {
		nums1Map[num] = true
	}

	// 遍历数组2，如果有元素存在于数组1的map中，则说明存在交集
	for _, num := range nums2 {
		if nums1Map[num] {
			return false
		}
	}

	return true
}

// InArray in_array()
// haystack supported types: slice, array or map
func InArray(haystack interface{}, needle interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}

func Column(slice interface{}, column string) []interface{} {
	column = gstr.CaseCamel(column)

	// 创建一个空 map 用于存储去重后的值
	uniqueMap := make(map[interface{}]bool)
	s := make([]interface{}, 0)

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			if item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			field := item.FieldByName(column)
			if field.IsValid() {
				uniqueMap[field.Interface()] = true
				//s = append(s, field.Interface())
			}
		}

		for key := range uniqueMap {
			s = append(s, key)
		}

	default:
		panic("expected a slice")
	}

	return s
}

// DeleteEmpty 删除切片中的空值并返回新切片
func DeleteEmpty(slice []interface{}) []interface{} {
	var result []interface{}
	for _, s := range slice {
		if !g.IsEmpty(s) {
			result = append(result, s)
		}
	}

	return result
}
