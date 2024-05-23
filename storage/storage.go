package storage

import (
	"context"
	"durger/models"
)

type IStorage interface {
	CloseDB()
	UserStorage() UserStorage
}

type UserStorage interface {
	Create(ctx context.Context, user models.AddUser) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (models.User, error)
	GetAll(ctx context.Context) (models.AllUsers, error)
}
