package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	tTemplate "text/template"

	_ "github.com/lib/pq"
)

const (
	DbUser     = "postgres"
	DbPassword = "postgres"
	DbName     = "test"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()   // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello astaxie123!") // send data to client side
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		_ = t.Execute(w, nil)

		psql()
	} else {
		_ = r.ParseForm()

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		//template.HTMLEscape(w, []byte(r.Form.Get("username")))

		t, _ := tTemplate.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		_ = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	psql()
}

func psql() {
	fmt.Println("test123")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# inserting values")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
