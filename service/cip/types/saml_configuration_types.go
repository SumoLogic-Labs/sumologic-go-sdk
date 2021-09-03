package types

import "time"

type AllowlistedUserResult struct {
	// Unique identifier of the user.
	UserId string `json:"userId"`
	// First name of the user.
	FirstName string `json:"firstName"`
	// Last name of the user.
	LastName string `json:"lastName"`
	// Email of the user.
	Email string `json:"email"`
	// If the user can manage SAML Configurations.
	CanManageSaml bool `json:"canManageSaml"`
	// Checks if the user is active.
	IsActive bool `json:"isActive"`
	// Timestamp of the last login of the user.
	LastLogin time.Time `json:"lastLogin"`
}
