package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

type User struct {
	Username string
	Password string
}

const (
	key            = `123456`
	ipRadiusServer = "192.168.3.128:9090"
)

func getlogin(rw http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("template.html"))
	tmp.Execute(rw, nil)
}
func getresult(rw http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("template.html"))
	tmp.Execute(rw, nil)
}
func postlogin(rw http.ResponseWriter, r *http.Request) {
	user := User{
		Username: r.FormValue("username"),
		Password: r.FormValue("passw"),
	}
<<<<<<< HEAD
=======
	fmt.Println(user.Username)
	fmt.Println(user.Password)
>>>>>>> 971846b0fa21010856e8758d0a851c1f5dc58b3f
	packet := radius.New(radius.CodeAccessRequest, []byte(key))
	rfc2865.UserName_SetString(packet, user.Username)
	rfc2865.UserPassword_SetString(packet, user.Password)

<<<<<<< HEAD
=======
		rfc2865.UserPassword_SetString(packet, user.Password)
	
	//res, er := radius.
	//radius.
>>>>>>> 971846b0fa21010856e8758d0a851c1f5dc58b3f
	response, err := radius.Exchange(context.Background(), packet, ipRadiusServer)

	if err != nil {
		log.Fatal(err)
	}
	var Result string
	if response.Code == radius.CodeAccessAccept {
		Result = "Đăng nhập thành công"
	} else {
		Result = "Đăng nhập thất bại"
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(Result)
	//log.Println("Code:", response.Code)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", getlogin).Methods("GET")
	r.HandleFunc("/login", postlogin).Methods("POST")

	http.ListenAndServe(":8080", r)
}
