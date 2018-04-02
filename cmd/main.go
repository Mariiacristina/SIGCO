package main

import(
  "net/http"
  "log"
  "SIGCO/API/rutas"
  "github.com/gorilla/handlers"
)


func main(){
    headers := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
    methods := handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"})

    m := rutas.Routering()
    log.Fatal(http.ListenAndServe(":8000",handlers.CORS(headers,methods,origins)(m)))

}
