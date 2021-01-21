package dal

import (
	"database/sql"
	"fmt"
	"juejinCollections/model"
	"juejinCollections/tool"
	"reflect"
	"strings"

	"github.com/cockroachdb/errors"
	_ "github.com/mattn/go-sqlite3"
)

/** 插入或更新 */
func insertOrUpdate(mainField, noUpdateField []string, valList interface{}, whereSqlFn func(tbName string) string) (sql.Result, error) {
	sliceValue := reflect.Indirect(reflect.ValueOf(valList))
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("valList needs a slice")
	}

	sliceLen := sliceValue.Len()
	if sliceLen == 0 {
		return nil, nil
	}

	// 获取表名 和 key
	sliceFirst := reflect.Indirect(sliceValue.Index(0))
	tableNameRv := sliceFirst.Addr().MethodByName("TableName").Call([]reflect.Value{})[0]
	tabelName := tableNameRv.String()
	sqlKeyMap, err := tool.GetSqlKeyMap(sliceFirst)
	if err != nil {
		return nil, err
	}

	// 生成Key
	insertKey := []string{}
	updateKey := []string{}
	mainKey := []string{}
	fileKeyArr := []string{}
	for sqlKey, fieldKey := range sqlKeyMap {
		insertKey = append(insertKey, sqlKey)
		fileKeyArr = append(fileKeyArr, fieldKey)
		isUpdate := true

		for _, mV := range mainField {
			if fieldKey == mV {
				mainKey = append(mainKey, sqlKey)
				isUpdate = false
				continue
			}
		}

		if isUpdate {
			for _, noV := range noUpdateField {
				if fieldKey == noV {
					isUpdate = false
					continue
				}
			}
			if isUpdate {
				updateKey = append(updateKey, sqlKey)
			}
		}
	}
	insertKeySql := strings.Join(insertKey, ",")

	// 插入的values
	var args []interface{}
	valStrArr := []string{}
	for i := 0; i < sliceLen; i++ {
		fieldItem := sliceValue.Index(i)
		if fieldItem.Kind() == reflect.Ptr {
			fieldItem = fieldItem.Elem()
		}
		var valArr = []string{}
		for _, keyItem := range fileKeyArr {
			fieldValue := fieldItem.FieldByName(keyItem)
			args = append(args, tool.ValueToDbInterface(fieldValue.Interface()))
			valArr = append(valArr, "?")
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
	whereKeySql := strings.Join(whereSqlList, " AND ")
	if whereSqlFn != nil {
		whereKeySql += " AND " + whereSqlFn(tabelName)
	}

	// 冲突键
	conflictKeySql := strings.Join(mainKey, ",")

	sql := fmt.Sprintf(
		"INSERT INTO `%s` (%s) VALUES %s ON CONFLICT (%s) DO UPDATE SET %s WHERE %s;",
		tabelName, insertKeySql, valStr, conflictKeySql, updateSql, whereKeySql,
	)

	var sqlArgs = []interface{}{sql}
	sqlArgs = append(sqlArgs, args...)
	return DbDal.Engine.Exec(sqlArgs...)
}

// 添加收藏列表
func AddTags(list *[]model.TagModel) (sql.Result, error) {
	main := []string{"Id"}
	noUpdate := []string{"CreateTime"}
	tool.SetTimeFile([]string{"CreateTime", "UpdateTime"}, list)
	return insertOrUpdate(main, noUpdate, list, nil)
}

func AddArticle(article *model.ArticleModel) (sql.Result, error) {
	tool.SetTimeFile([]string{"CreateTime", "UpdateTime"}, article)
	main := []string{"ArticleId"}
	noUpdate := []string{"CreateTime"}
	return insertOrUpdate(main, noUpdate, []*model.ArticleModel{article}, func(tbName string) string {
		return fmt.Sprintf("excluded.%[1]s!=%[2]s.%[1]s", "ctime", tbName)
	})
}
