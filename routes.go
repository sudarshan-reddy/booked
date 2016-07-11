package routes

import ("net/http"
        "fmt"
		"html/template")

func login(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.gtpl")
		t.Execute(w, nil)
	}else{
		r.ParseForm()

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}



