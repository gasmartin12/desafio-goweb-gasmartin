package tickets

import (
	"context"

	"github.com/gasmartin12/desafio-goweb-gasmartin/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.repo.GetTicketByDestination(ctx, destination)
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	return s.repo.AverageDestination(ctx, destination)
}
