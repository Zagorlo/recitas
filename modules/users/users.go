package users

import (
	"context"
	"strings"

	errors2 "recitas/errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Uid            uuid.UUID `db:"uid"`
	Login          string    `db:"login"`
	Token          string    `db:"token"`
	Password       string    `db:"password"`
	HashedPassword string    `db:"hashed_password"`
}

func (u *User) create(ctx context.Context, conn *sqlx.DB) error {
	const createUser = `
		insert into rec.users
		(uid, login, password)
		values($1, $2, $3)
		returning uid;
	`

	err := conn.Get(&u.Uid, createUser, u.Uid, u.Login, u.Password)
	if err != nil {
		if strings.Contains(err.Error(), "violate") {
			return errors2.UserAlreadyExistsErr
		} else {
			return err
		}
	}

	return nil
}

func (u *User) login(ctx context.Context, conn *sqlx.DB) error {
	const loginUser = `
		select u.uid, u.password as hashed_password
		from rec.users u
		where u.login = $1
	`

	err := conn.Get(u, loginUser, u.Login)
	if err != nil {
		return err
	}

	return nil
}
