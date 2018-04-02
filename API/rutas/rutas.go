package rutas

import(
  "net/http"
  "github.com/gorilla/mux"
  "SIGCO/API/controller"
)

func Routering()(http.Handler){
  router := mux.NewRouter()

//vecinosss
  router.HandleFunc("/Vecino", controller.GetAllVecinos).Methods("GET")
  router.HandleFunc("/Vecino", controller.PostVecino).Methods("POST")
  //router.HandleFunc("/Vecino/{id}", controler.DeleteVecino).Methods("DELETE")

//fichas
    router.HandleFunc("/Ficha", controller.GetAllFichas).Methods("GET")
    router.HandleFunc("/Ficha", controller.PostFicha).Methods("POST")
    router.HandleFunc("/Ficha/{id}", controller.UpdateFicha).Methods("PUT")

//FACTURAS
  router.HandleFunc("/Pagar", controller.PostFactura).Methods("POST")

  return router
}
