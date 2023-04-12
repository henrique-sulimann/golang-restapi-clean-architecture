package repository

import "golang.org/x/net/context"

type Connection interface {
	Insert(ctx context.Context, statement ...any) error
	Delete(ctx context.Context, statement ...any) error
	Update(ctx context.Context, statement ...any) error
	FindAll(ctx context.Context, statement ...any) error
	FindOne(ctx context.Context, statement ...any) error
	Close(ctx context.Context) error
}
