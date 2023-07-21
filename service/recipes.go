package service

import (
	"context"

	errors2 "recitas/errors"
	recipes2 "recitas/modules/recipes"
	"recitas/modules/users"
	"recitas/proto/rec"

	"github.com/google/uuid"
)

func (a Api) CreateRecipe(ctx context.Context, req *rec.CreateRecipeRequest) (res *rec.CreateRecipeResponse, err error) {
	res = new(rec.CreateRecipeResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	recipe := recipes2.Recipe{
		Steps: func() []recipes2.Step {
			steps := make([]recipes2.Step, 0, len(req.Steps))

			for _, s := range req.Steps {
				steps = append(steps, recipes2.Step{
					Duration:    s.Duration,
					Description: s.Description,
				})
			}

			return steps
		}(),
		Name:        req.Name,
		Description: req.Description,
		Ingredients: req.Ingredients,
	}

	token := ctx.Value(users.UserTokenKey)
	if tokenString, ok := token.(users.Token); !ok {
		return res, errors2.InvalidTokenErr
	} else {
		uid, err := a.users.CheckUserToken(ctx, tokenString)
		if err != nil {
			return res, err
		}

		recipe.UserUid, err = uuid.Parse(uid)
		if err != nil {
			return res, err
		}
	}

	err = a.recipes.CreateRecipe(ctx, &recipe)
	res.Id = recipe.Id
	return res, err
}

func (a Api) UpdateRecipe(ctx context.Context, req *rec.UpdateRecipeRequest) (res *rec.UpdateRecipeResponse, err error) {
	res = new(rec.UpdateRecipeResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	recipe := recipes2.Recipe{
		Id: req.Id,
		Steps: func() []recipes2.Step {
			steps := make([]recipes2.Step, 0, len(req.Steps))

			for _, s := range req.Steps {
				steps = append(steps, recipes2.Step{
					Duration:    s.Duration,
					Description: s.Description,
				})
			}

			return steps
		}(),
		Name:        req.Name,
		Description: req.Description,
		Ingredients: req.Ingredients,
	}

	token := ctx.Value(users.UserTokenKey)
	if tokenString, ok := token.(users.Token); !ok {
		return res, errors2.InvalidTokenErr
	} else {
		uid, err := a.users.CheckUserToken(ctx, tokenString)
		if err != nil {
			return res, err
		}

		recipe.UserUid, err = uuid.Parse(uid)
		if err != nil {
			return res, err
		}
	}

	err = a.recipes.UpdateRecipe(ctx, &recipe)
	res.Id = recipe.Id
	return res, err
}

func (a Api) DeleteRecipe(ctx context.Context, req *rec.DeleteRecipeRequest) (res *rec.DeleteRecipeResponse, err error) {
	res = new(rec.DeleteRecipeResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	recipe := recipes2.Recipe{
		Id: req.Id,
	}

	token := ctx.Value(users.UserTokenKey)
	if tokenString, ok := token.(users.Token); !ok {
		return res, errors2.InvalidTokenErr
	} else {
		uid, err := a.users.CheckUserToken(ctx, tokenString)
		if err != nil {
			return res, err
		}

		recipe.UserUid, err = uuid.Parse(uid)
		if err != nil {
			return res, err
		}
	}

	err = a.recipes.DeleteRecipe(ctx, &recipe)
	return res, err
}

func (a Api) GetRecipe(ctx context.Context, req *rec.GetRecipeRequest) (res *rec.GetRecipeResponse, err error) {
	res = new(rec.GetRecipeResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	recipe := recipes2.Recipe{
		Id: req.Id,
	}
	err = a.recipes.GetRecipe(ctx, &recipe)
	if err != nil {
		return res, err
	}

	res.Id = recipe.Id
	res.Name = recipe.Name
	res.Description = recipe.Description
	res.Ingredients = recipe.Ingredients
	res.Steps = func() []*rec.Step {
		steps := make([]*rec.Step, 0, len(recipe.Steps))

		for _, s := range recipe.Steps {
			steps = append(steps, &rec.Step{
				Description: s.Description,
				Duration:    s.Duration,
			})
		}

		return steps
	}()
	res.TotalTime = recipe.TotalTime

	return res, err
}

func (a Api) GetAllRecipes(ctx context.Context, req *rec.GetAllRecipesRequest) (res *rec.GetAllRecipesResponse, err error) {
	res = new(rec.GetAllRecipesResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	recipes := recipes2.Recipes{
		Ingredients: req.Ingredients,
		MinTime:     req.MinTime,
		MaxTime:     req.MaxTime,
		TimeOrder:   req.TimeOrder,
	}
	err = a.recipes.GetRecipes(ctx, &recipes)
	if err != nil {
		return res, err
	}

	for _, r := range recipes.Recipes {
		res.Recipes = append(res.Recipes, &rec.GetRecipeResponse{
			Id:          r.Id,
			Name:        r.Name,
			Description: r.Description,
			Ingredients: r.Ingredients,
			Steps: func() []*rec.Step {
				steps := make([]*rec.Step, 0, len(r.Steps))

				for _, s := range r.Steps {
					steps = append(steps, &rec.Step{
						Description: s.Description,
						Duration:    s.Duration,
					})
				}

				return steps
			}(),
			TotalTime: r.TotalTime,
		})
	}

	return res, err
}

func (a Api) GetAllRecipesByUser(ctx context.Context, req *rec.GetAllRecipesByUserRequest) (res *rec.GetAllRecipesByUserResponse, err error) {
	res = new(rec.GetAllRecipesByUserResponse)

	if err = req.Validate(); err != nil {
		return res, errors2.RequestValidationError
	}

	defer func() {
		// если встретилась ошибка, которую мы не переописали, то вернём стандартное "не удалось выполнить запрос"
		if err != nil && !errors2.IsCustom(err) {
			err = errors2.ErrInvalidRequest
		}
	}()

	recipes := recipes2.Recipes{
		Ingredients: req.Ingredients,
		MinTime:     req.MinTime,
		MaxTime:     req.MaxTime,
		TimeOrder:   req.TimeOrder,
	}
	token := ctx.Value(users.UserTokenKey)
	if tokenString, ok := token.(users.Token); !ok {
		return res, errors2.InvalidTokenErr
	} else {
		uid, err := a.users.CheckUserToken(ctx, tokenString)
		if err != nil {
			return res, err
		}

		recipes.UserUid, err = uuid.Parse(uid)
		if err != nil {
			return res, err
		}
	}

	err = a.recipes.GetRecipes(ctx, &recipes)
	if err != nil {
		return res, err
	}

	for _, r := range recipes.Recipes {
		res.Recipes = append(res.Recipes, &rec.GetRecipeResponse{
			Id:          r.Id,
			Name:        r.Name,
			Description: r.Description,
			Ingredients: r.Ingredients,
			Steps: func() []*rec.Step {
				steps := make([]*rec.Step, 0, len(r.Steps))

				for _, s := range r.Steps {
					steps = append(steps, &rec.Step{
						Description: s.Description,
						Duration:    s.Duration,
					})
				}

				return steps
			}(),
			TotalTime: uint32(r.TotalTime),
		})
	}
	return res, err
}
