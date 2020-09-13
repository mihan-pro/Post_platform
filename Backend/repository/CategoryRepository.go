package Repository

import (

	// JSon Parser
	"encoding/json"

	// import ModuleController

	CategoryController "go-postgres/controller"

	Helpers "go-postgres/helpers"

	//  access the request and response object of the api
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

// ================================== Repositories to handle controllers =========================================

// GetCategory - sindle module
func GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if Helpers.IsValidUUID(params["category_id"]) {

		product, err := CategoryController.GetSingleCategory(params["category_id"])

		if err != nil {
			log.Errorln("Unable to get product.", err)
			res := error{
				Message: "Организации не существует",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res)
		} else {
			json.NewEncoder(w).Encode(product)
		}
	} else {
		res := error{
			Message: "Нет данных для отображения",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
	}
}
