package controller

import(
  "SIGCO/API/model"
  "net/http"
  "log"
  "SIGCO/API/schema"
  "encoding/json"
)

func GetAllFichas(w http.ResponseWriter, r *http.Request){
  log.Println("get FICHAS")
}
//insertar marciano
func PostFicha(w http.ResponseWriter, r *http.Request){
  var ficha schema.Ficha
  _=json.NewDecoder(r.Body).Decode(&ficha)
  err := model.InsertFicha(ficha)
  if (err != nil) {
    w.WriteHeader(http.StatusInternalServerError)
  }else{
    w.WriteHeader(http.StatusOK)
  }
}

func UpdateFicha(w http.ResponseWriter, r *http.Request){
  log.Println("UPDATEFICHA")
}
