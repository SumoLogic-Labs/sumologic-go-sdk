package types

import "github.com/antihax/optional"

type Baselines struct {
	// The average daily amount of continuous logs this child org is expected to ingest, in GB.
	ContinuousIngest int64 `json:"continuousIngest,omitempty"`
	// The average daily amount of frequent logs this child org is expected to ingest, in GB.
	FrequentIngest int64 `json:"frequentIngest,omitempty"`
	// The average daily amount of infrequent logs this child org is expected to ingest, in GB.
	InfrequentIngest int64 `json:"infrequentIngest,omitempty"`
	// The average daily amount of metrics this child org is expected to ingest, in DPMs(Data points per minute).
	Metrics int64 `json:"metrics,omitempty"`
	// The average daily amount of tracing data this child org is expected to ingest, in GB. This is currently not available for all customers. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	TracingIngest int64 `json:"tracingIngest,omitempty"`
}

type BasicUsage struct {
	// The amount of credits used by the organization in form of deployment charge.
	DeploymentChargeCreditsUsed float64 `json:"deploymentChargeCreditsUsed"`
	// The amount of credits used by the organization excluding deployment charge.
	AllocatedCreditsUsed float64 `json:"allocatedCreditsUsed"`
	// The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash `-` character.
	OrgId string `json:"orgId"`
	// Denotes the total number of credits provisioned for the child organization to use.
	TotalCapacity float64 `json:"totalCapacity"`
	// Denotes the total number of credits that have been utilized.
	TotalCreditsUsed float64 `json:"totalCreditsUsed"`
}

type CreditsSubscription struct {
	// Start date of the contract.
	StartDate string `json:"startDate"`
	// End date of the contract.
	EndDate string `json:"endDate"`
	// Status of the subscription.
	Status string `json:"status"`
	Plan   *Plan  `json:"plan"`
	// The total number of credits allocated to the organization.
	Credits int64 `json:"credits"`
	// The number of credits allocated to the organization in form of deployment charge.
	DeploymentChargeCredits int64 `json:"deploymentChargeCredits,omitempty"`
	// The number of credits allocated to the organization excluding deployment charge.
	AllocatedCredits int64      `json:"allocatedCredits,omitempty"`
	Baselines        *Baselines `json:"baselines"`
}

type Deployment struct {
	// Identifier of the deployment.
	DeploymentId string `json:"deploymentId"`
	// The URL to interact with the Sumo Logic service on the corresponding deployment.
	ServiceUrl string `json:"serviceUrl,omitempty"`
	// URL to interact with Sumo Logic APIs on the corresponding deployment.
	ApiUrl string `json:"apiUrl,omitempty"`
}

type DeploymentCharge struct {
	// Identifier of the deployment for the child org for which deployment charge is applied.
	DeploymentId string `json:"deploymentId,omitempty"`
	// Deployment charge is a charge that applies to child orgs deployed in different regions. This number is a percentage applied to the total credits being allocated to the child org.
	DeploymentCharge float64 `json:"deploymentCharge,omitempty"`
}

type DetailedUsage struct {
	// The amount of credits used by the organization in form of deployment charge.
	DeploymentChargeCreditsUsed float64 `json:"deploymentChargeCreditsUsed"`
	// The amount of credits used by the organization excluding deployment charge.
	AllocatedCreditsUsed float64 `json:"allocatedCreditsUsed"`
	// Contains details of the credits used per product variable.
	Usages []UsagePerProductVariable `json:"usages"`
	// Denotes the total number of credits provisioned for the child organization to use.
	TotalCapacity float64 `json:"totalCapacity"`
	// Denotes the total number of credits that have been utilized.
	TotalCreditsUsed float64 `json:"totalCreditsUsed"`
}

type ListOrganizationsOpts struct {
	Limit  optional.Int32
	Token  optional.String
	Status optional.String
}

type ListOrganizationResponse struct {
	// List of organizations with subscription details.
	Data []ReadOrganizationResponse `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ListUsagesResponse struct {
	// Usage details of the requested organizations.
	Data []BasicUsage `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type OrganizationsUsagesOpts struct {
	Limit optional.Int32
	Token optional.String
}

type OrganizationWithSubscriptionDetails struct {
	// Identifier of the deployment in which the organization should be created.
	DeploymentId string `json:"deploymentId"`
	// Specify the duration of the Trial plan. If not specified, your subscription plan will be used for the created organization.
	TrialPlanPeriod int32      `json:"trialPlanPeriod,omitempty"`
	Baselines       *Baselines `json:"baselines,omitempty"`
	// Email address of the account owner.
	Email string `json:"email"`
	// Name of the organization.
	OrganizationName string `json:"organizationName"`
	// First name of the account owner.
	FirstName string `json:"firstName"`
	// Last name of the account owner.
	LastName string `json:"lastName,omitempty"`
}

type ParentOrgInfo struct {
	// Tells whether the parent org can set up trial child orgs subscriptions.
	IsEligibleForTrialOrgs bool `json:"isEligibleForTrialOrgs,omitempty"`
	// Tells whether the org is subject to deployment charges.
	IsEligibleForDeploymentCharge bool `json:"isEligibleForDeploymentCharge,omitempty"`
	// List of deployment charges for the customer for setting up child org in each deployment.
	DeploymentCharges []DeploymentCharge `json:"deploymentCharges,omitempty"`
	// Plan name of the account.
	PlanName string `json:"planName,omitempty"`
}

type ParentUsage struct {
	// Denotes the total number of credits that have been allocated to the child organizations.
	CreditsAllocated float64 `json:"creditsAllocated"`
	// Denotes the total number of credits provisioned for the child organization to use.
	TotalCapacity float64 `json:"totalCapacity"`
	// Denotes the total number of credits that have been utilized.
	TotalCreditsUsed float64 `json:"totalCreditsUsed"`
}

type Plan struct {
	// Name of the subscription plan.
	PlanName string `json:"planName"`
}

type ReadOrganizationResponse struct {
	Subscription *CreditsSubscription `json:"subscription"`
	// The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash `-` character.
	OrgId string `json:"orgId"`
	// Identifier of the deployment in which the organization is present.
	DeploymentId string `json:"deploymentId,omitempty"`
	// Email address of the account owner.
	Email string `json:"email"`
	// Name of the organization.
	OrganizationName string `json:"organizationName"`
	// First name of the account owner.
	FirstName string `json:"firstName"`
	// Last name of the account owner.
	LastName string `json:"lastName,omitempty"`
}

type Subdomain struct {
	// Subdomain login URL of the organization.
	SubdomainLoginUrl string `json:"subdomainLoginUrl"`
}

type UsagePerProductVariable struct {
	// A Product Variable is a unique service performance feature that is tracked through credit utilization. Valid values are 'continuousIngest', 'frequentIngest', 'storage', 'metrics', 'infrequentScan', 'infrequentIngest', 'inFrequentStorage', 'cseIngest', 'cseStorage'.
	ProductVariable string `json:"productVariable"`
	// Denotes the total number of actual credits that have been used.
	CreditsUsed float64 `json:"creditsUsed"`
	// Denotes the total number of credits that have been used in form of deployment charges.
	DeploymentChargeCredits float64 `json:"deploymentChargeCredits"`
	// Denotes the total number of credits that have been used including deployment charges.
	CreditsDeducted float64 `json:"creditsDeducted"`
	// The native utilization of the product variable.
	Utilization float64 `json:"utilization"`
	// The unit in which the native utilization is measured.
	Unit string `json:"unit"`
}
