package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"juejinCollections/tool"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// 下滑线和首字母转换为大写
func getFileName(keys []string) []string {
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

func insertOrUpdate(tabelName string, mainKey, updateKey, otherKey []string, valList interface{}) (sql.Result, error) {

	sliceValue := reflect.Indirect(reflect.ValueOf(valList))
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("valList needs a slice")
	}

	sliceLen := sliceValue.Len()
	if sliceLen == 0 {
		return nil, nil
	}

	// 插入的字段名
	insertKey := append(mainKey, updateKey...)
	insertKey = append(insertKey, otherKey...)
	insertKeySql := strings.Join(insertKey, ",")

	// 插入的values
	valStrArr := []string{}
	fileKey := getFileName(insertKey)
	for i := 0; i < sliceLen; i++ {
		fieldItem := sliceValue.Index(i)
		var valArr = []string{}
		for _, keyItem := range fileKey {
			fieldValue := fieldItem.FieldByName(keyItem)
			valArr = append(valArr, tool.ValueToDbString(fieldValue.Interface()))
		}
		valStrArr = append(valStrArr, "("+strings.Join(valArr, ",")+")")
	}
	valStr := strings.Join(valStrArr, ",")

	// 冲突时写入字段
	updateSqlList := []string{}
	for _, v := range updateKey {
		updateSqlList = append(updateSqlList, fmt.Sprintf("%[1]s=excluded.%[1]s", v))
	}
	updateSql := strings.Join(updateSqlList, ",")

	// 冲突时更新条件
	whereSqlList := []string{}
	for _, v := range mainKey {
		whereSqlList = append(whereSqlList, fmt.Sprintf("excluded.%[1]s=%[2]s.%[1]s", v, tabelName))
	}
	whereKeySql := strings.Join(whereSqlList, ",")

	// 冲突键
	conflictKeySql := strings.Join(mainKey, ",")

	sql := fmt.Sprintf(
		"INSERT INTO `%s` (%s) VALUES %s ON CONFLICT (%s) DO UPDATE SET %s WHERE %s;",
		tabelName, insertKeySql, valStr, conflictKeySql, updateSql, whereKeySql,
	)

	return DbDal.Engine.Exec(sql)
}

// 添加收藏列表
func AddTags(list *[]TagModel) (sql.Result, error) {
	model := &TagModel{}
	tabelName := model.TableName()
	mainKey := []string{"id"}
	updateKey := []string{"tag_id", "tag_name", "color", "icon", "back_ground", "ctime", "mtime", "status", "creator_id", "user_name", "post_article_count", "concern_user_count", "isfollowed", "is_has_in", "update_time"}
	otherKey := []string{"create_time"}
	return insertOrUpdate(tabelName, mainKey, updateKey, otherKey, list)
}
