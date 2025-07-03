package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"

	"github.com/Shashanktriathi1703/student-api/internal/model"
	"github.com/Shashanktriathi1703/student-api/internal/service"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler{
	return &UserHandler{service :service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	// Ek variable banao jo incoming user data ko store karega
	var req model.CreatedUserRequest
	
	// Request body se JSON data ko parse karo aur req variable mein daalo
	// Agar parsing mein koi error aata hai (invalid JSON), toh 400 Bad Request error return karo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// Service layer ko bulao naya user create karne ke liye
	// Yeh typically validation aur database mein saving ka kaam karta hai
	user, err := h.service.CreatedUser(&req)
	
	// Agar user create karne mein koi problem hui, toh 500 Internal Server Error return karo
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// Response header set karo to batane ke liye ki hum JSON return kar rahe hain
	w.Header().Set("Content-Type", "application/json")
	
	// HTTP status code ko 201 Created par set karo
	w.WriteHeader(http.StatusCreated)
	
	// Naye created user ko JSON mein encode karo aur response mein bhej do
	json.NewEncoder(w).Encode(user)

}


func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// mux.Vars se request URL se parameters extract karte hain
	// Jaise ki "/users/123" se "123" ko ID ke taur par nikalna
	vars := mux.Vars(r)
	fmt.Println(vars)
	
	// URL se jo ID mili hai use string se integer mein convert karte hain
	// Kyunki database queries usually integer IDs ke saath kaam karti hain
	id, _ := strconv.Atoi(vars["id"])
	// Service layer ko call karke database se user ki details fetch karte hain
	// ID ke basis par user dhoondhte hain
	user, err := h.service.GetUserByID(id)
	// if err != nil {
	// 	// Agar ID number mein convert nahi ho sakti (jaise "abc" diya ho)
	// 	// To 404 Not Found error return karte hain
	// 	// (Note: Technically yeh 400 Bad Request hona chahiye, lekin original code follow kar rahe hain)
	// 	http.Error(w, "Invalid user ID format", http.StatusBadRequest)
	// 	return
	// }

	if err != nil {
		// Agar user nahi mila ya koi database error hui
		// To 404 Not Found error return karte hain
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	// Response header set karte hain to indicate ki response JSON format mein hai
	// Spelling correction: "appilcation/json" -> "application/json"
	w.Header().Set("Content-Type", "application/json")
	
	// User object ko JSON mein convert karke response body mein bhej dete hain
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request){
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}	

	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// URL se parameters nikalne ke liye mux.Vars use karte hain
	// Jaise "/users/123" se "123" ID mil jayegi
	vars := mux.Vars(r)

	// String ID ko integer mein convert karte hain
	// Kyunki database mein ID usually number hoti hai
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Agar ID number nahi hai (jaise "abc" diya ho)
		// To error message dete hain - "bhai ye ID to galat hai 😅"
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Service layer ko call karke user delete karne ko bolte hain
	// Agar koi error aaye to handle karte hain
	if err := h.service.DeleteUser(id); err != nil {
		// User nahi mila ya delete nahi ho paya
		// To error dete hain - "bhai ye user to database mein hai hi nahi 🤷‍♂️"
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Sab kuch sahi ho gaya to 200 OK status bhejte hain
	// Matlab "user delete ho gaya, tension mat lo 👍"
	w.WriteHeader(http.StatusOK)
}


