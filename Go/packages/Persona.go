package packages

type Persona struct {
	Nombre string
	Edad   int
	Altura float64
}

func (p Persona) Saludar() string {
	return "Hola, mi nombre es " + p.Nombre
}
