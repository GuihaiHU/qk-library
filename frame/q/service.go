package q

import (
	"reflect"

	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
)

// Find find records that match given conditions
func Find(tx *gorm.DB, dest interface{}, conds ...interface{}) error {
	if len(tx.Statement.Preloads) == 0 {
		return tx.Find(dest, conds).Error
	} else {
		arrType := reflect.SliceOf(reflect.TypeOf(tx.Statement.Model).Elem())
		arr := reflect.New(arrType).Interface()
		if err := tx.Find(arr, conds).Scan(dest).Error; err != nil {
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
		if err := tx.Take(tx.Statement.Model).Scan(res).Error; err != nil {
			return err
		}
		return gconv.StructDeep(tx.Statement.Model, res)
	}
}

func Post(tx *gorm.DB, m interface{}, param interface{}) error {
	gconv.Struct(param, m)
	return tx.Create(m).Error
}

func Patch(tx *gorm.DB, m interface{}, param interface{}) error {
	GenSqlByParam(tx, param)
	err := tx.Take(m).Error
	if err != nil {
		return err
	}
	gconv.Struct(param, m)
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(m).Error
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
	if err := Count(tx, param, total); err != nil {
		return err
	}
	return FindWithPaginate(tx, param, res)
}

func Count(tx *gorm.DB, param interface{}, total *int64) error {
	tx = tx.Session(&gorm.Session{})
	GenSqlByParam(tx, param)
	return tx.Count(total).Error
}

func FindWithPaginate(tx *gorm.DB, param interface{}, res interface{}) error {
	GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)

	tx.Scopes(Paginate(param))

	if len(tx.Statement.Preloads) == 0 {
		return tx.Find(res).Error
	} else {
		arrType := reflect.SliceOf(reflect.TypeOf(tx.Statement.Model).Elem())
		arr := reflect.New(arrType).Interface()
		if err := tx.Find(arr).Scan(res).Error; err != nil {
			return err
		}
		return gconv.ScanDeep(arr, res)
	}
}
