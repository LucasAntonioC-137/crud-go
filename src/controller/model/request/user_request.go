package request

import "time"

// UserRequest represents the input data for creating a new user.
// @Summary User Input Data
// @Description Structure containing the required fields for creating a new user.
type UserRequest struct {
	// User's email (required and must be a valid email address).
	// Example: johndoe@example.com
	Email string `json:"email" binding:"required,email" example:"johndoe@example.com"`

	// User's password (required, minimum of 6 characters,
	// and must contain at least one of the special characters: !@#$%*).
	// Example: Secret@123
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"Secret@123"`

	// User's full name (required, minimum of 4 characters, maximum of 100 characters).
	// Example: John Doe
	Name string `json:"name" binding:"required,min=4,max=100" example:"John Doe"`

	// User's age (required, must be between 2 and 140).
	// Example: 25
	Age int8 `json:"age" binding:"required,min=2,max=140" example:"25"`

	PasswordExpiration time.Time `bson:"password_expiration"`
}

type UserRequestCreate struct {
	// User's email (required and must be a valid email address).
	// Example: johndoe@example.com
	Email string `json:"email" binding:"required,email" example:"johndoe@example.com"`

	// User's password (required, minimum of 6 characters,
	// and must contain at least one of the special characters: !@#$%*).
	// Example: Secret@123
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"Secret@123"`

	// User's full name (required, minimum of 4 characters, maximum of 100 characters).
	// Example: John Doe
	Name string `json:"name" binding:"required,min=4,max=100" example:"John Doe"`

	// User's age (required, must be between 2 and 140).
	// Example: 25
	Age int8 `json:"age" binding:"required,min=2,max=140" example:"25"`
}

// UserUpdateRequest represents the input data for updating a user.
// @Summary User Update Data
// @Description Structure containing the optional fields for updating an existing user.
type UserUpdateRequest struct {
	// User's new full name (optional, minimum of 4 characters, maximum of 100 characters).
	// Example: Johnny Updated
	Name string `json:"name" binding:"omitempty,min=4,max=100" example:"Johnny Updated"`

	// User's new age (optional, must be between 2 and 140).
	// Example: 30
	Age int8 `json:"age" binding:"omitempty,min=2,max=140" example:"30"`
}
