package calorias

/*
import "time"

type Alimento struct {
	ID                 int
	Nombre             string
	CaloriasPorPorcion float64
}

type Ejercicio struct {
	ID                        int
	Nombre                    string
	CaloriasQuemadasPorMinuto float64
}

type RegistroAlimento struct {
	Fecha           time.Time
	Cantidad        float64
	CaloriasTotales float64
	Alimento        *Alimento
}

type RegistroEjercicio struct {
	Fecha            time.Time
	Duracion         int // en minutos
	CaloriasQuemadas float64
	Ejercicio        *Ejercicio
}

type ResumenDiario struct {
	Fecha              time.Time
	CaloriasConsumidas float64
	CaloriasQuemadas   float64
	CaloriasNetas      float64
}

func (r *ResumenDiario) CalcularNetas() float64 {
	r.CaloriasNetas = r.CaloriasConsumidas - r.CaloriasQuemadas
	return r.CaloriasNetas
}

type Usuario struct {
	ID                 int
	Nombre             string
	Peso               float64
	Altura             float64
	RegistrosAlimento  []RegistroAlimento
	RegistrosEjercicio []RegistroEjercicio
	ResumenesDiarios   []ResumenDiario
}

func (u *Usuario) IngresarPeso(peso float64) {
	u.Peso = peso
}

func (u *Usuario) IngresarAltura(altura float64) {
	u.Altura = altura
}

func (u *Usuario) RegistrarAlimento(alimento *Alimento, cantidad float64) {
	calorias := alimento.CaloriasPorPorcion * cantidad
	reg := RegistroAlimento{
		Fecha:           time.Now(),
		Cantidad:        cantidad,
		CaloriasTotales: calorias,
		Alimento:        alimento,
	}
	u.RegistrosAlimento = append(u.RegistrosAlimento, reg)
}

func (u *Usuario) RegistrarEjercicio(ejercicio *Ejercicio, duracion int) {
	quemadas := ejercicio.CaloriasQuemadasPorMinuto * float64(duracion)
	reg := RegistroEjercicio{
		Fecha:            time.Now(),
		Duracion:         duracion,
		CaloriasQuemadas: quemadas,
		Ejercicio:        ejercicio,
	}
	u.RegistrosEjercicio = append(u.RegistrosEjercicio, reg)
}

func (u *Usuario) VerResumenDiario(fecha time.Time) ResumenDiario {
	var consumidas, quemadas float64
	for _, reg := range u.RegistrosAlimento {
		if sameDay(reg.Fecha, fecha) {
			consumidas += reg.CaloriasTotales
		}
	}
	for _, reg := range u.RegistrosEjercicio {
		if sameDay(reg.Fecha, fecha) {
			quemadas += reg.CaloriasQuemadas
		}
	}
	resumen := ResumenDiario{
		Fecha:              fecha,
		CaloriasConsumidas: consumidas,
		CaloriasQuemadas:   quemadas,
	}
	resumen.CalcularNetas()
	u.ResumenesDiarios = append(u.ResumenesDiarios, resumen)
	return resumen
}

// Utilidad para comparar fechas por día
func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}
*/
