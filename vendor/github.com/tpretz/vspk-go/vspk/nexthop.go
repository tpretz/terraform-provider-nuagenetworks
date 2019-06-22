/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NextHopIdentity represents the Identity of the object
var NextHopIdentity = bambou.Identity{
	Name:     "nexthop",
	Category: "nexthops",
}

// NextHopsList represents a list of NextHops
type NextHopsList []*NextHop

// NextHopsAncestor is the interface that an ancestor of a NextHop must implement.
// An Ancestor is defined as an entity that has NextHop as a descendant.
// An Ancestor can get a list of its child NextHops, but not necessarily create one.
type NextHopsAncestor interface {
	NextHops(*bambou.FetchingInfo) (NextHopsList, *bambou.Error)
}

// NextHopsParent is the interface that a parent of a NextHop must implement.
// A Parent is defined as an entity that has NextHop as a child.
// A Parent is an Ancestor which can create a NextHop.
type NextHopsParent interface {
	NextHopsAncestor
	CreateNextHop(*NextHop) *bambou.Error
}

// NextHop represents the model of a nexthop
type NextHop struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	IPType             string `json:"IPType,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`
	Ip                 string `json:"ip,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
	Type               string `json:"type,omitempty"`
}

// NewNextHop returns a new *NextHop
func NewNextHop() *NextHop {

	return &NextHop{}
}

// Identity returns the Identity of the object.
func (o *NextHop) Identity() bambou.Identity {

	return NextHopIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NextHop) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NextHop) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NextHop from the server
func (o *NextHop) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NextHop into the server
func (o *NextHop) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NextHop from the server
func (o *NextHop) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NextHop
func (o *NextHop) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NextHop
func (o *NextHop) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NextHop
func (o *NextHop) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NextHop
func (o *NextHop) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
