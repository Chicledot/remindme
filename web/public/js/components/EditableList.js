class EditableList extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.data = [];
        this.editingId = null;
        this.render();
    }

    setData(data) {
        this.data = data;
        this.render();
    }

    render() {
        this.shadowRoot.innerHTML = `
            <style>
                table { width: 100%; border-collapse: collapse; margin-top: 1em; }
                th, td { border: 1px solid #ccc; padding: 8px; text-align: left; }
                th { background: #f3f3f3; }
                button { margin: 0 2px; }
                form { margin-top: 1em; }
            </style>
            <table>
                <thead>
                    <tr>
                        <th>Título</th>
                        <th>Descripción</th>
                        <th>Fecha</th>
                        <th>Cumplido</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    ${this.data.map(item => this.editingId === item.id ? this.editRow(item) : this.displayRow(item)).join('')}
                </tbody>
            </table>
            ${this.editingId ? '' : this.createForm()}
        `;

        // Add listeners for edit/delete/create/update
        this.shadowRoot.querySelectorAll('.delete-btn').forEach(btn => {
            btn.onclick = e => {
                const id = parseInt(btn.dataset.id);
                this.dispatchEvent(new CustomEvent('item-delete', { detail: { id } }));
            };
        });
        this.shadowRoot.querySelectorAll('.edit-btn').forEach(btn => {
            btn.onclick = e => {
                this.editingId = parseInt(btn.dataset.id);
                this.render();
            };
        });
        const createForm = this.shadowRoot.querySelector('#create-form');
        if (createForm) {
            createForm.onsubmit = e => {
                e.preventDefault();
                const titulo = createForm.titulo.value.trim();
                const descripcion = createForm.descripcion.value.trim();
                const fecha = createForm.fecha.value;
                if (titulo && descripcion && fecha) {
                    this.dispatchEvent(new CustomEvent('item-create', {
                        detail: { titulo, descripcion, fecha }
                    }));
                    createForm.reset();
                }
            };
        }
        const updateForm = this.shadowRoot.querySelector('#update-form');
        if (updateForm) {
            updateForm.onsubmit = e => {
                e.preventDefault();
                const id = parseInt(updateForm.id.value);
                const titulo = updateForm.titulo.value.trim();
                const descripcion = updateForm.descripcion.value.trim();
                const fecha = updateForm.fecha.value;
                const cumplido = updateForm.cumplido.checked;
                this.dispatchEvent(new CustomEvent('item-update', {
                    detail: { id, titulo, descripcion, fecha, cumplido }
                }));
                this.editingId = null;
            };
            this.shadowRoot.querySelector('#cancel-edit').onclick = () => {
                this.editingId = null;
                this.render();
            };
        }
    }

    displayRow(item) {
        return `
            <tr>
                <td>${item.titulo}</td>
                <td>${item.descripcion}</td>
                <td>${item.fecha ? item.fecha.split('T')[0] : ''}</td>
                <td>${item.cumplido ? '✅' : '❌'}</td>
                <td>
                    <button class="edit-btn" data-id="${item.id}">✏️ Editar</button>
                    <button class="delete-btn" data-id="${item.id}">🗑️ Eliminar</button>
                </td>
            </tr>
        `;
    }

    editRow(item) {
        return `
            <tr>
                <form id="update-form">
                    <input type="hidden" name="id" value="${item.id}">
                    <td><input name="titulo" value="${item.titulo}" required></td>
                    <td><input name="descripcion" value="${item.descripcion}" required></td>
                    <td><input name="fecha" type="date" value="${item.fecha ? item.fecha.split('T')[0] : ''}" required></td>
                    <td><input name="cumplido" type="checkbox" ${item.cumplido ? 'checked' : ''}></td>
                    <td>
                        <button type="submit">💾 Actualizar</button>
                        <button type="button" id="cancel-edit">Cancelar</button>
                    </td>
                </form>
            </tr>
        `;
    }

    createForm() {
        return `
            <form id="create-form">
                <input name="titulo" placeholder="Título" required>
                <input name="descripcion" placeholder="Descripción" required>
                <input name="fecha" type="date" required>
                <button type="submit">➕ Crear</button>
            </form>
        `;
    }
}

customElements.define('editable-list', EditableList);