package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "VERSION: %s \n", os.Getenv("VERSION"))
    fmt.Fprintf(w, "BRANCH: %s", os.Getenv("BRANCH"))
}
