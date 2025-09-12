package model

type Developer struct {
	Tarea        string
	Comentarios  string
	IdTarea      string
	Cliente      string
	Prioridad    string
	Estado       string
	Estadokanvan string
	FechaRe      string
	FechaIni     string
	FechaFin     string
	DiasDesarro  string
	Developer    string
	Url          string
	Link         string
}

func NewDeveloper(data []string, link string) Developer {
	return Developer{
		Tarea:        data[0],
		Cliente:      data[1],
		Prioridad:    data[2],
		Estado:       data[3],
		Estadokanvan: data[4],
		FechaRe:      data[5],
		FechaIni:     data[6],
		FechaFin:     data[7],
		DiasDesarro:  data[8],
		Developer:    data[9],
		Link:         link,
	}
}
