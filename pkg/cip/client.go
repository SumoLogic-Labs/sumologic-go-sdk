package cip

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Sumo Logic API API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AccessKeyManagementApi          *AccessKeyManagementApiService
	AccountManagementApi            *AccountManagementApiService
	AppManagementApi                *AppManagementApiService
	ArchiveManagementApi            *ArchiveManagementApiService
	ConnectionManagementApi         *ConnectionManagementApiService
	ContentManagementApi            *ContentManagementApiService
	ContentPermissionsApi           *ContentPermissionsApiService
	DashboardManagementApi          *DashboardManagementApiService
	DynamicParsingRuleManagementApi *DynamicParsingRuleManagementApiService
	ExtractionRuleManagementApi     *ExtractionRuleManagementApiService
	FieldManagementV1Api            *FieldManagementV1ApiService
	FolderManagementApi             *FolderManagementApiService
	HealthEventsApi                 *HealthEventsApiService
	IngestBudgetManagementV1Api     *IngestBudgetManagementV1ApiService
	IngestBudgetManagementV2Api     *IngestBudgetManagementV2ApiService
	LogSearchesEstimatedUsageApi    *LogSearchesEstimatedUsageApiService
	LookupManagementApi             *LookupManagementApiService
	MetricsSearchesManagementApi    *MetricsSearchesManagementApiService
	MonitorsLibraryManagementApi    *MonitorsLibraryManagementApiService
	PartitionManagementApi          *PartitionManagementApiService
	PasswordPolicyApi               *PasswordPolicyApiService
	RoleManagementApi               *RoleManagementApiService
	SamlConfigurationManagementApi  *SamlConfigurationManagementApiService
	ScheduledViewManagementApi      *ScheduledViewManagementApiService
	ServiceAllowlistManagementApi   *ServiceAllowlistManagementApiService
	TokensLibraryManagementApi      *TokensLibraryManagementApiService
	TransformationRuleManagementApi *TransformationRuleManagementApiService
	UserManagementApi               *UserManagementApiService
}

type service struct {
	client *APIClient
}
