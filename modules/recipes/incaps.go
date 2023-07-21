package recipes

import "context"

type IRecipes interface {
	UpdateRecipe(ctx context.Context, recipe *Recipe) (err error)
	DeleteRecipe(ctx context.Context, recipe *Recipe) (err error)
	CreateRecipe(ctx context.Context, recipe *Recipe) error
	GetRecipes(ctx context.Context, recipes *Recipes) error
	GetRecipe(ctx context.Context, recipe *Recipe) (err error)
}
