/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VPNConnectionIdentity represents the Identity of the object
var VPNConnectionIdentity = bambou.Identity{
	Name:     "vpnconnection",
	Category: "vpnconnections",
}

// VPNConnectionsList represents a list of VPNConnections
type VPNConnectionsList []*VPNConnection

// VPNConnectionsAncestor is the interface that an ancestor of a VPNConnection must implement.
// An Ancestor is defined as an entity that has VPNConnection as a descendant.
// An Ancestor can get a list of its child VPNConnections, but not necessarily create one.
type VPNConnectionsAncestor interface {
	VPNConnections(*bambou.FetchingInfo) (VPNConnectionsList, *bambou.Error)
}

// VPNConnectionsParent is the interface that a parent of a VPNConnection must implement.
// A Parent is defined as an entity that has VPNConnection as a child.
// A Parent is an Ancestor which can create a VPNConnection.
type VPNConnectionsParent interface {
	VPNConnectionsAncestor
	CreateVPNConnection(*VPNConnection) *bambou.Error
}

// VPNConnection represents the model of a vpnconnection
type VPNConnection struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	Name                   string `json:"name,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	Description            string `json:"description,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	AssociatedWANServiceID string `json:"associatedWANServiceID,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewVPNConnection returns a new *VPNConnection
func NewVPNConnection() *VPNConnection {

	return &VPNConnection{}
}

// Identity returns the Identity of the object.
func (o *VPNConnection) Identity() bambou.Identity {

	return VPNConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VPNConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VPNConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VPNConnection from the server
func (o *VPNConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VPNConnection into the server
func (o *VPNConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VPNConnection from the server
func (o *VPNConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VPNConnection
func (o *VPNConnection) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VPNConnection
func (o *VPNConnection) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VPNConnection
func (o *VPNConnection) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VPNConnection
func (o *VPNConnection) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
