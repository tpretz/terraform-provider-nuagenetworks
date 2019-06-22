/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NetconfSessionIdentity represents the Identity of the object
var NetconfSessionIdentity = bambou.Identity{
	Name:     "netconfsession",
	Category: "netconfsessions",
}

// NetconfSessionsList represents a list of NetconfSessions
type NetconfSessionsList []*NetconfSession

// NetconfSessionsAncestor is the interface that an ancestor of a NetconfSession must implement.
// An Ancestor is defined as an entity that has NetconfSession as a descendant.
// An Ancestor can get a list of its child NetconfSessions, but not necessarily create one.
type NetconfSessionsAncestor interface {
	NetconfSessions(*bambou.FetchingInfo) (NetconfSessionsList, *bambou.Error)
}

// NetconfSessionsParent is the interface that a parent of a NetconfSession must implement.
// A Parent is defined as an entity that has NetconfSession as a child.
// A Parent is an Ancestor which can create a NetconfSession.
type NetconfSessionsParent interface {
	NetconfSessionsAncestor
	CreateNetconfSession(*NetconfSession) *bambou.Error
}

// NetconfSession represents the model of a netconfsession
type NetconfSession struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	LastUpdatedBy         string `json:"lastUpdatedBy,omitempty"`
	GatewayModel          string `json:"gatewayModel,omitempty"`
	GatewayVendor         string `json:"gatewayVendor,omitempty"`
	GatewayVersion        string `json:"gatewayVersion,omitempty"`
	EntityScope           string `json:"entityScope,omitempty"`
	AssociatedGatewayID   string `json:"associatedGatewayID,omitempty"`
	AssociatedGatewayName string `json:"associatedGatewayName,omitempty"`
	Status                string `json:"status,omitempty"`
	ExternalID            string `json:"externalID,omitempty"`
}

// NewNetconfSession returns a new *NetconfSession
func NewNetconfSession() *NetconfSession {

	return &NetconfSession{}
}

// Identity returns the Identity of the object.
func (o *NetconfSession) Identity() bambou.Identity {

	return NetconfSessionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetconfSession) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetconfSession) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetconfSession from the server
func (o *NetconfSession) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetconfSession into the server
func (o *NetconfSession) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetconfSession from the server
func (o *NetconfSession) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NetconfSession
func (o *NetconfSession) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NetconfSession
func (o *NetconfSession) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NetconfSession
func (o *NetconfSession) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NetconfSession
func (o *NetconfSession) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
