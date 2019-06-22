/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EnterprisePermissionIdentity represents the Identity of the object
var EnterprisePermissionIdentity = bambou.Identity{
	Name:     "enterprisepermission",
	Category: "enterprisepermissions",
}

// EnterprisePermissionsList represents a list of EnterprisePermissions
type EnterprisePermissionsList []*EnterprisePermission

// EnterprisePermissionsAncestor is the interface that an ancestor of a EnterprisePermission must implement.
// An Ancestor is defined as an entity that has EnterprisePermission as a descendant.
// An Ancestor can get a list of its child EnterprisePermissions, but not necessarily create one.
type EnterprisePermissionsAncestor interface {
	EnterprisePermissions(*bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error)
}

// EnterprisePermissionsParent is the interface that a parent of a EnterprisePermission must implement.
// A Parent is defined as an entity that has EnterprisePermission as a child.
// A Parent is an Ancestor which can create a EnterprisePermission.
type EnterprisePermissionsParent interface {
	EnterprisePermissionsAncestor
	CreateEnterprisePermission(*EnterprisePermission) *bambou.Error
}

// EnterprisePermission represents the model of a enterprisepermission
type EnterprisePermission struct {
	ID                         string `json:"ID,omitempty"`
	ParentID                   string `json:"parentID,omitempty"`
	ParentType                 string `json:"parentType,omitempty"`
	Owner                      string `json:"owner,omitempty"`
	Name                       string `json:"name,omitempty"`
	LastUpdatedBy              string `json:"lastUpdatedBy,omitempty"`
	PermittedAction            string `json:"permittedAction,omitempty"`
	PermittedEntityDescription string `json:"permittedEntityDescription,omitempty"`
	PermittedEntityID          string `json:"permittedEntityID,omitempty"`
	PermittedEntityName        string `json:"permittedEntityName,omitempty"`
	PermittedEntityType        string `json:"permittedEntityType,omitempty"`
	EntityScope                string `json:"entityScope,omitempty"`
	ExternalID                 string `json:"externalID,omitempty"`
}

// NewEnterprisePermission returns a new *EnterprisePermission
func NewEnterprisePermission() *EnterprisePermission {

	return &EnterprisePermission{}
}

// Identity returns the Identity of the object.
func (o *EnterprisePermission) Identity() bambou.Identity {

	return EnterprisePermissionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnterprisePermission) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnterprisePermission) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EnterprisePermission from the server
func (o *EnterprisePermission) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EnterprisePermission into the server
func (o *EnterprisePermission) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EnterprisePermission from the server
func (o *EnterprisePermission) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EnterprisePermission
func (o *EnterprisePermission) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EnterprisePermission
func (o *EnterprisePermission) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EnterprisePermission
func (o *EnterprisePermission) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EnterprisePermission
func (o *EnterprisePermission) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
