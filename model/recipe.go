package model

import (
	"fmt"
	"time"
	"database/sql"

	response "../handlers/responses"
)

type Recipe struct {
	ID   int
	Title string
	PreparationTime  string
	Serves  string
	Ingredients string
	Cost string
	CreatedAt string
	UpdatedAt string
}


func (r *Recipe) ToGetRecipeResponse() *response.CreateRecipeResponse {
	rl := make([]*response.RecipeWithTime, 0)
	rl = append(rl, r.ToRecipeWithTime())
	return &response.CreateRecipeResponse{
		Message: "Recipe successfully created!",
		RecipeList: rl,
	}
}

func (r *Recipe) ToRecipeWithTime() *response.RecipeWithTime {
	return &response.RecipeWithTime{
		ID: r.ID,
		Title: r.Title,
		PreparationTime: r.PreparationTime,
		Serves: r.Serves,
		Ingredients: r.Ingredients,
		Cost: r.Cost,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func (r *Recipe) ToRecipe() *response.Recipe {
	return &response.Recipe{
		ID: r.ID,
		Title: r.Title,
		PreparationTime: r.PreparationTime,
		Serves: r.Serves,
		Ingredients: r.Ingredients,
		Cost: r.Cost,
	}
}

func (r *Recipe) CreateRecipe(db *sql.DB) error {
	now := time.Now()
	r.CreatedAt = now.Format("2006-01-02 15:04:05")
	r.UpdatedAt = now.Format("2006-01-02 15:04:05")

	_, err := db.Exec("INSERT INTO recipe(title, preparation_time, serves, ingredients, cost, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)", r.Title, r.PreparationTime, r.Serves, r.Ingredients, r.Cost, r.CreatedAt, r.UpdatedAt)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT max(id) FROM recipe").Scan(&r.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *Recipe) UpdateRecipe(db *sql.DB) error {
	now := time.Now()
	old := Recipe{ID: r.ID}
	old.GetRecipe(db)
	old.UpdatedAt = now.Format("2006-01-02 15:04:05")
	if r.Title == "" {
		r.Title = old.Title
	}
	if r.PreparationTime == "" {
		r.PreparationTime = old.PreparationTime
	}
	if r.Serves == "" {
		r.Serves = old.Serves
	}
	if r.Ingredients == "" {
		r.Ingredients = old.Ingredients
	}
	if r.Cost == "" {
		r.Cost = old.Cost 
	}
	_, err := db.Exec("UPDATE recipe SET title = ?, preparation_time = ?, serves = ?, ingredients = ?, cost = ?, updated_at = ? WHERE id = ?", r.Title, r.PreparationTime, r.Serves, r.Ingredients, r.Cost, r.UpdatedAt, r.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetRecipeList(db *sql.DB) ([]Recipe, error) {
	statement := "SELECT * FROM recipe"

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipeList := []Recipe{}

	for rows.Next() {
		var u Recipe
		if err := rows.Scan(&u.ID, &u.Title, &u.PreparationTime, &u.Serves, &u.Ingredients, &u.Cost, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		recipeList = append(recipeList, u)
	}

	return recipeList, nil

}

func (r *Recipe) GetRecipe(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM recipe WHERE id=%d", r.ID)
	return db.QueryRow(statement).Scan(&r.ID, &r.Title, &r.PreparationTime, &r.Serves, &r.Ingredients, &r.Cost, &r.CreatedAt, &r.UpdatedAt)
}

func (r *Recipe) DeleteRecipe(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM recipe WHERE id=%d", r.ID)
	_, err := db.Exec(statement)
	return err
}