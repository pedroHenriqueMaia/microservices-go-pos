package service

import (
	"context"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"

	"microservices/ticket-process/pkg/model"

	"github.com/pkg/errors"
)

type Service interface {
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
}

type ServiceImpl struct {
	messagePublisher rabbitmq.MessagePublisher
}

func NewServiceImpl(messagePublisher rabbitmq.MessagePublisher) Service {
	return &ServiceImpl{
		messagePublisher: messagePublisher,
	}
}

func (s ServiceImpl) Create(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {

	if err := s.messagePublisher.Publish("notification", ticket); err != nil {
		return nil, errors.Wrap(err, "fail to publish order")
	}

	coreLog.Logger.Infof("ticket publish. %v", ticket)

	return ticket, nil
}
