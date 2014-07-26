package main

import "log"
import "net/http"
import "flag"

func init() {
  log.Println("Initializing ...")
}

func main() {

  var folder = flag.String("f", "./", "root folder to serve files from ...")

  flag.Parse()

  log.Println("Starting server ...")
  log.Println("Serving files from", *folder)

  http.ListenAndServe(":8081", http.FileServer(http.Dir(*folder)))
}
