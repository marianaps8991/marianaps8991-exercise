package crud

import (
	"errors"

	"api/security"

	"api/models"
	"api/utils/channels"

	"github.com/jinzhu/gorm"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		if user.FamilyID != 0 {
			fam := models.Family{}
			err := r.db.Debug().Model(&models.Family{}).Where("family_Id = ?", user.FamilyID).Take(&fam).Error
			if err != nil {
				ch <- false
				return
			}
			user.Family = fam
		}
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		if len(users) > 0 {
			for i, _ := range users {
				err = r.db.Debug().Model(&models.Family{}).Where("family_id= ?", users[i].FamilyID).Take(&users[i].Family).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (r *repositoryUsersCRUD) FindById(id uint32) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("user_id = ?", id).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		err = r.db.Debug().Model(&models.Family{}).Where("family_id= ?", user.FamilyID).Take(&user.Family).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) Update(uid uint32, user models.User) (int64, error) {

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		family := models.Family{}
		rs = r.db.Debug().Model(&models.User{}).Where("family_id = ?", user.FamilyID).Take(&family)
		password, err := security.Hash(user.Password)
		if err != nil {
			ch <- false
			return
		}

		rs = r.db.Debug().Model(&models.User{}).Where("user_id = ?", uid).Take(&models.User{}).Update(
			map[string]interface{}{
				"Name":     user.Name,
				"Age":      user.Age,
				"Username": user.Username,
				"Password": password,
				"FamilyId": user.FamilyID,
				"Family":   family,
				"Role":     user.Role,
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

func (r *repositoryUsersCRUD) Delete(uid uint32) (int64, error) {

	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("user_id = ?", uid).Take(&models.User{}).Delete(&models.User{})
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
