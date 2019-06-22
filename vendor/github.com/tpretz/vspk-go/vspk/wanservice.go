/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// WANServiceIdentity represents the Identity of the object
var WANServiceIdentity = bambou.Identity{
	Name:     "service",
	Category: "services",
}

// WANServicesList represents a list of WANServices
type WANServicesList []*WANService

// WANServicesAncestor is the interface that an ancestor of a WANService must implement.
// An Ancestor is defined as an entity that has WANService as a descendant.
// An Ancestor can get a list of its child WANServices, but not necessarily create one.
type WANServicesAncestor interface {
	WANServices(*bambou.FetchingInfo) (WANServicesList, *bambou.Error)
}

// WANServicesParent is the interface that a parent of a WANService must implement.
// A Parent is defined as an entity that has WANService as a child.
// A Parent is an Ancestor which can create a WANService.
type WANServicesParent interface {
	WANServicesAncestor
	CreateWANService(*WANService) *bambou.Error
}

// WANService represents the model of a service
type WANService struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	WANServiceIdentifier   string `json:"WANServiceIdentifier,omitempty"`
	IRBEnabled             bool   `json:"IRBEnabled"`
	Name                   string `json:"name,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	PermittedAction        string `json:"permittedAction,omitempty"`
	ServicePolicy          string `json:"servicePolicy,omitempty"`
	ServiceType            string `json:"serviceType,omitempty"`
	Description            string `json:"description,omitempty"`
	VnId                   int    `json:"vnId,omitempty"`
	EnterpriseName         string `json:"enterpriseName,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	DomainName             string `json:"domainName,omitempty"`
	ConfigType             string `json:"configType,omitempty"`
	Orphan                 bool   `json:"orphan"`
	UseUserMnemonic        bool   `json:"useUserMnemonic"`
	UserMnemonic           string `json:"userMnemonic,omitempty"`
	AssociatedDomainID     string `json:"associatedDomainID,omitempty"`
	AssociatedVPNConnectID string `json:"associatedVPNConnectID,omitempty"`
	TunnelType             string `json:"tunnelType,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
	ExternalRouteTarget    string `json:"externalRouteTarget,omitempty"`
}

// NewWANService returns a new *WANService
func NewWANService() *WANService {

	return &WANService{}
}

// Identity returns the Identity of the object.
func (o *WANService) Identity() bambou.Identity {

	return WANServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *WANService) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *WANService) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the WANService from the server
func (o *WANService) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the WANService into the server
func (o *WANService) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the WANService from the server
func (o *WANService) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the WANService
func (o *WANService) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the WANService
func (o *WANService) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the WANService
func (o *WANService) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the WANService
func (o *WANService) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the WANService
func (o *WANService) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the WANService
func (o *WANService) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the WANService
func (o *WANService) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the WANService
func (o *WANService) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the WANService
func (o *WANService) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the WANService
func (o *WANService) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
