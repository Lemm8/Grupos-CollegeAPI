package helpers

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}

type Grupo struct {
	ID         int    `json:"id"`
	Salon      string `json:"salon"`
	Materia_ID int    `json:"materia_id"`
	Docente_ID int    `json:"docente_id"`
	Alumno_ID  int    `json:"alumno_id"`
}

type Materia struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	IsTroncoComun bool   `json:"troncoComun"`
}

type Alumno struct {
	ID               int    `json:"id"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Matricula        string `json:"matricula"`
	Fecha_Nacimiento string `json:"fecha_nacimiento"`
	Semestre         string `json:"semestre"`
	Carreras_ID      int    `json:"carreras_id"`
}

type Docente struct {
	ID               int    `json:"ID"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Matricula        string `json:"matricula"`
	Fecha_Nacimiento string `json:"fecha_nacimiento"`
	Titulo           string `json:"titulo"`
	Correo           string `json:"correo"`
	Telefono         string `json:"telefono"`
}

type ListGruposResponse struct {
	Status int      `json:"status"`
	Grupos []*Grupo `json:"docentes"`
}

type GrupoResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Grupo   *Grupo `json:"docente"`
}
