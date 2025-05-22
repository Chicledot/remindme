package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	// User and Recordatorio types are defined below instead of importing from a missing package
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	Role         string `json:"role"`
}

var (
	users     = []User{}
	reminders = []Recordatorio{}
	sessions  = map[string]int{} // sessionID -> userID
	userIDSeq = 1
	remIDSeq  = 1
)

type Recordatorio struct {
	ID          int       `json:"id"`
	Titulo      string    `json:"titulo"`
	Descripcion string    `json:"descripcion"`
	Fecha       time.Time `json:"fecha"`
	Cumplido    bool      `json:"cumplido"`
}

func findUserByUsername(username string) *User {
	for i := range users {
		if users[i].Username == username {
			return &users[i]
		}
	}
	return nil
}

func findUserByID(id int) *User {
	for i := range users {
		if users[i].ID == id {
			return &users[i]
		}
	}
	return nil
}

func findReminderByID(id int) *Recordatorio {
	for i := range reminders {
		if reminders[i].ID == id {
			return &reminders[i]
		}
	}
	return nil
}

// --- Middleware ---
func withAuth(next http.HandlerFunc, requireAdmin bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		uid, ok := sessions[cookie.Value]
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		u := findUserByID(uid)
		if u == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if requireAdmin && u.Role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Attach user to context if needed
		next(w, r)
	}
}

// --- Auth Handlers ---
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // "admin" or "user"
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid fields", http.StatusBadRequest)
		return
	}
	if findUserByUsername(req.Username) != nil {
		http.Error(w, "Username exists", http.StatusBadRequest)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	u := User{
		ID:           userIDSeq,
		Username:     req.Username,
		PasswordHash: string(hash),
		Role:         req.Role,
	}
	userIDSeq++
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	u := findUserByUsername(req.Username)
	if u == nil || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	sid := uuid.New().String()
	sessions[sid] = u.ID
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sid,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure: true, // Uncomment if using HTTPS
	})
	w.WriteHeader(http.StatusOK)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		delete(sessions, cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
		})
	}
	w.WriteHeader(http.StatusOK)
}

// --- Reminder CRUD Handlers ---
func createReminderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var rec Recordatorio
	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	rec.ID = remIDSeq
	remIDSeq++
	rec.Cumplido = false
	reminders = append(reminders, rec)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rec)
}

func listRemindersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reminders)
}

func getReminderHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	id, _ := strconv.Atoi(parts[4])
	rec := findReminderByID(id)
	if rec == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

// --- Update Reminder Handler ---
func updateReminderHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	id, _ := strconv.Atoi(parts[4])
	rec := findReminderByID(id)
	if rec == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	var upd Recordatorio
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	rec.Titulo = upd.Titulo
	rec.Descripcion = upd.Descripcion
	rec.Fecha = upd.Fecha
	rec.Cumplido = upd.Cumplido
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rec)
}

func deleteReminderHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	id, _ := strconv.Atoi(parts[4])
	for i, rec := range reminders {
		if rec.ID == id {
			reminders = append(reminders[:i], reminders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

// --- Main ---
func main() {
	mux := http.NewServeMux()

	// Auth endpoints
	mux.HandleFunc("/api/auth/register", registerHandler)
	mux.HandleFunc("/api/auth/login", loginHandler)
	mux.HandleFunc("/api/auth/logout", logoutHandler)

	// Reminder endpoints (CRUD)
	mux.HandleFunc("/api/v1/reminders", withAuth(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listRemindersHandler(w, r)
		case http.MethodPost:
			createReminderHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}, false))

	mux.HandleFunc("/api/v1/reminders/", func(w http.ResponseWriter, r *http.Request) {
		// Only admin can PUT/DELETE
		switch r.Method {
		case http.MethodGet:
			withAuth(getReminderHandler, false)(w, r)
		case http.MethodPut, http.MethodDelete:
			withAuth(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPut {
					updateReminderHandler(w, r)
				} else if r.Method == http.MethodDelete {
					deleteReminderHandler(w, r)
				}
			}, true)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Static files
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
