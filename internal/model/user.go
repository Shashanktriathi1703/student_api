package model

import "time"

// ye jo User ka "U" captail hai, wo isiliye jisse wo export ho raha aur tum kahin import kara sako

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatedUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
