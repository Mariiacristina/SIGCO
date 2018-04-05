package model

import(
  "SIGCO/API/connection"
  "SIGCO/API/schema"
  "log"
)
var(
  errmod error
)
func InsertVecino(vecino schema.Vecino)(errmod error){
  db := connection.Connect()
  _,errmod = db.Exec("INSERT INTO VECINO VALUES (?,?,?,?,?,?,?)", vecino.Rut, vecino.Nombre, vecino.Apellido, vecino.Direccion, vecino.C_Postal, vecino.Ciudad, vecino.Tel)
    if(errmod != nil) {
      log.Println("error en el modelo")
      log.Println(errmod)
    }else{
      log.Println("Se insertó con exito")
    }
    connection.Disconnect(db)
    return errmod
}
