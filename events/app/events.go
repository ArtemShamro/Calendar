package app

import (
	"context"

	"github.com/ArtemShamro/Calendar/events/domain"
)

type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo : repo,
	}
}

func (s *Service) CreateNewEvent(ctx context.Context, hotel domain.Event) (*domain.Event, error) {
	return s.repo.CreateNewEvent(ctx, hotel)
}

func (s *Service) EditEvent(ctx context.Context, hotel *domain.Event) (*domain.Event, error) {
	return s.repo.EditEvent(ctx, hotel)
}

func (s *Service) GetEventsList(ctx context.Context) (*[]domain.Event, error) {
	return s.repo.GetEventsList(ctx)
}

func (s *Service) DeleteEvent(ctx context.Context, id int) (*domain.Event, error) {
	return s.repo.DeleteEvent(ctx, id)
}