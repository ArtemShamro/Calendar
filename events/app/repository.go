package app

import (
	"context"

	"github.com/ArtemShamro/Calendar/events/domain"
)

type Repository interface {
	CreateNewEvent(ctx context.Context, hotel domain.Event) (*domain.Event, error)
	GetEventsList(ctx context.Context) (*[]domain.Event, error)
	EditEvent(ctx context.Context, hotel *domain.Event) (*domain.Event, error)
	DeleteEvent(ctx context.Context, id int) (*domain.Event, error)
}