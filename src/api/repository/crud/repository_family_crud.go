package crud

import (
	"errors"

	"api/models"
	"api/utils/channels"

	"github.com/jinzhu/gorm"
)

type repositoryFamilyCRUD struct {
	db *gorm.DB
}

func NewRepositoryFamilyCRUD(db *gorm.DB) *repositoryFamilyCRUD {
	return &repositoryFamilyCRUD{db}
}

func (r *repositoryFamilyCRUD) Save(family models.Family) (models.Family, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Family{}).Create(&family).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return family, nil
	}
	return models.Family{}, err
}

func (r *repositoryFamilyCRUD) FindAll() ([]models.Family, error) {
	var err error
	done := make(chan bool)
	families := []models.Family{}
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Family{}).Find(&families).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return families, nil
	}
	return nil, err
}

func (r *repositoryFamilyCRUD) FindById(familyId uint32) (models.Family, error) {
	var err error
	family := models.Family{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Family{}).Where("family_Id = ?", familyId).Take(&family).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return family, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.Family{}, errors.New("family not found")
	}
	return models.Family{}, err
}

func (r *repositoryFamilyCRUD) Update(familyId uint32, family models.Family) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Family{}).Where("family_id = ?", familyId).Take(&models.Family{}).UpdateColumns(
			map[string]interface{}{
				"Name":        family.Name,
				"MaxPerson":   family.MaxPerson,
				"CurrentSize": family.CurrentSize,
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *repositoryFamilyCRUD) Delete(uid uint32) (int64, error) {

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Family{}).Where("family_id = ?", uid).Take(&models.Family{}).Delete(&models.Family{})
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
