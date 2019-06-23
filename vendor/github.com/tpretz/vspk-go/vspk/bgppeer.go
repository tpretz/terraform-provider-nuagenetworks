/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// BGPPeerIdentity represents the Identity of the object
var BGPPeerIdentity = bambou.Identity{
	Name:     "bgppeer",
	Category: "bgppeers",
}

// BGPPeersList represents a list of BGPPeers
type BGPPeersList []*BGPPeer

// BGPPeersAncestor is the interface that an ancestor of a BGPPeer must implement.
// An Ancestor is defined as an entity that has BGPPeer as a descendant.
// An Ancestor can get a list of its child BGPPeers, but not necessarily create one.
type BGPPeersAncestor interface {
	BGPPeers(*bambou.FetchingInfo) (BGPPeersList, *bambou.Error)
}

// BGPPeersParent is the interface that a parent of a BGPPeer must implement.
// A Parent is defined as an entity that has BGPPeer as a child.
// A Parent is an Ancestor which can create a BGPPeer.
type BGPPeersParent interface {
	BGPPeersAncestor
	CreateBGPPeer(*BGPPeer) *bambou.Error
}

// BGPPeer represents the model of a bgppeer
type BGPPeer struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	LastStateChange int    `json:"lastStateChange,omitempty"`
	Address         string `json:"address,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	Status          string `json:"status,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewBGPPeer returns a new *BGPPeer
func NewBGPPeer() *BGPPeer {

	return &BGPPeer{}
}

// Identity returns the Identity of the object.
func (o *BGPPeer) Identity() bambou.Identity {

	return BGPPeerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BGPPeer) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BGPPeer) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BGPPeer from the server
func (o *BGPPeer) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BGPPeer into the server
func (o *BGPPeer) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BGPPeer from the server
func (o *BGPPeer) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BGPPeer
func (o *BGPPeer) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BGPPeer
func (o *BGPPeer) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BGPPeer
func (o *BGPPeer) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BGPPeer
func (o *BGPPeer) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
