package types

type PasswordPolicy struct {
	// The minimum length of the password.
	MinLength int32 `json:"minLength,omitempty"`
	// The maximum length of the password.
	MaxLength int32 `json:"maxLength,omitempty"`
	// If the password must contain lower case characters.
	MustContainLowercase bool `json:"mustContainLowercase,omitempty"`
	// If the password must contain upper case characters.
	MustContainUppercase bool `json:"mustContainUppercase,omitempty"`
	// If the password must contain digits.
	MustContainDigits bool `json:"mustContainDigits,omitempty"`
	// If the password must contain special characters.
	MustContainSpecialChars bool `json:"mustContainSpecialChars,omitempty"`
	// Maximum number of days that a password can be used before user is required to change it. Put -1 if the user should not have to change their password.
	MaxPasswordAgeInDays int32 `json:"maxPasswordAgeInDays,omitempty"`
	// The minimum number of unique new passwords that a user must use before an old password can be reused.
	MinUniquePasswords int32 `json:"minUniquePasswords,omitempty"`
	// Number of failed login attempts allowed before account is locked-out.
	AccountLockoutThreshold int32 `json:"accountLockoutThreshold,omitempty"`
	// The duration of time in minutes that must elapse from the first failed login attempt after which failed login count is reset to 0.
	FailedLoginResetDurationInMins int32 `json:"failedLoginResetDurationInMins,omitempty"`
	// The duration of time in minutes that a locked-out account remained locked before getting unlocked automatically.
	AccountLockoutDurationInMins int32 `json:"accountLockoutDurationInMins,omitempty"`
	// If MFA should be required to log in. By default, this field is set to `false`.
	RequireMfa bool `json:"requireMfa,omitempty"`
	// If MFA should be remembered on the browser.
	RememberMfa bool `json:"rememberMfa,omitempty"`
}
