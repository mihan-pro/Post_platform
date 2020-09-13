package Helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofrs/uuid"
)

//ShowError dsdjl b htlbhtrn yf cnhfybwe jib,rb
func ShowModule(w http.ResponseWriter, number string) {
	errortpl, err := ioutil.ReadFile("public/modules/" + number + ".html")
	if err != nil {
		fmt.Print(err)
	}

	switch number {
	case "404":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errortpl))
	case "500":
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errortpl))
	case "405":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(errortpl))
	}
}

//CheckErr Дописать обработчик ошибок
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//IsValidUUID asdsad
func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
