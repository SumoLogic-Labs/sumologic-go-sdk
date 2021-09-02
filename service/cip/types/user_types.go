package types

import (
	"github.com/antihax/optional"
	"time"
)

type CreateUserDefinition struct {
	// First name of the user.
	FirstName string `json:"firstName"`
	// Last name of the user.
	LastName string `json:"lastName"`
	// Email address of the user.
	Email string `json:"email"`
	// List of roleIds associated with the user.
	RoleIds []string `json:"roleIds"`
}

type DisableMfaRequest struct {
	// Email of user whose mfa is being disabled.
	Email string `json:"email"`
	// Password of user whose mfa is being disabled.
	Password string `json:"password"`
}

type ListUserModelsResponse struct {
	// List of users.
	Data []UserModel `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ChangeEmailRequest struct {
	// New email address of the user.
	Email string `json:"email"`
}

type UpdateUserDefinition struct {
	// First name of the user.
	FirstName string `json:"firstName"`
	// Last name of the user.
	LastName string `json:"lastName"`
	// This has the value `true` if the user is active and `false` if they have been deactivated.
	IsActive bool `json:"isActive"`
	// List of role identifiers associated with the user.
	RoleIds []string `json:"roleIds"`
}

type UserInfo struct {
	// User's identifier.
	Id string `json:"id"`
	// User's email.
	Email string `json:"email"`
	// User's first name.
	FirstName string `json:"firstName"`
	// User's last name.
	LastName string `json:"lastName"`
}

type UserModel struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// First name of the user.
	FirstName string `json:"firstName"`
	// Last name of the user.
	LastName string `json:"lastName"`
	// Email address of the user.
	Email string `json:"email"`
	// List of roleIds associated with the user.
	RoleIds []string `json:"roleIds"`
	// Unique identifier for the user.
	Id string `json:"id"`
	// True if the user is active.
	IsActive bool `json:"isActive,omitempty"`
	// This has the value `true` if the user's account has been locked. If a user tries to log into their account several times and fails, his or her account will be locked for security reasons.
	IsLocked bool `json:"isLocked,omitempty"`
	// True if multi factor authentication is enabled for the user.
	IsMfaEnabled bool `json:"isMfaEnabled,omitempty"`
	// Timestamp of the last login for the user in UTC. Will be null if the user has never logged in.
	LastLoginTimestamp time.Time `json:"lastLoginTimestamp,omitempty"`
}

type DeleteUserOpts struct {
	TransferTo optional.String
}

type ListUsersOpts struct {
	Limit  optional.Int32
	Token  optional.String
	SortBy optional.String
	Email  optional.String
}
