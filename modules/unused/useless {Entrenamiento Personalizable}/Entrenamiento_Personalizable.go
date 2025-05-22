package entrenamiento

/*
type TipEdicion struct {
    ID       int
    Contenido string
}

func (t *TipEdicion) MostrarTip() string {
    return t.Contenido
}

type Ejercicio struct {
    ID           int
    Nombre       string
    ZonaMuscular string
    Repeticiones int
    Series       int
}

func (e *Ejercicio) Editar(nombre, zonaMuscular string, repeticiones, series int) {
    e.Nombre = nombre
    e.ZonaMuscular = zonaMuscular
    e.Repeticiones = repeticiones
    e.Series = series
}

type Rutina struct {
    ID            int
    Nombre        string
    Objetivo      string
    Duracion      int // en minutos
    TiempoDescanso int // en segundos
    Ejercicios    []Ejercicio
    TipsEdicion   []TipEdicion
}

func (r *Rutina) AgregarEjercicio(ejercicio Ejercicio) {
    r.Ejercicios = append(r.Ejercicios, ejercicio)
}

func (r *Rutina) EliminarEjercicio(ejercicioID int) {
    for i, e := range r.Ejercicios {
        if e.ID == ejercicioID {
            r.Ejercicios = append(r.Ejercicios[:i], r.Ejercicios[i+1:]...)
            break
        }
    }
}

func (r *Rutina) AjustarDuracion(nuevaDuracion int) {
    r.Duracion = nuevaDuracion
}

func (r *Rutina) AjustarDescanso(nuevoDescanso int) {
    r.TiempoDescanso = nuevoDescanso
}

type Usuario struct {
    ID      int
    Nombre  string
    Rutinas []Rutina
}

func (u *Usuario) CrearRutina(nombre string) *Rutina {
    rutina := Rutina{
        ID:     len(u.Rutinas) + 1,
        Nombre: nombre,
    }
    u.Rutinas = append(u.Rutinas, rutina)
    return &u.Rutinas[len(u.Rutinas)-1]
}

func (u *Usuario) EditarRutina(rutinaID int, nuevaRutina Rutina) {
    for i, r := range u.Rutinas {
        if r.ID == rutinaID {
            u.Rutinas[i] = nuevaRutina
            break
        }
    }
}

func (u *Usuario) VerRutinas() []Rutina {
    return u.Rutinas
}
*/
