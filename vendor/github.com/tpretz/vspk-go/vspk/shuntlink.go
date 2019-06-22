/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ShuntLinkIdentity represents the Identity of the object
var ShuntLinkIdentity = bambou.Identity{
	Name:     "shuntlink",
	Category: "shuntlinks",
}

// ShuntLinksList represents a list of ShuntLinks
type ShuntLinksList []*ShuntLink

// ShuntLinksAncestor is the interface that an ancestor of a ShuntLink must implement.
// An Ancestor is defined as an entity that has ShuntLink as a descendant.
// An Ancestor can get a list of its child ShuntLinks, but not necessarily create one.
type ShuntLinksAncestor interface {
	ShuntLinks(*bambou.FetchingInfo) (ShuntLinksList, *bambou.Error)
}

// ShuntLinksParent is the interface that a parent of a ShuntLink must implement.
// A Parent is defined as an entity that has ShuntLink as a child.
// A Parent is an Ancestor which can create a ShuntLink.
type ShuntLinksParent interface {
	ShuntLinksAncestor
	CreateShuntLink(*ShuntLink) *bambou.Error
}

// ShuntLink represents the model of a shuntlink
type ShuntLink struct {
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	VLANPeer1ID    string `json:"VLANPeer1ID,omitempty"`
	VLANPeer2ID    string `json:"VLANPeer2ID,omitempty"`
	Name           string `json:"name,omitempty"`
	LastUpdatedBy  string `json:"lastUpdatedBy,omitempty"`
	GatewayPeer1ID string `json:"gatewayPeer1ID,omitempty"`
	GatewayPeer2ID string `json:"gatewayPeer2ID,omitempty"`
	Peer1IPAddress string `json:"peer1IPAddress,omitempty"`
	Peer1Subnet    string `json:"peer1Subnet,omitempty"`
	Peer2IPAddress string `json:"peer2IPAddress,omitempty"`
	Peer2Subnet    string `json:"peer2Subnet,omitempty"`
	Description    string `json:"description,omitempty"`
	EntityScope    string `json:"entityScope,omitempty"`
	ExternalID     string `json:"externalID,omitempty"`
}

// NewShuntLink returns a new *ShuntLink
func NewShuntLink() *ShuntLink {

	return &ShuntLink{
		Name:           "null",
		GatewayPeer1ID: "null",
		GatewayPeer2ID: "null",
		Description:    "null",
	}
}

// Identity returns the Identity of the object.
func (o *ShuntLink) Identity() bambou.Identity {

	return ShuntLinkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ShuntLink) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ShuntLink) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ShuntLink from the server
func (o *ShuntLink) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ShuntLink into the server
func (o *ShuntLink) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ShuntLink from the server
func (o *ShuntLink) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ShuntLink
func (o *ShuntLink) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ShuntLink
func (o *ShuntLink) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the ShuntLink
func (o *ShuntLink) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// CreateAlarm creates a new child Alarm under the ShuntLink
func (o *ShuntLink) CreateAlarm(child *Alarm) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ShuntLink
func (o *ShuntLink) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ShuntLink
func (o *ShuntLink) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
