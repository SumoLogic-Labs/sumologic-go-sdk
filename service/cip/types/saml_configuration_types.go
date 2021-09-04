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

type OnDemandProvisioningInfo struct {
	// First name attribute of the new user account.
	FirstNameAttribute string `json:"firstNameAttribute,omitempty"`
	// Last name attribute of the new user account.
	LastNameAttribute string `json:"lastNameAttribute,omitempty"`
	// Sumo Logic RBAC roles to be assigned when user accounts are provisioned.
	OnDemandProvisioningRoles []string `json:"onDemandProvisioningRoles"`
}

type SamlIdentityProvider struct {
	// Authentication Request Signing Certificate for the user.
	Certificate string `json:"certificate"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// The identifier used to generate a unique URL for user login.
	SpInitiatedLoginPath string `json:"spInitiatedLoginPath,omitempty"`
	// Name of the SSO policy or another name used to describe the policy internally.
	ConfigurationName string `json:"configurationName"`
	// The unique URL assigned to the organization by the SAML Identity Provider.
	Issuer string `json:"issuer"`
	// True if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in.
	SpInitiatedLoginEnabled bool `json:"spInitiatedLoginEnabled,omitempty"`
	// The URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider.
	AuthnRequestUrl string `json:"authnRequestUrl,omitempty"`
	// The certificate is used to verify the signature in SAML assertions.
	X509cert1 string `json:"x509cert1"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires.
	X509cert2 string `json:"x509cert2,omitempty"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty.
	X509cert3                   string                    `json:"x509cert3,omitempty"`
	OnDemandProvisioningEnabled *OnDemandProvisioningInfo `json:"onDemandProvisioningEnabled,omitempty"`
	// The role that Sumo Logic will assign to users when they sign in.
	RolesAttribute string `json:"rolesAttribute,omitempty"`
	// True if users are redirected to a URL after signing out of Sumo Logic.
	LogoutEnabled bool `json:"logoutEnabled,omitempty"`
	// The URL that users will be redirected to after signing out of Sumo Logic.
	LogoutUrl string `json:"logoutUrl,omitempty"`
	// The email address of the new user account.
	EmailAttribute string `json:"emailAttribute,omitempty"`
	// True if additional details are included when a user fails to sign in.
	DebugMode bool `json:"debugMode,omitempty"`
	// True if Sumo Logic will send signed Authn requests to the identity provider.
	SignAuthnRequest bool `json:"signAuthnRequest,omitempty"`
	// True if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider.
	DisableRequestedAuthnContext bool `json:"disableRequestedAuthnContext,omitempty"`
	// True if the SAML binding is of HTTP Redirect type.
	IsRedirectBinding bool `json:"isRedirectBinding,omitempty"`
	// Unique identifier of the SAML Identity Provider.
	Id string `json:"id"`
	// The URL on Sumo Logic where the IdP will redirect to with its authentication response.
	AssertionConsumerUrl string `json:"assertionConsumerUrl,omitempty"`
}

type SamlIdentityProviderRequest struct {
	// The identifier used to generate a unique URL for user login.
	SpInitiatedLoginPath string `json:"spInitiatedLoginPath,omitempty"`
	// Name of the SSO policy or another name used to describe the policy internally.
	ConfigurationName string `json:"configurationName"`
	// The unique URL assigned to the organization by the SAML Identity Provider.
	Issuer string `json:"issuer"`
	// True if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in.
	SpInitiatedLoginEnabled bool `json:"spInitiatedLoginEnabled,omitempty"`
	// The URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider.
	AuthnRequestUrl string `json:"authnRequestUrl,omitempty"`
	// The certificate is used to verify the signature in SAML assertions.
	X509cert1 string `json:"x509cert1"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires.
	X509cert2 string `json:"x509cert2,omitempty"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty.
	X509cert3                   string                    `json:"x509cert3,omitempty"`
	OnDemandProvisioningEnabled *OnDemandProvisioningInfo `json:"onDemandProvisioningEnabled,omitempty"`
	// The role that Sumo Logic will assign to users when they sign in.
	RolesAttribute string `json:"rolesAttribute,omitempty"`
	// True if users are redirected to a URL after signing out of Sumo Logic.
	LogoutEnabled bool `json:"logoutEnabled,omitempty"`
	// The URL that users will be redirected to after signing out of Sumo Logic.
	LogoutUrl string `json:"logoutUrl,omitempty"`
	// The email address of the new user account.
	EmailAttribute string `json:"emailAttribute,omitempty"`
	// True if additional details are included when a user fails to sign in.
	DebugMode bool `json:"debugMode,omitempty"`
	// True if Sumo Logic will send signed Authn requests to the identity provider.
	SignAuthnRequest bool `json:"signAuthnRequest,omitempty"`
	// True if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider.
	DisableRequestedAuthnContext bool `json:"disableRequestedAuthnContext,omitempty"`
	// True if the SAML binding is of HTTP Redirect type.
	IsRedirectBinding bool `json:"isRedirectBinding,omitempty"`
}
