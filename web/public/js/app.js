const loginPage = document.getElementById('login');
const registerPage = document.getElementById('register');
const recordatoriosPage = document.getElementById('recordatorios');
const loginForm = document.getElementById('login-form');
const registerForm = document.getElementById('register-form');
const showRegister = document.getElementById('show-register');
const showLogin = document.getElementById('show-login');
const loginError = document.getElementById('login-error');
const registerError = document.getElementById('register-error');
const logoutBtn = document.getElementById('logout-btn');
const reminderError = document.getElementById('reminder-error');
const listComponent = document.querySelector('editable-list');

let isLoggedIn = false;
let userRole = "user";

// --- Navigation ---
function showPage(pageId) {
    [loginPage, registerPage, recordatoriosPage].forEach(p => p.style.display = 'none');
    document.getElementById(pageId).style.display = 'block';
}
window.addEventListener('hashchange', () => {
    if (isLoggedIn) {
        showPage('recordatorios');
    } else {
        showPage('login');
    }
});
showPage('login');

// --- Auth ---
showRegister.onclick = e => {
    e.preventDefault();
    showPage('register');
};
showLogin.onclick = e => {
    e.preventDefault();
    showPage('login');
};

loginForm.onsubmit = async e => {
    e.preventDefault();
    loginError.textContent = '';
    const username = document.getElementById('login-username').value;
    const password = document.getElementById('login-password').value;
    try {
        const res = await fetch('/api/auth/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });
        if (res.ok) {
            isLoggedIn = true;
            await fetchUserRole();
            showPage('recordatorios');
            loadReminders();
            logoutBtn.style.display = '';
        } else {
            loginError.textContent = 'Credenciales incorrectas';
        }
    } catch {
        loginError.textContent = 'Error de red';
    }
};

registerForm.onsubmit = async e => {
    e.preventDefault();
    registerError.textContent = '';
    const username = document.getElementById('register-username').value;
    const password = document.getElementById('register-password').value;
    const role = document.getElementById('register-role').value;
    try {
        const res = await fetch('/api/auth/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password, role })
        });
        if (res.ok) {
            showPage('login');
        } else {
            registerError.textContent = 'Error al registrar usuario';
        }
    } catch {
        registerError.textContent = 'Error de red';
    }
};

logoutBtn.onclick = async () => {
    await fetch('/api/auth/logout', { method: 'POST' });
    isLoggedIn = false;
    userRole = "user";
    showPage('login');
    logoutBtn.style.display = 'none';
};

// --- Fetch user role after login ---
async function fetchUserRole() {
    // This assumes your backend exposes /api/auth/me or similar.
    // If not, you can store the role in a cookie or return it in login.
    // For now, fallback to "user".
    userRole = "user";
}

// --- Reminders CRUD ---
async function loadReminders() {
    try {
        const res = await fetch('/api/v1/reminders');
        if (res.ok) {
            const data = await res.json();
            listComponent.setData(data);
        } else {
            reminderError.textContent = 'No se pudieron cargar los recordatorios';
        }
    } catch {
        reminderError.textContent = 'Error de red';
    }
}

listComponent.addEventListener('item-create', async e => {
    reminderError.textContent = '';
    try {
        const res = await fetch('/api/v1/reminders', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                titulo: e.detail.titulo,
                descripcion: e.detail.descripcion,
                fecha: e.detail.fecha,
                cumplido: false
            })
        });
        if (res.ok) {
            loadReminders();
        } else {
            reminderError.textContent = 'No tienes permisos para crear';
        }
    } catch {
        reminderError.textContent = 'Error de red';
    }
});

listComponent.addEventListener('item-update', async e => {
    reminderError.textContent = '';
    if (userRole !== "admin") {
        reminderError.textContent = 'Solo el administrador puede editar';
        return;
    }
    try {
        const res = await fetch(`/api/v1/reminders/${e.detail.id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(e.detail)
        });
        if (res.ok) {
            loadReminders();
        } else if (res.status === 403) {
            reminderError.textContent = 'No tienes permisos para editar';
        } else {
            reminderError.textContent = 'Error al editar';
        }
    } catch {
        reminderError.textContent = 'Error de red';
    }
});

listComponent.addEventListener('item-delete', async e => {
    reminderError.textContent = '';
    if (userRole !== "admin") {
        reminderError.textContent = 'Solo el administrador puede eliminar';
        return;
    }
    try {
        const res = await fetch(`/api/v1/reminders/${e.detail.id}`, {
            method: 'DELETE'
        });
        if (res.ok || res.status === 204) {
            loadReminders();
        } else if (res.status === 403) {
            reminderError.textContent = 'No tienes permisos para eliminar';
        } else {
            reminderError.textContent = 'Error al eliminar';
        }
    } catch {
        reminderError.textContent = 'Error de red';
    }
});

// --- On login, load reminders ---
if (isLoggedIn) {
    loadReminders();
}