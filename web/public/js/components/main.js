if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        navigator.serviceWorker.register('/js/components/sw.js')
            .then(registration => {
                console.log('Service Worker registered with scope:', registration.scope);
            }).catch(error => {
                console.log('Service Worker registration failed:', error);
            });
    });
}

class ReminderCard extends HTMLElement {
    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.innerHTML = `
            <style>
                .grid {
                    display: grid;
                    grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
                    gap: 16px;
                    margin-top: 20px;
                }
                .reminder {
                    background: #fffbe7;
                    border: 1px solid #e0c97f;
                    border-radius: 10px;
                    padding: 16px;
                    box-shadow: 0 2px 8px rgba(0,0,0,0.08);
                }
                .reminder h3 {
                    margin: 0 0 8px 0;
                    font-size: 1.1em;
                }
                .reminder p {
                    margin: 0;
                    font-size: 0.95em;
                }
                .add-form {
                    display: flex;
                    gap: 8px;
                    margin-bottom: 16px;
                    flex-wrap: wrap;
                }
                .add-form input, .add-form textarea {
                    font-size: 1em;
                    padding: 4px 8px;
                    border-radius: 4px;
                    border: 1px solid #ccc;
                }
                .add-form button {
                    background: chocolate;
                    color: white;
                    border: none;
                    border-radius: 4px;
                    padding: 6px 14px;
                    cursor: pointer;
                }
                .add-form button:hover {
                    background: #d2691e;
                }
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