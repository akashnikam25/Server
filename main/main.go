package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", loggingMiddleWare(http.HandlerFunc(handleGetRequest)))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	fmt.Println("c value is ", queryValues["c"])
	fmt.Println("name value is ", queryValues["name"])
	counter := r.Header.Get("counter")
	fmt.Println("counter ", counter)
	val, _ := strconv.Atoi(counter)
	sum := getCalSum(val)
	fmt.Fprintf(w, "sum of this counter is %d", sum)
}

func getCalSum(counter int) (res int) {
	for i := 1; i < counter; i++ {
		res += i
	}
	return
}
