package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  //"reflect"
)

type Page struct {
  Title string
  Body  []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)

  if err == nil {
    return &Page{Title: title, Body: body}, nil
  } else {
    return nil, err
  }
}

func router(response http.ResponseWriter, req *http.Request) {
  word := req.URL.Path[1:]
  loadedPage, err := loadPage(word)

  if err != nil {
    fmt.Fprintf(response, "<h1>Sanskrit dictionary</h1><div>Sorry. I do not know the word: <strong>%s</strong></div>", word)
  } else {
    fmt.Fprintf(response, "<h1>Sanskrit dictionary</h1><div>%s</div>", loadedPage.Body)
  }
}

func main() {
  port := 8080
  fmt.Println(fmt.Sprintf("Sanskrit dictionary is listening on port %d", port))

  http.HandleFunc("/", router)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
