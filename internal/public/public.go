package public

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func GetCurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

func ParamValidate(page int, pageSize int) (bool, error) {
	if pageSize > 5000 || pageSize < 10 {
		return false, errors.New("请求数据量非法，一次请求数据范围10-5000")
	}
	if page < 1 {
		return false, errors.New("页码非法")
	}
	return true, nil
}

func SJoin(sep string, args ...string) string {
	var array []string
	for _, v := range args {
		array = append(array, v)
	}
	return strings.Join(array, sep)
}

func SJoin2(sep string, args []string) string {
	var array []string
	for _, v := range args {
		array = append(array, v)
	}
	return strings.Join(array, sep)
}

func GetOrgX(src string) string {
	return strings.Replace(src, "MSP", "", -1)
}

func CheckStrs(src string, args ...string) bool {
	for _, subs := range args {
		if strings.Contains(src, subs) {
			continue
		} else {
			return false
		}
	}
	return true
}
