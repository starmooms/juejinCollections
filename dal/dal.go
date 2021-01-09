package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"juejinCollections/model"
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
	whereKeySql += " and excluded.mTime is NULL"

	// 冲突键
	conflictKeySql := strings.Join(mainKey, ",")

	sql := fmt.Sprintf(
		"INSERT INTO `%s` (%s) VALUES %s ON CONFLICT (%s) DO UPDATE SET %s WHERE %s;",
		tabelName, insertKeySql, valStr, conflictKeySql, updateSql, whereKeySql,
	)

	return DbDal.Engine.Exec(sql)
}

// 添加收藏列表
func AddTags(list *[]model.TagModel) (sql.Result, error) {
	tagModel := &model.TagModel{}
	tabelName := tagModel.TableName()
	mainKey := []string{"id"}
	updateKey := []string{"tag_id", "tag_name", "color", "icon", "back_ground", "ctime", "mtime", "status", "creator_id", "user_name", "post_article_count", "concern_user_count", "isfollowed", "is_has_in", "update_time"}
	otherKey := []string{"create_time"}
	// s := "\"计算机基础'\""
	// // s = "\"webpack\" or tag_name=\"监控\""

	// users := []model.TagModel{}
	// err2 := DbDal.Engine.Where("tag_name= ?", s).Find(&users)
	// r, err3 := DbDal.Engine.Exec(fmt.Sprintf("SELECT `id`, `tag_id`, `tag_name`, `color`, `icon`, `back_ground`, `ctime`, `mtime`, `status`, `creator_id`, `user_name`, `post_article_count`, `concern_user_count`, `isfollowed`, `is_has_in`, `create_time`, `update_time` FROM `tags` WHERE (tag_name= %s )", s))
	// fmt.Println("err2", err2, &users)
	// fmt.Println("err3", err3, r)

	return insertOrUpdate(tabelName, mainKey, updateKey, otherKey, list)
}
