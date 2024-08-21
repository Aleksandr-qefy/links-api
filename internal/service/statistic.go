package service

import (
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	model "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
)

type StatisticService struct {
	repo repository.Statistic
}

func NewStatisticService(repo repository.Statistic) *StatisticService {
	return &StatisticService{repo: repo}
}

func (s StatisticService) Create(category model.Statistic) (uuid.UUID, error) {
	return s.repo.Create(repoModel.Statistic{
		Id:        category.Id,
		UserId:    category.UserId,
		CreatedAt: category.CreatedAt,
		Activity:  category.Activity,
		Comment:   category.Comment,
	})
}

func (s StatisticService) GetAll(userId uuid.UUID) ([]model.Statistic, error) {
	repoStatistics, err := s.repo.GetAll(userId)
	if err != nil {
		return nil, err
	}

	statistics := make([]model.Statistic, len(repoStatistics))
	for i, repoStatistic := range repoStatistics {
		statistics[i] = model.Statistic{
			Id:        repoStatistic.Id,
			UserId:    repoStatistic.UserId,
			CreatedAt: repoStatistic.CreatedAt,
			Activity:  repoStatistic.Activity,
			Comment:   repoStatistic.Comment,
		}
	}

	return statistics, nil
}
