package q

import (
	"github.com/gogf/gf/util/gconv"
	"gorm.io/gorm"
)

func FindOne(tx *gorm.DB, param interface{}, res interface{}) error {
	GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)
	return tx.Take(res).Error
}

func CreateOne(tx *gorm.DB, m interface{}, param interface{}) error {
	gconv.Struct(param, m)
	return tx.Create(m).Error
}

func PatchOne(tx *gorm.DB, m interface{}, param interface{}) error {
	GenSqlByParam(tx, param)
	err := tx.Take(m).Error
	if err != nil {
		return err
	}
	gconv.Struct(param, m)
	return tx.Save(m).Error
}

func DeleteOne(tx *gorm.DB, m interface{}, param interface{}) error {
	GenSqlByParam(tx, param)
	err := tx.Take(m).Error
	if err != nil {
		return err
	}
	return tx.Delete(m).Error
}

func Find(tx *gorm.DB, param interface{}, res interface{}, total *int64) error {
	countSql := GenSqlByParam(tx, param)
	if result := countSql.Count(total); result.Error != nil {
		return result.Error
	}

	GenSqlByParam(tx, param)
	GenSqlByRes(tx, res)

	return tx.
		Scopes(Paginate(param)).
		Find(res).Error
}
