package main

import (
  "fmt"
  "log"
  "net/http"
)
const msg = "My proverb-native"
func main() {
  fmt.Println(msg)


  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, msg)
  })

  fmt.Println(1)
  log.Fatal(http.ListenAndServe(":8081",nil))

  fmt.Println(2)
}
