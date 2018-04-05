package schema

type Vecino struct {
  Rut string `json:"rut, omitempty"`
  Nombre string `json: "nombre, omitempty"`
  Apellido string `json: "apellido, omitempty"`
  Direccion string `json: "direccion, omitempty"`
  C_Postal int `json: "c_postal, omitempty"`
  Ciudad string `json: "ciudad, omitempty"`
  Tel int `json: "tel, omitempty"`
}

type Ficha struct {
  Id_visita int `json:"id_visita, omitempty"`
  Descripcion string `json: "descripcion, omitempty"`
  Costo int `json: "costo, omitempty"`
  Nombre_Admin string `json: "nombre_admin, omitempty"`
  Pagado int `json: "pagado, omitempty"`
}
