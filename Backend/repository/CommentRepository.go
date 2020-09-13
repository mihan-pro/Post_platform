package Repository

import (

	// JSon Parser
	"encoding/json"

	// import ModuleController

	CommentController "go-postgres/controller"
	Helpers "go-postgres/helpers"

	//  access the request and response object of the api
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

// ================================== Repositories to handle controllers =========================================

// GetAllModule - all modules
func GetAllCommentsRep(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if Helpers.IsValidUUID(params["product_id"]) {
		// retreive all the users in db
		modules, err := CommentController.GetAllComments(params["product_id"])

		if err != nil {
			log.Errorln("Unable to get modules.", err)
			res := error{
				Message: "Модулей нет",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res)
		} else {
			json.NewEncoder(w).Encode(modules)
		}
	} else {
		res := error{
			Message: "Нет данных для отображения",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
	}
}
