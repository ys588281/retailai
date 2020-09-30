package requests

import (
	model "../../model"
)

type CreateRecipeRequest struct {
	Title string `json:"title"`
	PreparationTime  string    `json:"preparation_time"`
	Serves  string    `json:"Serves"`
	Ingredients string `json:"ingredients"`
	Cost string `json:"cost"`
}

func (c *CreateRecipeRequest) ToRecipeModel() *model.Recipe {
	return &model.Recipe{
		Title: c.Title,
		PreparationTime: c.PreparationTime,
		Serves: c.Serves,
		Ingredients: c.Ingredients,
		Cost: c.Cost,
	}
}