package q

import (
	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
)

func Get(tx *gorm.DB, param interface{}, res interface{}) error {
	GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)
	return tx.Take(res).Error
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
	GenSqlByParam(tx, param)
	return tx.Count(total).Error
}

func FindWithPaginate(tx *gorm.DB, param interface{}, res interface{}) error {
	GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)

	return tx.
		Scopes(Paginate(param)).
		Find(res).Error
}
