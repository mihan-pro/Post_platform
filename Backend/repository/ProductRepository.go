package Repository

import (

	// JSon Parser
	"encoding/json"

	// import ModuleController

	CategoryController "go-postgres/controller"
	ModuleController "go-postgres/controller"
	OrganisationController "go-postgres/controller"
	ProductController "go-postgres/controller"

	Helpers "go-postgres/helpers"
	CategoryModel "go-postgres/models"
	ModuleModel "go-postgres/models"
	OrganisationModel "go-postgres/models"
	ProductModel "go-postgres/models"

	//  access the request and response object of the api
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

type feed struct {
	Products      []ProductModel.Product
	Organisations []OrganisationModel.Organisation
	Categorys     []CategoryModel.Category
	Module        ModuleModel.Module
}

// ================================== Repositories to handle controllers =========================================

// GetProduct - sindle module
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if Helpers.IsValidUUID(params["product_id"]) {

		product, err := ProductController.GetSingleProduct(params["product_id"])

		if err != nil {
			log.Errorln("Unable to get product.", err)
			res := error{
				Message: "Продукта не существует",
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

// GetHotProducts - all modules
func GetHotProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if Helpers.IsValidUUID(params["module_id"]) && len(params["sort_type"]) > 0 {
		products, err := ProductController.GetHot(params["module_id"], params["sort_type"])

		module, err := ModuleController.GetModule(params["module_id"])
		organisations, err := OrganisationController.GetOrganisations(params["module_id"])
		categorys, err := CategoryController.GetCategorys(params["module_id"])

		if err != nil {
			log.Errorln("Unable to get products.", err)
			res := error{
				Message: "Нет данных для отображения",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res)
		} else {
			var response = new(feed)
			response.Products = products
			response.Organisations = organisations
			response.Categorys = categorys
			response.Module = module

			json.NewEncoder(w).Encode(response)
		}
	} else {
		res := error{
			Message: "Нет данных для отображения",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
	}

}

// GetHotProducts - all modules
func GetHotProductsFeed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if Helpers.IsValidUUID(params["module_id"]) {
		products, err := ProductController.GetHotProductsCollection(params["module_id"])

		module, err := ModuleController.GetModule(params["module_id"])
		organisations, err := OrganisationController.GetOrganisations(params["module_id"])
		categorys, err := CategoryController.GetCategorys(params["module_id"])

		if err != nil {
			log.Errorln("Unable to get products.", err)
			res := error{
				Message: "Нет данных для отображения",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res)
		} else {
			var response = new(feed)
			response.Products = products
			response.Organisations = organisations
			response.Categorys = categorys
			response.Module = module
			json.NewEncoder(w).Encode(response)
		}
	} else {
		res := error{
			Message: "Нет данных для отображения",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
	}

}
