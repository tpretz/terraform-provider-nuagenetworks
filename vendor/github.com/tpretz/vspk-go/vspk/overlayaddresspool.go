/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// OverlayAddressPoolIdentity represents the Identity of the object
var OverlayAddressPoolIdentity = bambou.Identity{
	Name:     "overlayaddresspool",
	Category: "overlayaddresspools",
}

// OverlayAddressPoolsList represents a list of OverlayAddressPools
type OverlayAddressPoolsList []*OverlayAddressPool

// OverlayAddressPoolsAncestor is the interface that an ancestor of a OverlayAddressPool must implement.
// An Ancestor is defined as an entity that has OverlayAddressPool as a descendant.
// An Ancestor can get a list of its child OverlayAddressPools, but not necessarily create one.
type OverlayAddressPoolsAncestor interface {
	OverlayAddressPools(*bambou.FetchingInfo) (OverlayAddressPoolsList, *bambou.Error)
}

// OverlayAddressPoolsParent is the interface that a parent of a OverlayAddressPool must implement.
// A Parent is defined as an entity that has OverlayAddressPool as a child.
// A Parent is an Ancestor which can create a OverlayAddressPool.
type OverlayAddressPoolsParent interface {
	OverlayAddressPoolsAncestor
	CreateOverlayAddressPool(*OverlayAddressPool) *bambou.Error
}

// OverlayAddressPool represents the model of a overlayaddresspool
type OverlayAddressPool struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	IPType             string `json:"IPType,omitempty"`
	Name               string `json:"name,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	Description        string `json:"description,omitempty"`
	EndAddressRange    string `json:"endAddressRange,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	AssociatedDomainID string `json:"associatedDomainID,omitempty"`
	StartAddressRange  string `json:"startAddressRange,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewOverlayAddressPool returns a new *OverlayAddressPool
func NewOverlayAddressPool() *OverlayAddressPool {

	return &OverlayAddressPool{}
}

// Identity returns the Identity of the object.
func (o *OverlayAddressPool) Identity() bambou.Identity {

	return OverlayAddressPoolIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OverlayAddressPool) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OverlayAddressPool) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OverlayAddressPool from the server
func (o *OverlayAddressPool) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OverlayAddressPool into the server
func (o *OverlayAddressPool) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OverlayAddressPool from the server
func (o *OverlayAddressPool) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the OverlayAddressPool
func (o *OverlayAddressPool) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the OverlayAddressPool
func (o *OverlayAddressPool) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the OverlayAddressPool
func (o *OverlayAddressPool) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the OverlayAddressPool
func (o *OverlayAddressPool) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// OverlayPATNATEntries retrieves the list of child OverlayPATNATEntries of the OverlayAddressPool
func (o *OverlayAddressPool) OverlayPATNATEntries(info *bambou.FetchingInfo) (OverlayPATNATEntriesList, *bambou.Error) {

	var list OverlayPATNATEntriesList
	err := bambou.CurrentSession().FetchChildren(o, OverlayPATNATEntryIdentity, &list, info)
	return list, err
}

// CreateOverlayPATNATEntry creates a new child OverlayPATNATEntry under the OverlayAddressPool
func (o *OverlayAddressPool) CreateOverlayPATNATEntry(child *OverlayPATNATEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
