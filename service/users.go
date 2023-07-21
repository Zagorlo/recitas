package service

import (
	"context"

	errors2 "recitas/errors"
	"recitas/modules/users"
	"recitas/proto/rec"
)

func (a Api) Register(ctx context.Context, req *rec.RegisterRequest) (res *rec.RegisterResponse, err error) {
	res = new(rec.RegisterResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	user := users.User{
		Login:    req.Login,
		Password: req.Password,
	}

	err = a.users.RegisterUser(ctx, &user)
	res.Token = user.Token

	return res, err
}

func (a Api) Login(ctx context.Context, req *rec.LoginRequest) (res *rec.LoginResponse, err error) {
	res = new(rec.LoginResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	user := users.User{
		Login:    req.Login,
		Password: req.Password,
	}

	err = a.users.LoginUser(ctx, &user)
	if err != nil {
		return res, err
	}

	res.Token = user.Token
	return res, err
}
