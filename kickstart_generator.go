package main

import (
        "fmt"
        "html/template"
        "io/ioutil"
        "log"
        "net/http"

        "github.com/gorilla/mux"
)

func main() {
        router := mux.NewRouter().StrictSlash(true)
        router.HandleFunc("/", Index)
        router.HandleFunc("/status", Status)
        router.HandleFunc("/ks_generate/", ks_generate)
        log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to Kickstart File Generator")
}

func Status(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "API is up and running")
}

func check_error(err error) {
        if err != nil {
        panic(err)
	}
}

func ks_generate(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        check_error(err)

        tmpl_file := "ks.tmpl"
        tp, err := ioutil.ReadFile(tmpl_file)
        check_error(err)

        data := make(map[string]string)
        for i, j := range r.Form {
                data[i] = j[0]
        }
        t, err := template.New("index").Parse(string(tp))
        check_error(err)

        err = t.Execute(w, data)
        check_error(err)
}
