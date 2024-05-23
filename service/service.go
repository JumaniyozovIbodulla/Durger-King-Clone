package service

import (
	"durger/pkg/logger"
	"durger/storage"
)

type IServiceManager interface {
	User() userService
}

type Service struct {
	userService userService
	logger      logger.ILogger
}

func New(storage storage.IStorage, logger logger.ILogger) Service {
	services := Service{}

	services.userService = NewUserService(storage, logger)
	return services
}

func (s Service) User() userService {
	return s.userService
}
