package Router

import (
	CategoryRepository "go-postgres/repository"
	CommentRepository "go-postgres/repository"
	MainRepository "go-postgres/repository"
	ModuleRepository "go-postgres/repository"
	OrganisationRepository "go-postgres/repository"
	ProductRepository "go-postgres/repository"
	UserRepository "go-postgres/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//modules
	router.HandleFunc("/module/{module_id}", ModuleRepository.GetModule).Methods("GET", "OPTIONS")
	router.HandleFunc("/module", ModuleRepository.GetAllModule).Methods("GET", "OPTIONS")
	router.HandleFunc("/module/init/{module_id}", ModuleRepository.InitModule).Methods("GET", "OPTIONS")

	//product
	router.HandleFunc("/product/feed/{module_id}", ProductRepository.GetHotProducts).Queries("sort_type", "{sort_type}").Methods("GET", "OPTIONS")
	router.HandleFunc("/product/feed/{module_id}", ProductRepository.GetHotProductsFeed).Methods("GET", "OPTIONS")
	router.HandleFunc("/product/{product_id}", ProductRepository.GetProduct).Methods("GET", "OPTIONS")

	//organisation
	router.HandleFunc("/organisation/{organisation_id}", OrganisationRepository.GetOrganisation).Methods("GET", "OPTIONS")

	//category_id
	router.HandleFunc("/category/{category_id}", CategoryRepository.GetCategory).Methods("GET", "OPTIONS")
	//comments
	router.HandleFunc("/comment/{product_id}", CommentRepository.GetAllCommentsRep).Methods("GET", "OPTIONS")

	//user
	router.HandleFunc("/user/{id}", UserRepository.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/user", UserRepository.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/newuser", UserRepository.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/{id}", UserRepository.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deleteuser/{id}", UserRepository.DeleteUser).Methods("DELETE", "OPTIONS")

	//main endpoint
	router.HandleFunc("/main", MainRepository.Main).Methods("GET", "OPTIONS")
	router.HandleFunc("/orgpage", MainRepository.OrgPage).Methods("GET", "OPTIONS")
	router.PathPrefix("/template/").Handler(http.StripPrefix("/template/", http.FileServer(http.Dir("./public/template/"))))

	//icons
	router.HandleFunc("/static/{icon}", MainRepository.GetStatic).Methods("GET", "OPTIONS")
	return router
}
