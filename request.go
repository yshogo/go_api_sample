package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  
  "github.com/julienschmidt/httprouter"
)

func main() {
  router := httprouter.New()
  router.GET("/", requestJSON)

  log.Fatal(http.ListenAndServe(":8080", nil))
}


func requestJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://firebasestorage.googleapis.com/v0/b/blog-80835.appspot.com/o/test.json?alt=media&token=401dc208-99f8-4a27-a30c-a288ebf38ca7", nil)

  if (err != nil) {
    fmt.Println("error occur")
  }

  resp, err := client.Do(req)

  if(err != nil) {
    fmt.Println("error occur")
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if(err != nil) {
    fmt.Println("error occur")
  }

  fmt.Fprintf(w, string(body))
}
