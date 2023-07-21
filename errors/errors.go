package errors

import "fmt"

var (
	GeneratePwdErr         = fmt.Errorf("ошибка генерации пароля")
	IncorrectPwdErr        = fmt.Errorf("неверный пароль")
	InvalidTokenErr        = fmt.Errorf("вы не авторизованы: неверный формат токена")
	RecipeNotFoundErr      = fmt.Errorf("рецепт не найден")
	CannotUpdateRecipeErr  = fmt.Errorf("вам не принадлежит данный рецепт, вы не можете его изменить")
	UserAlreadyExistsErr   = fmt.Errorf("пользователь с такими данными уже существует")
	RecipeDuplicateErr     = fmt.Errorf("рецепт с таким названием уже существует уже существует")
	NotARecipeOwnerErr     = fmt.Errorf("вам не принадлежит данный рецепт, вы не можете его удалить")
	RequestValidationError = fmt.Errorf("невалидный запрос")

	ErrInvalidRequest = fmt.Errorf("не удалось выполнить запрос")

	CustomErrors = map[error]struct{}{
		GeneratePwdErr:         {},
		IncorrectPwdErr:        {},
		InvalidTokenErr:        {},
		RecipeNotFoundErr:      {},
		CannotUpdateRecipeErr:  {},
		NotARecipeOwnerErr:     {},
		RequestValidationError: {},
		UserAlreadyExistsErr:   {},
		RecipeDuplicateErr:     {},
	}
)

// проверка на кастомные ошибки введена для того,
// чтобы пользователю возвращался человекоподобный текст, вместо условный "err no rows in result set"
func IsCustom(e error) bool {
	_, ok := CustomErrors[e]
	return ok
}
