package service

import (
	"errors"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	model "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"strings"
)

type LinkService struct {
	repo repository.Link
}

func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}

func (s LinkService) Create(link model.Link) (uuid.UUID, error) {
	var descriptionP *string
	if link.Description != nil {
		description := *link.Description
		description = strings.TrimSpace(description)
		descriptionP = &description
	}
	trimmedRef := strings.TrimSpace(link.Ref)
	if trimmedRef == "" {
		return "", errors.New("link's ref should contain url")
	}
	return s.repo.Create(repoModel.Link{
		UserId:      link.UserId,
		Ref:         trimmedRef,
		Description: descriptionP,
		Categories:  link.Categories,
	})
}

func (s LinkService) GetAll(userId uuid.UUID) ([]model.Link, error) {
	repoLinks, err := s.repo.GetAll(userId)
	if err != nil {
		return nil, err
	}

	links := make([]model.Link, len(repoLinks))
	for i, repoLink := range repoLinks {
		links[i] = model.Link{
			Id:          repoLink.Id,
			UserId:      repoLink.UserId,
			Ref:         repoLink.Ref,
			Description: repoLink.Description,
			Categories:  repoLink.Categories,
		}
	}

	return links, nil
}

func (s LinkService) GetById(userId, linkId uuid.UUID) (model.Link, error) {
	repoLink, err := s.repo.GetById(userId, linkId)
	if err != nil {
		return model.Link{}, err
	}
	link := model.Link{
		Id:          repoLink.Id,
		UserId:      repoLink.UserId,
		Ref:         repoLink.Ref,
		Description: repoLink.Description,
		Categories:  repoLink.Categories,
	}
	return link, nil
}

func (s LinkService) DeleteById(userId, linkId uuid.UUID) error {
	return s.repo.DeleteById(userId, linkId)
}

func (s LinkService) Update(link model.Link) error {
	var descriptionP *string
	if link.Description != nil {
		description := *link.Description
		description = strings.TrimSpace(description)
		descriptionP = &description
	}

	trimmedRef := strings.TrimSpace(link.Ref)
	if link.Ref != "" && trimmedRef == "" { // check if in update link was not ref at all
		return errors.New("link's ref should contain url")
	}
	return s.repo.Update(repoModel.Link{
		Id:          link.Id,
		UserId:      link.UserId,
		Ref:         trimmedRef,
		Description: descriptionP,
		Categories:  link.Categories,
	})
}
