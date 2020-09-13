package Repository

import (
	Template "go-postgres/config"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

type error struct {
	Message string `json:"message,omitempty"`
}

type response struct {
	ID      int64     `json:"id,omitempty"`
	Message string    `json:"message,omitempty"`
	UserID  uuid.UUID `json:"user_id,omitempty"`
}

// Main test
func Main(w http.ResponseWriter, r *http.Request) {

	Template.Main.ExecuteTemplate(w, "mainpage.html", nil)
}

// OrgPage test
func OrgPage(w http.ResponseWriter, r *http.Request) {

	Template.Main.ExecuteTemplate(w, "orgpage.html", nil)
}

//GetStatic - Получение статики
func GetStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	vars := mux.Vars(r)
	key := vars["icon"]
	var url = "public/images/" + key
	http.ServeFile(w, r, url)
}
