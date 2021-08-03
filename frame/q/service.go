package q

import (
	"reflect"

	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
)

// Find find records that match given conditions
func Find(tx *gorm.DB, dest interface{}, conds ...interface{}) error {
	if len(tx.Statement.Preloads) == 0 {
		return tx.Find(dest, conds...).Error
	} else {
		arrType := reflect.SliceOf(reflect.TypeOf(tx.Statement.Model).Elem())
		arr := reflect.New(arrType).Interface()
		if err := tx.Find(arr, conds...).Scan(dest).Error; err != nil {
			return err
		}
		return gconv.ScanDeep(arr, dest)
	}
}

func Get(tx *gorm.DB, param interface{}, res interface{}) error {
	GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)
	if len(tx.Statement.Preloads) == 0 {
		return tx.Take(res).Error
	} else {
		// TODO 判断是否有select别名字段或者外键字段，当有的时候才select
		result := make(map[string]interface{})
		if err := tx.Scan(&result).Error; err != nil {
			return err
		}
		if err := gconv.StructDeep(result, res); err != nil {
			return err
		}

		// 这里是为了防止select里不包含外键字段,所以select设置为*
		if err := tx.Select("*").Take(tx.Statement.Model).Error; err != nil {
			return err
		}
		return gconv.StructDeep(tx.Statement.Model, res)
	}
}

func Post(tx *gorm.DB, m interface{}, param interface{}) error {
	if err := gconv.Struct(param, m); err != nil {
		return err
	}
	return tx.Create(m).Error
}

func Patch(tx *gorm.DB, m interface{}, param interface{}) error {
	GenSqlByParam(tx, param)
	err := tx.Take(m).Error
	if err != nil {
		return err
	}
	if err := gconv.Struct(param, m); err != nil {
		return err
	}
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(m).Error
}

func Delete(tx *gorm.DB, m interface{}, param interface{}) error {
	GenSqlByParam(tx, param)
	err := tx.Take(m).Error
	if err != nil {
		return err
	}
	return tx.Delete(m).Error
}

func List(tx *gorm.DB, param interface{}, res interface{}, total *int64) error {
	GenSqlByParam(tx, param)
	if err := Count(tx, param, total); err != nil {
		return err
	}
	return FindWithPaginate(tx, param, res)
}

func Count(tx *gorm.DB, param interface{}, total *int64) error {
	tx = tx.Session(&gorm.Session{})
	return tx.Count(total).Error
}

func FindWithPaginate(tx *gorm.DB, param interface{}, res interface{}) error {
	// GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)

	tx.Scopes(Paginate(param))

	if len(tx.Statement.Preloads) == 0 {
		return tx.Find(res).Error
	} else {
		// TODO 判断是否有select别名字段或者外键字段，当有的时候才select
		var results []map[string]interface{}
		if err := tx.Scan(&results).Error; err != nil {
			return err
		}
		if err := gconv.StructsDeep(results, res); err != nil {
			return err
		}

		arrType := reflect.SliceOf(reflect.TypeOf(tx.Statement.Model).Elem())
		arr := reflect.New(arrType).Interface()
		// 这里是为了防止select里不包含外键字段,所以select设置为*
		if err := tx.Select("*").Find(arr).Error; err != nil {
			return err
		}

		return gconv.StructsDeep(arr, res)
	}
}
