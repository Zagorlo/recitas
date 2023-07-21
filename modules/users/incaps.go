package users

import "context"

type IUsers interface {
	CheckUserToken(ctx context.Context, token Token) (uid string, err error)
	RegisterUser(ctx context.Context, user *User) error
	LoginUser(ctx context.Context, user *User) error
}
