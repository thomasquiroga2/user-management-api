package handlers

import (
	"encoding/json"
	"net/http"
	"user-management-api/models"
	"user-management-api/utils"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Validaciones
	if user.ID == "" || user.Name == "" || user.Email == "" {
		http.Error(w, "Campos obligatorios faltantes", http.StatusBadRequest)
		return
	}

	if _, exists := models.Users[user.ID]; exists {
		http.Error(w, "El ID ya existe", http.StatusConflict)
		return
	}

	if !utils.IsValidEmail(user.Email) {
		http.Error(w, "Correo electrónico no válido", http.StatusBadRequest)
		return
	}

	for _, u := range models.Users {
		if u.Email == user.Email {
			http.Error(w, "El correo electrónico ya está en uso", http.StatusConflict)
			return
		}
	}

	models.Users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario creado con éxito"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, exists := models.Users[id]
	if !exists {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, exists := models.Users[id]
	if !exists {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	var updates models.User
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if updates.Email != "" && !utils.IsValidEmail(updates.Email) {
		http.Error(w, "Correo electrónico no válido", http.StatusBadRequest)
		return
	}

	if updates.Name != "" {
		user.Name = updates.Name
	}
	if updates.Email != "" {
		user.Email = updates.Email
	}
	if updates.Age != 0 {
		user.Age = updates.Age
	}

	models.Users[id] = user
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if _, exists := models.Users[id]; !exists {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	delete(models.Users, id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario eliminado con éxito"})
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]models.User, 0, len(models.Users))
	for _, user := range models.Users {
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
