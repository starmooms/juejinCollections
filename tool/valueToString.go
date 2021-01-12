package tool

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// /reflect/tostring_test.go
// xorm.io/xorm/convert.go

// 转换为字符
func ValueToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}

	rv := reflect.ValueOf(val)
	k := rv.Kind()

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	default:
		return ""
	}
}

func toDbString(val string) string {
	if val == "" {
		return "NULL"
	}
	return fmt.Sprintf("\"%s\"", val)
}

// 转换为sql字符
func ValueToDbString(val interface{}) string {
	switch s := val.(type) {
	case string:
		return toDbString(s)
	case []byte:
		return toDbString(string(s))
	case time.Time:
		if s.IsZero() {
			s = time.Now()
		}
		timeStr := s.Format("2006-01-02 15:04:05")
		return toDbString(timeStr)
	case bool:
		if s {
			return "1"
		}
		return "0"
	case uint, uint8, uint16, uint32, uint64:
		return ValueToString(s)
	case int, int8, int16, int32, int64:
		return ValueToString(s)
	case float32, float64:
		return ValueToString(s)
	case nil:
		return toDbString("")
	default:
		str := ValueToString(val)
		return str
	}
}

// 转换为sql要写入的interface参数
func ValueToDbInterface(val interface{}) interface{} {
	switch s := val.(type) {
	case string:
		if s == "" {
			return nil
		}
	case time.Time:
		if s.IsZero() {
			return nil
		}
		return s.Format("2006-01-02 15:04:05")
	}
	return val
}
