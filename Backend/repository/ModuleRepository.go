package Repository

import (

	// JSon Parser
	"encoding/json"

	// import ModuleController
	Template "go-postgres/config"
	ModuleController "go-postgres/controller"

	//  access the request and response object of the api
	"net/http"

	"github.com/gorilla/mux"

	"github.com/prometheus/common/log"
)

// ================================== Repositories to handle controllers =========================================

// GetModule - sindle module
func GetModule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// call the get User function
	module, err := ModuleController.GetModule(params["module_id"])

	if err != nil {
		log.Errorln("Unable to get module.", err)
		res := error{
			Message: "Модуля не существует",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(module)
	}
}

//InitModule - asdasd
func InitModule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	apiKey, err := ModuleController.RenderModule(params["module_id"])
	log.Infof(apiKey)
	if err != nil {
		log.Errorln("Unable to get module.", err)
		res := error{
			Message: "Модуля нет",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
	} else {
		log.Info("render module: " + apiKey + ".html")
		Template.Module.ExecuteTemplate(w, apiKey+".html", nil)
	}

}

// GetAllModule - all modules
func GetAllModule(w http.ResponseWriter, r *http.Request) {

	// retreive all the users in db
	modules, err := ModuleController.GetAllModules()

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

}
