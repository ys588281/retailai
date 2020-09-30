package response

type CreateRecipeResponse struct {
	Message string `json:"message"`
	RecipeList []*RecipeWithTime `json:"recipe"`
}

type GetRecipeListResponse struct {
	RecipeList []*Recipe `json:"recipes"`
}

type GetRecipeResponse struct {
	Message string `json:"message"`
	RecipeList []*Recipe `json:"recipe"`
}

type GeneralRecipeResponse struct {
	Message string `json:"message"`
	Required string `json:"required,omitempty"`
}

type Recipe struct {
	ID   int    `json:"id"`
	Title string `json:"title"`
	PreparationTime  string    `json:"preparation_time"`
	Serves  string    `json:"Serves"`
	Ingredients string `json:"ingrediencts"`
	Cost string `json:"costs"`
}

type RecipeWithTime struct {
	ID   int    `json:"id"`
	Title string `json:"title"`
	PreparationTime  string    `json:"preparation_time"`
	Serves  string    `json:"Serves"`
	Ingredients string `json:"ingrediencts"`
	Cost string `json:"costs"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}