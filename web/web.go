package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	tTemplate "text/template"
	"time"

	_ "github.com/lib/pq"
)

const (
	DbUser     = "postgres"
	DbPassword = "mysecretpassword"
	DbName     = "test"
	DbPort     = "10000"
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

	connStr := "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	dbinfo := fmt.Sprintf(connStr, DbUser, DbPassword, DbPort, DbName)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# inserting values")

	var lastInsertId int
	err = db.QueryRow("INSERT INTO test_schema.userinfo(username,departname, created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)

	fmt.Println("last insert id=", lastInsertId)

	fmt.Println("# updating")
	stmt, err := db.Prepare("update test_schema.userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("astaxie_updated", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "Rows changed")
	fmt.Println("# Querying")

	rows, err := db.Query("select * from test_schema.userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)

		fmt.Println("uid | username | department | created")
		fmt.Printf("%3v | %18v | %6v | %6v\n", uid, username, department, created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
