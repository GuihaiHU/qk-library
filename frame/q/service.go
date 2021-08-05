package q

import (
	"reflect"

	"github.com/gogf/gf/util/gconv"
	"github.com/iWinston/qk-library/frame/qfield"
	"github.com/iWinston/qk-library/frame/qservice"
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
	id := param.(qfield.IdParam).GetId()
	if len(tx.Statement.Preloads) == 0 {
		return tx.Take(res, id).Error
	} else {
		// 这里是为了防止select里不包含外键字段,所以select设置为空，Take启动智能select所有字段
		preloadTx := tx.Session(&gorm.Session{})
		preloadTx.Statement.Selects = []string{}
		if err := preloadTx.Take(preloadTx.Statement.Model, id).Error; err != nil {
			return err
		}
		if err := gconv.StructDeep(preloadTx.Statement.Model, res); err != nil {
			return err
		}

		// TODO 判断是否有select别名字段或者外键字段，当有的时候才select
		result := make(map[string]interface{})
		if err := tx.Where(id).Scan(&result).Error; err != nil {
			return err
		}
		return gconv.StructDeep(result, res)
	}
}

func Post(tx *gorm.DB, m interface{}, param interface{}) (err error) {
	if err = gconv.Struct(param, m); err != nil {
		return err
	}
	if err = tx.Create(m).Error; err != nil {
		return err
	}
	qservice.ReqContext.SetAfterModelByTx(tx, m)
	qservice.ReqContext.SetRowsAffectedByTx(tx)
	return
}

func Patch(tx *gorm.DB, m interface{}, param interface{}) (err error) {
	GenSqlByParam(tx, param)
	id := param.(qfield.IdParam).GetId()
	if err = tx.Take(m, id).Error; err != nil {
		return
	}
	qservice.ReqContext.SetBeforeModelByTx(tx, m)
	if err = gconv.Struct(param, m); err != nil {
		return
	}
	if err = tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(m).Error; err != nil {
		return
	}
	qservice.ReqContext.SetAfterModelByTx(tx, m)
	qservice.ReqContext.SetRowsAffectedByTx(tx)
	return
}

func Delete(tx *gorm.DB, m interface{}, param interface{}) (err error) {
	GenSqlByParam(tx, param)
	id := param.(qfield.IdParam).GetId()
	if err = tx.Take(m, id).Error; err != nil {
		return err
	}
	qservice.ReqContext.SetBeforeModelByTx(tx, m)
	if err = tx.Delete(m, id).Error; err != nil {
		return err
	}
	qservice.ReqContext.SetRowsAffectedByTx(tx)
	return
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
		// 这里是为了防止select里不包含外键字段,所以select设置为空，Find启动智能select所有字段
		arrType := reflect.SliceOf(reflect.TypeOf(tx.Statement.Model).Elem())
		arr := reflect.New(arrType).Interface()
		preloadTx := tx.Session(&gorm.Session{})
		preloadTx.Statement.Selects = []string{}
		if err := preloadTx.Find(arr).Error; err != nil {
			return err
		}
		if err := gconv.StructDeep(arr, res); err != nil {
			return err
		}

		// TODO 判断是否有select别名字段或者外键字段，当有的时候才select
		var results []map[string]interface{}
		if err := tx.Scan(&results).Error; err != nil {
			return err
		}
		return gconv.StructsDeep(results, res)
	}
}
