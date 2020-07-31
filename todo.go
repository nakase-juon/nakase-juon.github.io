package main

import (
        "io/ioutil"
        "fmt"
        "net/http"
)

var slice []string

func main() {
      //fmt.Println(slice)
      //slice = append(slice, a)
        http.HandleFunc("/todo",TodoHandler)
        http.HandleFunc("/", RootHandler)
        http.HandleFunc("/hello", HelloHandler)
        http.HandleFunc("/hoge",HogeHandler)
        http.ListenAndServe(":8080", nil)
}
func TodoHandler(w http.ResponseWriter, r *http.Request) {
        switch r.Method {

    case http.MethodGet:
        fmt.Fprint(w, "GETするよ!\n")
        for _, s := range slice{
          fmt.Fprintln(w, s)
        }

    case http.MethodPost:
        fmt.Fprint(w, "POSTしたよ!\n")
        a, err := ioutil.ReadAll(r.Body)
        if err != nil{
          fmt.Fprint(w, "AHO\n")
        }

        //aをちゃんとした形にする

        slice = append(slice, string(a))

    default:
        fmt.Fprint(w, "Method not allowed.\n")
    }
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Print("アクセスされたよ\n")
        fmt.Fprint(w, "アクセスされたよ")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello World")
}

func HogeHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "hoge")
}

