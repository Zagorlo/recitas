package recipes

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"

	"recitas/config"
)

func NewRecipeModule(cfg config.Config) RecipeModule {
	return RecipeModule{conn: cfg.Postgres.Conn}
}

type RecipeModule struct {
	conn *sqlx.DB
}

func (rm RecipeModule) CreateRecipe(ctx context.Context, recipe *Recipe) (err error) {
	recipe.StepsJson, err = json.Marshal(recipe.Steps)
	if err != nil {
		return err
	}

	for _, s := range recipe.Steps {
		recipe.TotalTime += s.Duration
	}

	err = recipe.create(ctx, rm.conn)
	if err != nil {
		return err
	}

	return nil
}

func (rm RecipeModule) UpdateRecipe(ctx context.Context, recipe *Recipe) (err error) {
	recipe.StepsJson, err = json.Marshal(recipe.Steps)
	if err != nil {
		return err
	}

	for _, s := range recipe.Steps {
		recipe.TotalTime += s.Duration
	}

	err = recipe.update(ctx, rm.conn)
	if err != nil {
		return err
	}

	return nil
}

func (rm RecipeModule) DeleteRecipe(ctx context.Context, recipe *Recipe) (err error) {
	err = recipe.delete(ctx, rm.conn)
	if err != nil {
		return err
	}

	return nil
}

func (rm RecipeModule) GetRecipes(ctx context.Context, recipes *Recipes) (err error) {
	err = recipes.get(ctx, rm.conn)
	if err != nil {
		return err
	}

	for i, r := range recipes.Recipes {
		err = json.Unmarshal(r.StepsJson, &recipes.Recipes[i].Steps)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rm RecipeModule) GetRecipe(ctx context.Context, recipe *Recipe) (err error) {
	err = recipe.get(ctx, rm.conn)
	if err != nil {
		return err
	}

	err = json.Unmarshal(recipe.StepsJson, &recipe.Steps)
	if err != nil {
		return err
	}

	return nil
}
