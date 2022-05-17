package repository

import (
	"context"
	"errors"
	"microservices/ticket-process/pkg/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Repository interface {
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
	FindById(context.Context, uuid.UUID) (*model.Ticket, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) Create(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {
	sql := "insert into ticket (id, order_id, description, email, status) values ($1, $2, $3, $4, $5) returning id"

	var id uuid.UUID
	err := r.databaseManager.QueryRow(ctx, sql,
		ticket.Id, ticket.OrderId, ticket.Description, ticket.Email, ticket.Status).Scan(&id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r RepositoryImpl) FindById(ctx context.Context, ticketId uuid.UUID) (*model.Ticket, error) {
	sql := "select t.id, t.order_id, t.description, t.email, t.status from ticket t where t.id = $1"

	var id uuid.UUID
	var order_id uuid.UUID
	var description string
	var email string
	var status string

	err := r.databaseManager.QueryRow(ctx, sql, ticketId).Scan(&id, &order_id, &description, &email, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &model.Ticket{
		Id:          id,
		OrderId:     order_id,
		Description: description,
		Email:       email,
		Status:      status,
	}, nil
}
