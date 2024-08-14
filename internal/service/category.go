package service

import (
	"errors"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	model "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"strings"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s CategoryService) Create(category model.Category) (uuid.UUID, error) {
	name := strings.TrimSpace(category.Name)
	if name == "" {
		return "", errors.New("category name should contain one char at least")
	}

	return s.repo.Create(repoModel.Category{
		Id:     category.Id,
		UserId: category.UserId,
		Name:   name,
	})
}

func (s CategoryService) GetAll(userId uuid.UUID) ([]model.Category, error) {
	repoCategories, err := s.repo.GetAll(userId)
	if err != nil {
		return nil, err
	}

	categories := make([]model.Category, len(repoCategories))
	for i, repoCategory := range repoCategories {
		categories[i] = model.Category{
			Id:     repoCategory.Id,
			UserId: repoCategory.UserId,
			Name:   repoCategory.Name,
		}
	}

	return categories, nil
}

func (s CategoryService) GetById(userId, categoryId uuid.UUID) (model.Category, error) {
	repoCategory, err := s.repo.GetById(userId, categoryId)
	if err != nil {
		return model.Category{}, err
	}
	category := model.Category{
		Id:     repoCategory.Id,
		UserId: repoCategory.UserId,
		Name:   repoCategory.Name,
	}
	return category, nil
}

func (s CategoryService) DeleteById(userId, categoryId uuid.UUID) error {
	return s.repo.DeleteById(userId, categoryId)
}

func (s CategoryService) Update(category model.Category) error {
	name := strings.TrimSpace(category.Name)
	if name == "" {
		return errors.New("category name should contain one char at least")
	}
	return s.repo.Update(repoModel.Category{
		Id:     category.Id,
		UserId: category.UserId,
		Name:   name,
	})
}
