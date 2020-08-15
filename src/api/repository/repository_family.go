package repository

import "api/models"

type FamilyRepository interface {
	Save(models.Family) (models.Family, error)
	FindAll() ([]models.Family, error)
	FindById(uint32) (models.Family, error)
	Update(uint32, models.Family) (int64, error)
	Delete(uint32) (int64, error)
}
