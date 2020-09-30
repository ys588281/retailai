package common

import(
	"net/http"
	"encoding/json"
	
	responses "../handlers/responses"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondNotFound(w http.ResponseWriter) {
	RespondWithJSON(w, http.StatusNotFound, &responses.GeneralRecipeResponse{
		Message: "No recipe found!",
	})
}

func RespondBadRequest(w http.ResponseWriter, messgae, required string) {
	res := &responses.GeneralRecipeResponse{
		Message: "No recipe found!",
	}
	if required != "" {
		res.Required = required
	}
	RespondWithJSON(w, http.StatusBadRequest, res)
}

func RespondWithError(w http.ResponseWriter, message string) {
	RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": message})
}