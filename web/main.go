package main

import (
  "fmt"
  "html/template"
  "io/ioutil"
  "log"
  "net/http"
)

type Translation struct {
  English string
  Sanskrit  []byte
}

func (p *Translation) save() error {
  filename := p.English + ".txt"
  return ioutil.WriteFile(filename, p.Sanskrit, 0600)
}

func loadTranslation(english string) (*Translation, error) {
  filename := english + ".txt"
  sanskrit, err := ioutil.ReadFile(filename)

  if err == nil {
    return &Translation{English: english, Sanskrit: sanskrit}, nil
  } else {
    return nil, err
  }
}

func indexHandler(response http.ResponseWriter, req *http.Request) {
  t, _ := template.ParseFiles("templates/index.html")
  t.Execute(response, nil)
}

func english2SanskritHandler(response http.ResponseWriter, req *http.Request) {
  word, ok := req.URL.Query()["english"]

  if ok {
    loadedTranslation, err := loadTranslation(word[0])

    if err == nil {
      t, _ := template.ParseFiles("templates/english2sanskrit.html")
      t.Execute(response, loadedTranslation)
    } else {
      t, _ := template.ParseFiles("templates/word_not_found.html")
      translation := Translation{English: word[0], Sanskrit: nil}
      t.Execute(response, &translation)
    }

  } else {
    indexHandler(response, req)
  }
}

func main() {
  port := 8080
  fmt.Println(fmt.Sprintf("Sanskrit dictionary is listening on port %d", port))

  http.HandleFunc("/english2sanskrit", english2SanskritHandler)
  http.HandleFunc("/", indexHandler)

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
