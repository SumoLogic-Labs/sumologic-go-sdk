package types

type App struct {
	AppDefinition *AppDefinition `json:"appDefinition"`
	AppManifest   *AppManifest   `json:"appManifest"`
}

type AppDefinition struct {
	// Content identifier of the app in hexadecimal format.
	ContentId string `json:"contentId"`
	// Unique identifier for the app.
	Uuid string `json:"uuid"`
	// Name of the app.
	Name string `json:"name"`
	// Version of the app.
	AppVersion string `json:"appVersion"`
	// Indicates whether the app is in preview or not.
	Preview bool `json:"preview,omitempty"`
	// Manifest version of the app
	ManifestVersion string `json:"manifestVersion,omitempty"`
}

type AppInstallRequest struct {
	// Preferred name of the app to be installed. This will be the name of the app in the selected installation folder.
	Name string `json:"name"`
	// Preferred description of the app to be installed. This will be displayed as the app description in the selected installation folder.
	Description string `json:"description"`
	// Identifier of the folder in which the app will be installed in hexadecimal format.
	DestinationFolderId string `json:"destinationFolderId"`
	// Dictionary of properties specifying log-source name and value.
	DataSourceValues map[string]string `json:"dataSourceValues,omitempty"`
}

type AppManifest struct {
	// The app family
	Family string `json:"family,omitempty"`
	// Description of the app.
	Description string `json:"description"`
	// Categories that the app belongs to.
	Categories []string `json:"categories,omitempty"`
	// Text to be displayed when hovered over in UI.
	HoverText string `json:"hoverText"`
	// App icon URL.
	IconURL string `json:"iconURL"`
	// App screenshot URLs.
	ScreenshotURLs []string `json:"screenshotURLs,omitempty"`
	// App help page URL.
	HelpURL string `json:"helpURL,omitempty"`
	// the IDs of the docs pages for this app
	HelpDocIdMap map[string]string `json:"helpDocIdMap,omitempty"`
	// App community page URL.
	CommunityURL string `json:"communityURL,omitempty"`
	// Requirements for the app.
	Requirements []string `json:"requirements,omitempty"`
	// Account types that are allowed to install the app
	AccountTypes []string `json:"accountTypes,omitempty"`
	// Indicates whether installation instructions are required or not.
	RequiresInstallationInstructions bool `json:"requiresInstallationInstructions,omitempty"`
	// Installation instructions for the app.
	InstallationInstructions string `json:"installationInstructions,omitempty"`
	// Content identifier of the app.
	Parameters []ServiceManifestDataSourceParameter `json:"parameters,omitempty"`
	// App author.
	Author string `json:"author,omitempty"`
	// App author website URL.
	AuthorWebsite string `json:"authorWebsite,omitempty"`
}

type ListAppsResult struct {
	// An array of Apps
	Apps []App `json:"apps"`
}

type ServiceManifestDataSourceParameter struct {
	// Parameter type.
	ParameterType string `json:"parameterType"`
	// Parameter identifier.
	ParameterId string `json:"parameterId"`
	// Data source type.
	DataSourceType string `json:"dataSourceType,omitempty"`
	// Label.
	Label string `json:"label,omitempty"`
	// Description.
	Description string `json:"description,omitempty"`
	// Example.
	Example string `json:"example,omitempty"`
	// Should the UI display?
	Hidden bool `json:"hidden,omitempty"`
}
