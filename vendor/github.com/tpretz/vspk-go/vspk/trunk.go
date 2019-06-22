/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// TrunkIdentity represents the Identity of the object
var TrunkIdentity = bambou.Identity{
	Name:     "trunk",
	Category: "trunks",
}

// TrunksList represents a list of Trunks
type TrunksList []*Trunk

// TrunksAncestor is the interface that an ancestor of a Trunk must implement.
// An Ancestor is defined as an entity that has Trunk as a descendant.
// An Ancestor can get a list of its child Trunks, but not necessarily create one.
type TrunksAncestor interface {
	Trunks(*bambou.FetchingInfo) (TrunksList, *bambou.Error)
}

// TrunksParent is the interface that a parent of a Trunk must implement.
// A Parent is defined as an entity that has Trunk as a child.
// A Parent is an Ancestor which can create a Trunk.
type TrunksParent interface {
	TrunksAncestor
	CreateTrunk(*Trunk) *bambou.Error
}

// Trunk represents the model of a trunk
type Trunk struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Name              string `json:"name,omitempty"`
	LastUpdatedBy     string `json:"lastUpdatedBy,omitempty"`
	EntityScope       string `json:"entityScope,omitempty"`
	AssociatedVPortID string `json:"associatedVPortID,omitempty"`
	ExternalID        string `json:"externalID,omitempty"`
}

// NewTrunk returns a new *Trunk
func NewTrunk() *Trunk {

	return &Trunk{}
}

// Identity returns the Identity of the object.
func (o *Trunk) Identity() bambou.Identity {

	return TrunkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Trunk) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Trunk) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Trunk from the server
func (o *Trunk) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Trunk into the server
func (o *Trunk) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Trunk from the server
func (o *Trunk) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Trunk
func (o *Trunk) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Trunk
func (o *Trunk) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Trunk
func (o *Trunk) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Trunk
func (o *Trunk) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the Trunk
func (o *Trunk) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}
