package users

import (
	"context"
	"time"

	"recitas/config"
	"recitas/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const pwdHashCost = 13

const (
	UserTokenKey             = "x-user-token"
	tokenClaimUsername       = "_user_name"
	tokenClaimExpirationTime = "_exp"
	tokenClaimGroupDn        = "_gdn"
	tokenExpirationTimeHour  = 720

	tokenGroupName = "_casting_recipes"
)

func SetSignKey(key string) {
	signKey = key
}

type Token = string

var signKey = ""

func NewUserModule(cfg config.Config) UserModule {
	return UserModule{conn: cfg.Postgres.Conn}
}

type UserModule struct {
	conn *sqlx.DB
}

func (rm UserModule) CheckUserToken(ctx context.Context, token Token) (uid string, err error) {
	uid, err = checkUserToken(token)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (rm UserModule) RegisterUser(ctx context.Context, user *User) (err error) {
	user.Uid = uuid.New()

	user.Password, err = generatePwdHash(user.Password)
	if err != nil {
		return err
	}

	err = user.create(ctx, rm.conn)
	if err != nil {
		return err
	}

	user.Token, err = createToken(user.Uid.String())
	if err != nil {
		return err
	}

	return nil
}

func (rm UserModule) LoginUser(ctx context.Context, user *User) (err error) {
	err = user.login(ctx, rm.conn)
	if err != nil {
		return err
	}

	if !checkPwsHash(user.Password, user.HashedPassword) {
		return errors.IncorrectPwdErr
	}

	user.Token, err = createToken(user.Uid.String())
	if err != nil {
		return err
	}

	return nil
}

func generatePwdHash(pass string) (string, error) {
	hash, err := getPwdHash(pass, pwdHashCost)
	if err != nil {
		return "", err
	}

	if checkPwsHash(pass, hash) {
		return hash, nil
	}

	return "", errors.GeneratePwdErr
}

func getPwdHash(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func checkPwsHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(creds string) (string, error) {
	timeNow := time.Now()

	tokenClaims := make(jwt.MapClaims)
	tokenClaims[tokenClaimUsername] = creds
	tokenClaims[tokenClaimExpirationTime] = timeNow.Add(time.Hour * tokenExpirationTimeHour).Unix()
	tokenClaims[tokenClaimGroupDn] = tokenGroupName

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaims)
	tokenString, err := token.SignedString([]byte(signKey))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func checkUserToken(reqToken string) (string, error) {
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err == nil && token.Valid {
		claims, err := extractClaims(reqToken)
		if err == nil && claims[tokenClaimGroupDn] == tokenGroupName {
			user := claims[tokenClaimUsername].(string)
			if len(user) > 0 {
				return user, nil
			}
		}
		return "", err
	} else {
		return "", errors.InvalidTokenErr
	}
}

func extractClaims(tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(signKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, errors.GeneratePwdErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.GeneratePwdErr
	}
}
