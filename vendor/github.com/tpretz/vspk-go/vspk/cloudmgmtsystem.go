/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CloudMgmtSystemIdentity represents the Identity of the object
var CloudMgmtSystemIdentity = bambou.Identity{
	Name:     "cms",
	Category: "cms",
}

// CloudMgmtSystemsList represents a list of CloudMgmtSystems
type CloudMgmtSystemsList []*CloudMgmtSystem

// CloudMgmtSystemsAncestor is the interface that an ancestor of a CloudMgmtSystem must implement.
// An Ancestor is defined as an entity that has CloudMgmtSystem as a descendant.
// An Ancestor can get a list of its child CloudMgmtSystems, but not necessarily create one.
type CloudMgmtSystemsAncestor interface {
	CloudMgmtSystems(*bambou.FetchingInfo) (CloudMgmtSystemsList, *bambou.Error)
}

// CloudMgmtSystemsParent is the interface that a parent of a CloudMgmtSystem must implement.
// A Parent is defined as an entity that has CloudMgmtSystem as a child.
// A Parent is an Ancestor which can create a CloudMgmtSystem.
type CloudMgmtSystemsParent interface {
	CloudMgmtSystemsAncestor
	CreateCloudMgmtSystem(*CloudMgmtSystem) *bambou.Error
}

// CloudMgmtSystem represents the model of a cms
type CloudMgmtSystem struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewCloudMgmtSystem returns a new *CloudMgmtSystem
func NewCloudMgmtSystem() *CloudMgmtSystem {

	return &CloudMgmtSystem{}
}

// Identity returns the Identity of the object.
func (o *CloudMgmtSystem) Identity() bambou.Identity {

	return CloudMgmtSystemIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudMgmtSystem) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudMgmtSystem) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CloudMgmtSystem from the server
func (o *CloudMgmtSystem) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CloudMgmtSystem into the server
func (o *CloudMgmtSystem) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CloudMgmtSystem from the server
func (o *CloudMgmtSystem) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the CloudMgmtSystem
func (o *CloudMgmtSystem) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the CloudMgmtSystem
func (o *CloudMgmtSystem) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the CloudMgmtSystem
func (o *CloudMgmtSystem) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the CloudMgmtSystem
func (o *CloudMgmtSystem) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
