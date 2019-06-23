/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PermissionIdentity represents the Identity of the object
var PermissionIdentity = bambou.Identity{
	Name:     "permission",
	Category: "permissions",
}

// PermissionsList represents a list of Permissions
type PermissionsList []*Permission

// PermissionsAncestor is the interface that an ancestor of a Permission must implement.
// An Ancestor is defined as an entity that has Permission as a descendant.
// An Ancestor can get a list of its child Permissions, but not necessarily create one.
type PermissionsAncestor interface {
	Permissions(*bambou.FetchingInfo) (PermissionsList, *bambou.Error)
}

// PermissionsParent is the interface that a parent of a Permission must implement.
// A Parent is defined as an entity that has Permission as a child.
// A Parent is an Ancestor which can create a Permission.
type PermissionsParent interface {
	PermissionsAncestor
	CreatePermission(*Permission) *bambou.Error
}

// Permission represents the model of a permission
type Permission struct {
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

// NewPermission returns a new *Permission
func NewPermission() *Permission {

	return &Permission{}
}

// Identity returns the Identity of the object.
func (o *Permission) Identity() bambou.Identity {

	return PermissionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Permission) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Permission) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Permission from the server
func (o *Permission) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Permission into the server
func (o *Permission) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Permission from the server
func (o *Permission) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Permission
func (o *Permission) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Permission
func (o *Permission) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Permission
func (o *Permission) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Permission
func (o *Permission) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the Permission
func (o *Permission) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
