package recipes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"recitas/errors"
)

type Recipe struct {
	UserUid     uuid.UUID       `db:"-"`
	Steps       []Step          `db:"-"`
	StepsJson   json.RawMessage `db:"steps"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Id          uint32          `db:"id"`
	Ingredients pq.StringArray  `db:"ingredients"`
	TotalTime   uint32          `db:"total_time"`
}

type Step struct {
	Duration    uint32 `json:"duration"`
	Description string `json:"description"`
}

type Recipes struct {
	Ingredients pq.StringArray
	MinTime     uint32
	MaxTime     uint32
	TimeOrder   string
	UserUid     uuid.UUID `db:"-"`
	Recipes     []Recipe
}

func (r *Recipe) create(ctx context.Context, conn *sqlx.DB) (err error) {
	const createRecipe = `
		insert into rec.recipes(user_uid, name, description, steps, total_time)
		values($1, $2, $3, $4, $5)
		returning id
	`

	const putRecipeIngredients = `
		insert into rec.recipe_ingredients(recipe_id, name)
		select $1, unnest($2::_text)
	`

	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	err = tx.Get(&r.Id, createRecipe, r.UserUid, r.Name, r.Description, r.StepsJson, r.TotalTime)
	if err != nil {
		_ = tx.Rollback()
		if strings.Contains(err.Error(), "violate") {
			return errors.RecipeDuplicateErr
		} else {
			return err
		}
	}

	_, err = tx.Exec(putRecipeIngredients, r.Id, r.Ingredients)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Recipe) update(ctx context.Context, conn *sqlx.DB) (err error) {
	const clearIngredients = `
		delete from rec.recipe_ingredients
		where recipe_id = $1
	`

	const updateRecipe = `
		update rec.recipes
		set 
			name = $1,
			description = $2,
			steps = $3,
			total_time = $4
		where id = $5
		returning user_uid
	`

	const putRecipeIngredients = `
		insert into rec.recipe_ingredients(recipe_id, name)
		select $1, unnest($2::_text)
	`

	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(clearIngredients, r.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	var uid uuid.UUID
	err = tx.Get(&uid, updateRecipe, r.Name, r.Description, r.StepsJson, r.TotalTime, r.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.RecipeNotFoundErr
		}
		_ = tx.Rollback()
		return err
	}

	if uid != r.UserUid {
		_ = tx.Rollback()
		return errors.CannotUpdateRecipeErr
	}

	_, err = tx.Exec(putRecipeIngredients, r.Id, r.Ingredients)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Recipe) delete(ctx context.Context, conn *sqlx.DB) (err error) {
	const clearIngredients = `
		delete from rec.recipe_ingredients
		where recipe_id = $1
	`

	const deleteRecipe = `
		delete from rec.recipes
		where id = $1
		returning user_uid
	`

	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(clearIngredients, r.Id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	var uid uuid.UUID
	err = tx.Get(&uid, deleteRecipe, r.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.RecipeNotFoundErr
		}
		_ = tx.Rollback()
		return err
	}

	if uid != r.UserUid {
		_ = tx.Rollback()
		return errors.NotARecipeOwnerErr
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Recipes) get(ctx context.Context, conn *sqlx.DB) (err error) {
	getRecipes := `
		select id, steps, r.name, description, total_time, array_agg(ri.name) as ingredients
		from rec.recipes r
			left join rec.recipe_ingredients ri 
				on ri.recipe_id = r.id
		%s
		group by id, steps, r.name,description, total_time
		%s
		%s
	`

	whereClause := ""
	if r.MinTime > 0 {
		whereClause += fmt.Sprintf("total_time >= %d", r.MinTime)
	}
	if r.MaxTime > 0 {
		if len(whereClause) > 0 {
			whereClause += " and "
		}
		whereClause += fmt.Sprintf("total_time <= %d", r.MaxTime)
	}
	if r.UserUid != uuid.Nil {
		if len(whereClause) > 0 {
			whereClause += " and "
		}
		// в тип uuid.UUID инъекцию не засунешь
		whereClause += fmt.Sprintf("r.user_uid = '%s'", r.UserUid.String())
	}
	if len(whereClause) > 0 {
		whereClause = fmt.Sprintf("where %s", whereClause)
	}

	orderByClause := ""
	if r.TimeOrder == "desc" {
		orderByClause = "order by total_time desc"
	} else {
		orderByClause = "order by total_time asc"
	}

	if len(r.Ingredients) > 0 {
		// slqx через вставку параметров ($) защищает от инъекций
		getRecipes = fmt.Sprintf(getRecipes, whereClause, fmt.Sprintf("having array_agg(ri.name) @> $1"), orderByClause)
		err = conn.Select(&r.Recipes, getRecipes, r.Ingredients)
	} else {
		getRecipes = fmt.Sprintf(getRecipes, whereClause, "", orderByClause)
		err = conn.Select(&r.Recipes, getRecipes)
	}

	if err != nil {
		return
	}

	return nil
}

func (r *Recipe) get(ctx context.Context, conn *sqlx.DB) (err error) {
	getRecipe := `
		select id, steps, r.name, description, total_time, array_agg(ri.name) as ingredients
		from rec.recipes r
			left join rec.recipe_ingredients ri 
				on ri.recipe_id = r.id
		where id = $1
		group by id, steps, r.name,description, total_time		
	`

	err = conn.Get(r, getRecipe, r.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.RecipeNotFoundErr
		}
		return
	}

	return nil
}
