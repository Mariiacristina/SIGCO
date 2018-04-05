package controller

import(
  "SIGCO/API/model"
  "net/http"
  "log"
  "SIGCO/API/schema"
  "encoding/json"
)

func GetAllVecinos(w http.ResponseWriter, r *http.Request){
  log.Println("get VECINOS")
}


//insertar marciano
func PostVecino(w http.ResponseWriter, r *http.Request){
  var vecino schema.Vecino
  _=json.NewDecoder(r.Body).Decode(&vecino)
  err := model.InsertVecino(vecino)
  if (err != nil) {
    w.WriteHeader(http.StatusInternalServerError)
  }else{
    w.WriteHeader(http.StatusOK)
  }
}
