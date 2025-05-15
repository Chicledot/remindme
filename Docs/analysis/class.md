# Modelado de Clases y POO:

Comenzamos por reunir todos los archivos y descripciones, y lo alimentamos a Copilot, ocupando el modulo GTP4, le entregamos el contexto necesario y los parametros para realizar el trabajo, el prompt usado fue el siguiente:

"Basado en los siguientes casos de uso del sistema de la app RemindME descritos en usecases.md, genera un diagrama de clases en formato PlantUML. Identifica las clases principales, sus atributos (usando tipos int, string, bool donde sea apropiado) y métodos básicos relacionados con los casos de uso. Considera las relaciones entre clases (asociación, composición, etc.)"

La IA procedió a generar codigo en lenguaje PlantUML, el cual es el que se ocupó para realizar el diseño presente en los archivos siguientes:

# Diagramas de Clase :

## [Diagrama de clases - Recordatorio]


![Image](/Docs/analysis/diagrams/class/ClassDiagram_Reminder.svg)


## [Diagrama de clases - Seguimiento de Calorias]

![Image](/Docs/analysis/diagrams/class/ClassDiagram_SeguimientoCalorias.svg)



## [Diagrama de clases - Conectar dispositivo bluetooth]

![Image](/Docs/analysis/diagrams/class/ClassDiagram_Sincronizacion.svg)

## [Diagrama de clases - Entrenamiento personalizado]

![Image](/Docs/analysis/diagrams/class/ClassDiagram_EntrenamientoPersonalizable.svg)

# Codigo Creado: Implementacion de Structs en lenguago GO

Recordatorio

**`modules/RF01/Reminder.go`**
```go

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
```

# Reflexión sobre el uso de la IA:

Siempre que ocupo la IA siento que estoy haciendo trampa, siento que aprendo poco, por lo tanto prefiero evitarla lo más posible. No me gusta defender la IA, ultimamente se a demostrado que gracias a la IA la gente esta perdiendo la capacida de pensar de manera critica, aun asi, mucho de ese problema se puede atribuir a la gente siendo muy floja y ocupando mál la herramienta para cosas que no deberia de hacer. Como herramienta para el developer esta bien, pero para que llegue a hacer el trabajo por mi, no me agrada para nada. -Diego (Junior dev)



