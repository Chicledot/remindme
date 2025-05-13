package user

import "time"

type Categoria struct {
	ID     int
	Nombre string
}

type Recordatorio struct {
	ID          int
	Titulo      string
	Descripcion string
	Fecha       time.Time
	Cumplido    bool
	Categoria   *Categoria
}

func (r *Recordatorio) MarcarCumplido() {
	r.Cumplido = true
}

func (r *Recordatorio) Categorizar(categoria *Categoria) {
	r.Categoria = categoria
}

type Notificacion struct {
	ID         int
	Mensaje    string
	FechaEnvio time.Time
}

func (n *Notificacion) Enviar(usuario *Usuario) {
	// Implementar lógica de envío de notificación al usuario
}

type Usuario struct {
	ID             int
	Nombre         string
	Email          string
	Recordatorios  []*Recordatorio
	Notificaciones []*Notificacion
}

func (u *Usuario) CrearRecordatorio(titulo, descripcion string, fecha time.Time) *Recordatorio {
	rec := &Recordatorio{
		ID:          len(u.Recordatorios) + 1,
		Titulo:      titulo,
		Descripcion: descripcion,
		Fecha:       fecha,
		Cumplido:    false,
	}
	u.Recordatorios = append(u.Recordatorios, rec)
	return rec
}

func (u *Usuario) VerRecordatorios() []*Recordatorio {
	return u.Recordatorios
}

func (u *Usuario) MarcarComoCumplido(recordatorio *Recordatorio) {
	recordatorio.MarcarCumplido()
}
