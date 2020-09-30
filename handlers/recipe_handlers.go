package RecipeHandlers
import (
	"net/http"
	"database/sql"
	"encoding/json"
	"strconv"
	
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	
	requests "./requests"
	responses "./responses"
	model "../model"
	common "../common"
)
type recipeHandlers struct {
	router *mux.Router
	db     *sql.DB
}

func InitializeRecipeHandlers(router *mux.Router, db *sql.DB) {
	r := &recipeHandlers{
		router: router,
		db: db,
	}
	r.router.HandleFunc("/recipe", r.getRecipeList).Methods("GET")
	r.router.HandleFunc("/recipe", r.createRecipe).Methods("POST")
	r.router.HandleFunc("/recipe/{id:[0-9]+}", r.updateRecipe).Methods("PATCH")
	r.router.HandleFunc("/recipe/{id:[0-9]+}", r.getRecipe).Methods("GET")
	r.router.HandleFunc("/recipe/{id:[0-9]+}", r.deleteRecipe).Methods("DELETE")
}

func (rh *recipeHandlers) createRecipe(w http.ResponseWriter, r *http.Request) {
	var req requests.CreateRecipeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		common.RespondBadRequest(w, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()

	recipe := req.ToRecipeModel()
	if err := recipe.CreateRecipe(rh.db); err != nil {
		common.RespondWithError(w, err.Error())
		return
	}

	res := recipe.ToGetRecipeResponse()
	common.RespondWithJSON(w, http.StatusCreated, res)
}

func (rh *recipeHandlers) updateRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		common.RespondBadRequest(w, "Invalid recipe ID", "")
		return
	}
	var req requests.CreateRecipeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		common.RespondBadRequest(w, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()

	recipe := req.ToRecipeModel()
	recipe.ID = id
	if err := recipe.UpdateRecipe(rh.db); err != nil {
		common.RespondWithError(w, err.Error())
		return
	}

	res := &responses.GetRecipeResponse{
		Message: "Recipe details by id",
		RecipeList: toRecipe(recipe),
	}
	common.RespondWithJSON(w, http.StatusOK, res)
}

func (rh *recipeHandlers) getRecipeList(w http.ResponseWriter, r *http.Request) {
	recipeList, err := model.GetRecipeList(rh.db)
	if err != nil {
		common.RespondWithError(w, err.Error())
		return
	}

	resList := make([]*responses.Recipe, 0)
	for _, recipe :=  range recipeList {
		resList = append(resList, recipe.ToRecipe())
	}

	common.RespondWithJSON(w, http.StatusOK, &responses.GetRecipeListResponse{
		RecipeList: resList,
	})
}

func (rh *recipeHandlers) getRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		common.RespondBadRequest(w, "Invalid recipe ID", "")
		return
	}

	recipe := &model.Recipe{ID: id}

	if err := recipe.GetRecipe(rh.db); err != nil {
		switch err {
		case sql.ErrNoRows:
			common.RespondNotFound(w)
		default:
			common.RespondWithError(w, err.Error())
		}
		return
	}
	
	res := &responses.GetRecipeResponse{
		Message: "Recipe details by id",
		RecipeList: toRecipe(recipe),
	}
	
	common.RespondWithJSON(w, http.StatusOK, res)
}

func toRecipe(recipe *model.Recipe) []*responses.Recipe {
	rl := make([]*responses.Recipe, 0)
	rl = append(rl, recipe.ToRecipe())
	return rl
}

func (rh *recipeHandlers) deleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		common.RespondBadRequest(w, "Invalid user ID", "")
		return
	}

	recipe := model.Recipe{ID: id}

	if err := recipe.GetRecipe(rh.db); err != nil {
		switch err {
		case sql.ErrNoRows:
			common.RespondNotFound(w)
		default:
			common.RespondWithError(w, err.Error())
		}
		return
	}


	if err := recipe.DeleteRecipe(rh.db); err != nil {
		common.RespondWithError(w, err.Error())
		return
	}
	
	common.RespondWithJSON(w, http.StatusOK, &responses.GeneralRecipeResponse{
		Message: "Recipe successfully removed!",
	})
}