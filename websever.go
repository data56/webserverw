package main

import ("fmt"
  "net/http")

func main() {



    http.HandleFunc("/earth", handler2)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    		http.ServeFile(w, r, r.URL.Path[1:])
    	})


    http.ListenAndServe(":8080", nil)


}

func handler(w http.ResponseWriter, r *http.Request){
      fmt.Fprintf(w, "Hello World") }

func handler2(w http.ResponseWriter, r *http.Request){
            fmt.Fprintf(w, "Hello earth")

            //hi

}
