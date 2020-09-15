package repository

import "gome/api/models"

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindByID(uint32) (models.Post, error)
	Update(uint32, models.Post) (int64, error)
	Delete(uint32, uint64) (int64, error)
}