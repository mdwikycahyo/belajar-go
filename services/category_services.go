package services

import (
	"category-api/models"
	"category-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll(search string) ([]models.Category, error) {
	return s.repo.GetAll(search)
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}