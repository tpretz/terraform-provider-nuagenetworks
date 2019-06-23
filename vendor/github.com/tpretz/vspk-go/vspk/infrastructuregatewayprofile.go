/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// InfrastructureGatewayProfileIdentity represents the Identity of the object
var InfrastructureGatewayProfileIdentity = bambou.Identity{
	Name:     "infrastructuregatewayprofile",
	Category: "infrastructuregatewayprofiles",
}

// InfrastructureGatewayProfilesList represents a list of InfrastructureGatewayProfiles
type InfrastructureGatewayProfilesList []*InfrastructureGatewayProfile

// InfrastructureGatewayProfilesAncestor is the interface that an ancestor of a InfrastructureGatewayProfile must implement.
// An Ancestor is defined as an entity that has InfrastructureGatewayProfile as a descendant.
// An Ancestor can get a list of its child InfrastructureGatewayProfiles, but not necessarily create one.
type InfrastructureGatewayProfilesAncestor interface {
	InfrastructureGatewayProfiles(*bambou.FetchingInfo) (InfrastructureGatewayProfilesList, *bambou.Error)
}

// InfrastructureGatewayProfilesParent is the interface that a parent of a InfrastructureGatewayProfile must implement.
// A Parent is defined as an entity that has InfrastructureGatewayProfile as a child.
// A Parent is an Ancestor which can create a InfrastructureGatewayProfile.
type InfrastructureGatewayProfilesParent interface {
	InfrastructureGatewayProfilesAncestor
	CreateInfrastructureGatewayProfile(*InfrastructureGatewayProfile) *bambou.Error
}

// InfrastructureGatewayProfile represents the model of a infrastructuregatewayprofile
type InfrastructureGatewayProfile struct {
	ID                           string `json:"ID,omitempty"`
	ParentID                     string `json:"parentID,omitempty"`
	ParentType                   string `json:"parentType,omitempty"`
	Owner                        string `json:"owner,omitempty"`
	NTPServerKey                 string `json:"NTPServerKey,omitempty"`
	NTPServerKeyID               int    `json:"NTPServerKeyID,omitempty"`
	Name                         string `json:"name,omitempty"`
	LastUpdatedBy                string `json:"lastUpdatedBy,omitempty"`
	DatapathSyncTimeout          int    `json:"datapathSyncTimeout,omitempty"`
	DeadTimer                    string `json:"deadTimer,omitempty"`
	DeadTimerEnabled             bool   `json:"deadTimerEnabled"`
	RemoteLogMode                string `json:"remoteLogMode,omitempty"`
	RemoteLogServerAddress       string `json:"remoteLogServerAddress,omitempty"`
	RemoteLogServerPort          int    `json:"remoteLogServerPort,omitempty"`
	Description                  string `json:"description,omitempty"`
	MetadataUpgradePath          string `json:"metadataUpgradePath,omitempty"`
	FlowEvictionThreshold        int    `json:"flowEvictionThreshold,omitempty"`
	EnterpriseID                 string `json:"enterpriseID,omitempty"`
	EntityScope                  string `json:"entityScope,omitempty"`
	ControllerLessDuration       string `json:"controllerLessDuration,omitempty"`
	ControllerLessForwardingMode string `json:"controllerLessForwardingMode,omitempty"`
	ControllerLessRemoteDuration string `json:"controllerLessRemoteDuration,omitempty"`
	ForceImmediateSystemSync     bool   `json:"forceImmediateSystemSync"`
	OpenFlowAuditTimer           int    `json:"openFlowAuditTimer,omitempty"`
	UpgradeAction                string `json:"upgradeAction,omitempty"`
	ProxyDNSName                 string `json:"proxyDNSName,omitempty"`
	UseTwoFactor                 bool   `json:"useTwoFactor"`
	StatsCollectorPort           int    `json:"statsCollectorPort,omitempty"`
	ExternalID                   string `json:"externalID,omitempty"`
	SystemSyncScheduler          string `json:"systemSyncScheduler,omitempty"`
}

// NewInfrastructureGatewayProfile returns a new *InfrastructureGatewayProfile
func NewInfrastructureGatewayProfile() *InfrastructureGatewayProfile {

	return &InfrastructureGatewayProfile{
		DatapathSyncTimeout:          1000,
		DeadTimerEnabled:             false,
		RemoteLogMode:                "DISABLED",
		RemoteLogServerPort:          514,
		FlowEvictionThreshold:        2500,
		ControllerLessDuration:       "P7DT0H0M",
		ControllerLessForwardingMode: "DISABLED",
		ControllerLessRemoteDuration: "P3DT0H0M",
		ForceImmediateSystemSync:     false,
		OpenFlowAuditTimer:           180,
		UpgradeAction:                "DOWNLOAD_AND_UPGRADE_AT_WINDOW",
		UseTwoFactor:                 true,
		StatsCollectorPort:           39090,
		SystemSyncScheduler:          "0 0 * * *",
	}
}

// Identity returns the Identity of the object.
func (o *InfrastructureGatewayProfile) Identity() bambou.Identity {

	return InfrastructureGatewayProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructureGatewayProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructureGatewayProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the InfrastructureGatewayProfile from the server
func (o *InfrastructureGatewayProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the InfrastructureGatewayProfile into the server
func (o *InfrastructureGatewayProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the InfrastructureGatewayProfile from the server
func (o *InfrastructureGatewayProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the InfrastructureGatewayProfile
func (o *InfrastructureGatewayProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the InfrastructureGatewayProfile
func (o *InfrastructureGatewayProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the InfrastructureGatewayProfile
func (o *InfrastructureGatewayProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the InfrastructureGatewayProfile
func (o *InfrastructureGatewayProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
