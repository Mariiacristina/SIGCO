package model

import(
  "SIGCO/API/connection"
  "SIGCO/API/schema"
  "log"
)
var(
  errmod error
)
func InsertFicha(ficha schema.Ficha)(errmod error){
  db := connection.Connect()
  _,errmod = db.Exec("INSERT INTO Marciano VALUES (?,?,?,?,?,?)", ficha.Id_visita, ficha.Id_vecino, ficha.Descripcion, ficha.Costo,ficha.Nombre_Admin,ficha.Pagado)
    if(errmod != nil) {
      log.Println("error en el modelo")
      log.Println(errmod)
    }else{
      log.Println("Se insertó con exito")
    }
    connection.Disconnect(db)
    return errmod
}
