package main

import("net/http"
	   "booked/routes"
)

func main(){
	http.HandleFunc("/login" , routes.Login)
	http.ListenAndServe(":8001" , nil)
}
