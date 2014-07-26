package main

import "log"
import "io"
import "net/http"

func main() {
  log.Println("Starting server ...")

  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
    log.Println("new request from", req.RemoteAddr, "at", req.RequestURI)
    io.WriteString(w, "hello, world!\n")
  })

  err := http.ListenAndServe(":8080", nil)

  if err != nil {
    log.Fatal("ERROR: ", err)
  }
}
