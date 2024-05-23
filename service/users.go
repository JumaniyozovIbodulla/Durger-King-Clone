package service

import (
	"context"
	"durger/models"
	"durger/pkg/logger"
	"durger/storage"
)

type userService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewUserService(storage storage.IStorage, logger logger.ILogger) userService {
	return userService{
		storage: storage,
		logger:  logger,
	}
}

func (u userService) Create(ctx context.Context, user models.AddUser) (int64, error) {
	id, err := u.storage.UserStorage().Create(ctx, user)

	if err != nil {
		u.logger.Error("error with service create a user: ", logger.Error(err))
		return id, err
	}
	return id, nil
}

func (u userService) Delete(ctx context.Context, id int64) error {
	err := u.storage.UserStorage().Delete(ctx, id)

	if err != nil {
		u.logger.Error("error with service delete a user: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userService) GetUser(ctx context.Context, id int64) (models.User, error) {
	user, err := u.storage.UserStorage().GetUser(ctx, id)

	if err != nil {
		u.logger.Error("error with service get a user: ", logger.Error(err))
		return user, err
	}
	return user, nil
}

func (u userService) GetAll(ctx context.Context) (models.AllUsers, error) {
	users, err := u.storage.UserStorage().GetAll(ctx)

	if err != nil {
		u.logger.Error("error with service get all users: ", logger.Error(err))
		return users, err
	}
	return users, nil
}