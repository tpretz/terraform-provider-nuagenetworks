/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// NetworkLayoutIdentity represents the Identity of the object
var NetworkLayoutIdentity = bambou.Identity{
	Name:     "networklayout",
	Category: "networklayout",
}

// NetworkLayoutsList represents a list of NetworkLayouts
type NetworkLayoutsList []*NetworkLayout

// NetworkLayoutsAncestor is the interface that an ancestor of a NetworkLayout must implement.
// An Ancestor is defined as an entity that has NetworkLayout as a descendant.
// An Ancestor can get a list of its child NetworkLayouts, but not necessarily create one.
type NetworkLayoutsAncestor interface {
	NetworkLayouts(*bambou.FetchingInfo) (NetworkLayoutsList, *bambou.Error)
}

// NetworkLayoutsParent is the interface that a parent of a NetworkLayout must implement.
// A Parent is defined as an entity that has NetworkLayout as a child.
// A Parent is an Ancestor which can create a NetworkLayout.
type NetworkLayoutsParent interface {
	NetworkLayoutsAncestor
	CreateNetworkLayout(*NetworkLayout) *bambou.Error
}

// NetworkLayout represents the model of a networklayout
type NetworkLayout struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	ServiceType         string `json:"serviceType,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	RouteReflectorIP    string `json:"routeReflectorIP,omitempty"`
	AutonomousSystemNum int    `json:"autonomousSystemNum,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

// NewNetworkLayout returns a new *NetworkLayout
func NewNetworkLayout() *NetworkLayout {

	return &NetworkLayout{}
}

// Identity returns the Identity of the object.
func (o *NetworkLayout) Identity() bambou.Identity {

	return NetworkLayoutIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkLayout) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkLayout) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetworkLayout from the server
func (o *NetworkLayout) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetworkLayout into the server
func (o *NetworkLayout) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetworkLayout from the server
func (o *NetworkLayout) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NetworkLayout
func (o *NetworkLayout) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NetworkLayout
func (o *NetworkLayout) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NetworkLayout
func (o *NetworkLayout) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NetworkLayout
func (o *NetworkLayout) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
