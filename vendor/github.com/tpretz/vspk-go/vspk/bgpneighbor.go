/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// BGPNeighborIdentity represents the Identity of the object
var BGPNeighborIdentity = bambou.Identity{
	Name:     "bgpneighbor",
	Category: "bgpneighbors",
}

// BGPNeighborsList represents a list of BGPNeighbors
type BGPNeighborsList []*BGPNeighbor

// BGPNeighborsAncestor is the interface that an ancestor of a BGPNeighbor must implement.
// An Ancestor is defined as an entity that has BGPNeighbor as a descendant.
// An Ancestor can get a list of its child BGPNeighbors, but not necessarily create one.
type BGPNeighborsAncestor interface {
	BGPNeighbors(*bambou.FetchingInfo) (BGPNeighborsList, *bambou.Error)
}

// BGPNeighborsParent is the interface that a parent of a BGPNeighbor must implement.
// A Parent is defined as an entity that has BGPNeighbor as a child.
// A Parent is an Ancestor which can create a BGPNeighbor.
type BGPNeighborsParent interface {
	BGPNeighborsAncestor
	CreateBGPNeighbor(*BGPNeighbor) *bambou.Error
}

// BGPNeighbor represents the model of a bgpneighbor
type BGPNeighbor struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	BFDEnabled                      bool   `json:"BFDEnabled"`
	IPType                          string `json:"IPType,omitempty"`
	IPv6Address                     string `json:"IPv6Address,omitempty"`
	Name                            string `json:"name,omitempty"`
	DampeningEnabled                bool   `json:"dampeningEnabled"`
	PeerAS                          int    `json:"peerAS,omitempty"`
	PeerIP                          string `json:"peerIP,omitempty"`
	Description                     string `json:"description,omitempty"`
	Session                         string `json:"session,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	AssociatedExportRoutingPolicyID string `json:"associatedExportRoutingPolicyID,omitempty"`
	AssociatedImportRoutingPolicyID string `json:"associatedImportRoutingPolicyID,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
}

// NewBGPNeighbor returns a new *BGPNeighbor
func NewBGPNeighbor() *BGPNeighbor {

	return &BGPNeighbor{
		BFDEnabled: false,
		IPType:     "IPV4",
	}
}

// Identity returns the Identity of the object.
func (o *BGPNeighbor) Identity() bambou.Identity {

	return BGPNeighborIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BGPNeighbor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BGPNeighbor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BGPNeighbor from the server
func (o *BGPNeighbor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BGPNeighbor into the server
func (o *BGPNeighbor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BGPNeighbor from the server
func (o *BGPNeighbor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BGPNeighbor
func (o *BGPNeighbor) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BGPNeighbor
func (o *BGPNeighbor) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BGPNeighbor
func (o *BGPNeighbor) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BGPNeighbor
func (o *BGPNeighbor) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
