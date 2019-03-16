package main

import (
  "io/ioutil"
  "net/http"
  
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  e.GET("/", requestJSON)

  e.Logger.Fatal(e.Start(":8080"))
}


func requestJSON(c echo.Context) error {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://firebasestorage.googleapis.com/v0/b/blog-80835.appspot.com/o/test.json?alt=media&token=401dc208-99f8-4a27-a30c-a288ebf38ca7", nil)

  if (err != nil) {
    return c.String(http.StatusInternalServerError, "err")
  }

  resp, err := client.Do(req)

  if(err != nil) {
    return c.String(http.StatusInternalServerError, "err")
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if(err != nil) {
    return c.String(http.StatusInternalServerError, "err")
  }

  return c.String(http.StatusOK, string(body))
}
