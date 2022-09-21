package account

import (
	"context"
	"log"

	"github.com/go-kit/kit/log"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func CreateUser(ctx context.Context, email string, password string) (string, error) {

}

func GetUser(ctx context.Context, id string) (string, error) {

}
