package types

import "time"

type AccountStatusResponse struct {
	// Whether the account is `cloudflex` or `credits`
	PricingModel string `json:"pricingModel"`
	// If the plan can be updated by the given user
	CanUpdatePlan bool `json:"canUpdatePlan"`
	// Whether the account is `Free`/`Trial`/`Paid`
	PlanType string `json:"planType"`
	// The number of days in which the plan will expire
	PlanExpirationDays int32 `json:"planExpirationDays,omitempty"`
	// The current usage of the application.
	ApplicationUse string `json:"applicationUse"`
}

type ConfigureSubdomainRequest struct {
	// The new subdomain.
	Subdomain string `json:"subdomain"`
}

type SubdomainDefinitionResponse struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// The new subdomain.
	Subdomain string `json:"subdomain"`
	// Login URL corresponding to the subdomain.
	Url string `json:"url"`
}
