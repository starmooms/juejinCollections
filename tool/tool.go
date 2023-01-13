package tool

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
)

func rvSetTime(itemVal reflect.Value, timeRv reflect.Value, fileKey []string) {
	for _, key := range fileKey {
		fileItem := itemVal.FieldByName(key)
		fileItem.Set(timeRv)
	}
}

/** 将结构体中特定时间字段 设置为当前时间 */
func SetTimeFile(fileKey []string, list interface{}) error {
	listRv := reflect.Indirect(reflect.ValueOf(list))

	now := time.Now()
	nowRv := reflect.ValueOf(now)
	if listRv.Kind() == reflect.Slice {
		len := listRv.Len()
		if len == 0 {
			return nil
		}
		for i := 0; i < len; i++ {
			itemVal := reflect.Indirect(listRv.Index(i))
			rvSetTime(itemVal, nowRv, fileKey)
		}
	} else {
		rvSetTime(listRv, nowRv, fileKey)
	}
	return nil
}

// func SetTimeFile(fileKey []string, list interface{}) error {
// 	sliceValue := reflect.Indirect(reflect.ValueOf(list).Elem())
// 	if sliceValue.Kind() != reflect.Slice {
// 		return errors.New("list needs a slice")
// 	}

// 	sliceLen := sliceValue.Len()
// 	if sliceLen == 0 {
// 		return nil
// 	}

// 	now := time.Now()
// 	nowRv := reflect.ValueOf(now)
// 	for i := 0; i < sliceLen; i++ {
// 		itemVal := sliceValue.Index(i)
// 		for _, key := range fileKey {
// 			fileItem := itemVal.FieldByName(key)
// 			fileItem.Set(nowRv)
// 		}
// 	}
// 	return nil
// }

// 下滑线和首字母转换为大写
func GetFileName(keys []string) []string {
	fileKey := []string{}
	toLowerNum := 'a' - 'A'
	for _, val := range keys {
		newstr := make([]rune, 0)
		name := strings.ToLower(val)
		parts := strings.Split(name, "_")
		for _, partItem := range parts {
			for i, r := range partItem {
				if i == 0 && 'a' <= r && r <= 'z' {
					r -= toLowerNum
				}
				newstr = append(newstr, r)
			}
		}
		fileKey = append(fileKey, string(newstr))
	}
	return fileKey
}

/** 获取结构体中 sql字段和field 的相关字典 */
func GetSqlKeyMap(rv reflect.Value) (sqlKeyMap map[string]string, err error) {
	rt := rv.Type()
	fieldNum := rt.NumField()
	// fieldKeyMap = map[string]string{}
	sqlKeyMap = map[string]string{}

	for i := 0; i < fieldNum; i++ {
		field := rt.Field(i)
		sqlKey := field.Tag.Get("json")
		ormTagStr := field.Tag.Get("xorm")
		if ormTagStr != "" {
			r, err := regexp.Compile("'\\w+'")
			if err != nil {
				return nil, errors.Wrap(err, "Reg Error")
			}
			regResult := r.FindStringSubmatch(ormTagStr)
			if len(regResult) == 2 {
				sqlKey = regResult[1]
			}
		}

		fieldName := field.Name
		// fieldKeyMap[fieldName] = sqlKey
		sqlKeyMap[sqlKey] = fieldName
	}
	return
}

/** 限制字符传长度 */
func LimtStr(str string, limt int) string {
	strLen := len(str)
	if strLen > limt {
		str = fmt.Sprintf("%s ...+%d", string(str[0:limt]), strLen-limt)
	}
	return str
}

func RegExpReplace(expr string, src string, repl string) string {
	reg := regexp.MustCompile(expr)
	return reg.ReplaceAllString(src, repl)
}
