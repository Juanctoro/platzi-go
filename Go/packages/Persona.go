package packages

type Persona struct {
	Nombre string
	Edad   int
	Altura float64
}

func (p Persona) Saludar() string {
	return "Hola, ni nombre es " + p.Nombre
}
