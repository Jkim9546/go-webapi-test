package main
 
import (
    "net/http"
)
 
func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })
    http.HandleFunc("/sang", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello sang"))
    })
    http.HandleFunc("/jung", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello jung"))
    })
 
    http.ListenAndServe(":5000", nil)
}
