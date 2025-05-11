# Documentación de Ruteo y Componentes

## Lista de Rutas

| Ruta (Hash)      | Sección / Página           | Descripción                        |
|------------------|---------------------------|------------------------------------|
| `#principal`     | Página Principal          | Página de bienvenida               |
| `#pagina2`       | Recordatorios activos     | Gestión y visualización de recordatorios |
| `#pagina3`       | Página 3                  | Otra sección de contenido          |
| `#acerca`        | Acerca de nosotros        | Información sobre el proyecto      |

---

## Enfoque de Ruteo

**Tipo:** Ruteo con CSS usando `:target`  
**Justificación:**  
Se utiliza el selector CSS `:target` para mostrar/ocultar secciones de la SPA según el hash de la URL. Esto permite navegación instantánea entre "páginas" sin recargar el sitio, manteniendo la simplicidad y sin necesidad de JavaScript adicional para el ruteo.

**Código clave:**

```html
<!-- index.html -->
<ul>
    <li><a href="#principal">Inicio</a></li>
    <li><a href="#pagina2">Pagina2</a></li>
    <li><a href="#pagina3">Pagina3</a></li>
    <li style="float:right;"><a class="active" href="#acerca">Acerca de nosotros</a></li>
</ul>

<div id="principal" class="page">...</div>
<div id="pagina2" class="page">...</div>
<div id="pagina3" class="page">...</div>
<div id="acerca" class="page">...</div>
```

```css
/* style.css */
.page {
    display: none;
    transition: opacity 0.5s;
    opacity: 0;
}
.page:target {
    display: block;
    opacity: 1;
}
.page:first-of-type:target,
.page:first-of-type:not(:target) {
    display: block;
    opacity: 1;
}
```

---

## Web Component Creado

**Etiqueta:** `<reminder-card>`  
**Propósito:** Permitir al usuario crear y visualizar recordatorios en formato de tarjetas, agregando título y descripción, y mostrándolos en una cuadrícula.

**Código clave:**

```javascript
// main.js
class ReminderCard extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = `
            <style>
                .grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(220px, 1fr)); gap: 16px; margin-top: 20px; }
                .reminder { background: #fffbe7; border: 1px solid #e0c97f; border-radius: 10px; padding: 16px; box-shadow: 0 2px 8px rgba(0,0,0,0.08);}
                .reminder h3 { margin: 0 0 8px 0; font-size: 1.1em; }
                .reminder p { margin: 0; font-size: 0.95em; }
                .add-form { display: flex; gap: 8px; margin-bottom: 16px; flex-wrap: wrap; }
                .add-form input, .add-form textarea { font-size: 1em; padding: 4px 8px; border-radius: 4px; border: 1px solid #ccc; }
                .add-form button { background: chocolate; color: white; border: none; border-radius: 4px; padding: 6px 14px; cursor: pointer; }
                .add-form button:hover { background: #d2691e; }
            </style>
            <form class="add-form">
                <input type="text" placeholder="Título" required>
                <textarea placeholder="Descripción" rows="1" required></textarea>
                <button type="submit">Agregar</button>
            </form>
            <div class="grid"></div>
        `;
        this.grid = this.shadowRoot.querySelector('.grid');
        this.form = this.shadowRoot.querySelector('.add-form');
        this.form.addEventListener('submit', (e) => {
            e.preventDefault();
            const title = this.form.querySelector('input').value.trim();
            const desc = this.form.querySelector('textarea').value.trim();
            if (title && desc) {
                this.addReminder(title, desc);
                this.form.reset();
            }
        });
    }

    addReminder(title, desc) {
        const card = document.createElement('div');
        card.className = 'reminder';
        card.innerHTML = `<h3>${title}</h3><p>${desc}</p>`;
        this.grid.appendChild(card);
    }
}
customElements.define('reminder-card', ReminderCard);
```

**Uso en HTML:**

```html
<div id="pagina2" class="page">
    <h2>Recordatorios activos</h2>
    <reminder-card></reminder-card>
</div>
```

---

## Relación con Arquitecturas SPA

Este enfoque implementa una SPA (Single Page Application) básica, donde la navegación entre secciones no recarga la página. El ruteo se maneja con CSS y el contenido dinámico (recordatorios) se gestiona con Web Components nativos, manteniendo el código modular y reutilizable.

---

## Instrucciones de Prueba

1. Inicia el servidor Go (`go run main.go`) en la carpeta `web`.
2. Abre [http://localhost:8080/](http://localhost:8080/) en tu navegador.
3. Navega usando el menú superior:
   - `Inicio` (`#principal`)
   - `Recordatorios activos` (`#pagina2`)
   - `Página 3` (`#pagina3`)
   - `Acerca de nosotros` (`#acerca`)
4. En la sección "Recordatorios activos", agrega recordatorios usando el formulario y verifica que se muestren en tarjetas dentro de una cuadrícula.
5. Verifica que la navegación entre secciones es instantánea y no recarga la página.
6. Si tienes problemas, revisa la consola del navegador para errores.

---

## Enlaces a Issues de GitHub

- [Issue #11: Definir rutas de carpetas y ordenar Git](https://github.com/Chicledot/remindme/issues/11)
- [Issue #12: Crear e implementar PWA](https://github.com/Chicledot/remindme/issues/12)
- [Issue #13: Definir Rutas](https://github.com/Chicledot/remindme/issues/13)
- [Issue #14: Crear Web Component Nativo](https://github.com/Chicledot/remindme/issues/14)

---

**Autor:**
Diego  y Copilot.
Fecha: 2025-05-11
