package main

import "log"
import "io"
import "io/ioutil"
import "net/http"
import "math/rand"
import "encoding/json"
import "flag"


func main()  {

  var fileName = flag.String("f", "./fortunes.json", "file containing fortunes as JSON")
  flag.Parse()

  log.Println("Starting server ...")

  data, err := ioutil.ReadFile(*fileName)

  if err != nil {
    log.Fatal(err)
  }


  var fortunes []string
  err = json.Unmarshal(data, &fortunes)

  if err != nil {
    log.Fatal(err)
  }


  http.HandleFunc("/", func(w http.ResponseWriter,req *http.Request){
    i := rand.Intn(len(fortunes))
    io.WriteString(w, fortunes[i])
  })

  http.ListenAndServe(":8081", nil)
}
