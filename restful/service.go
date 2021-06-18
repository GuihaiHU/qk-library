package restful

import (
	"gorm.io/gorm"
)

func TakeOne(sql *gorm.DB, id uint, res interface{}) error {
	if result := sql.Take(res, id); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func CreateOne(sql *gorm.DB, modelToCreate interface{}) error {
	if result := sql.Create(modelToCreate); result.Error != nil {
		return result.Error
	}
	return nil
}

func PatchOne(sql *gorm.DB, id uint, modelToPatch interface{}) error {
	if result := sql.Save(modelToPatch); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteOne(sql *gorm.DB, id uint, modelToDelete interface{}) error {
	if result := sql.Delete(modelToDelete); result.Error != nil {
		return result.Error
	}
	return nil
}

func Find(sql *gorm.DB, param interface{}, res interface{}) error {
	if result := sql.
		Scopes(Paginate(param)).
		Find(res); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
