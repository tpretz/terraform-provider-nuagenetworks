/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// UnderlayIdentity represents the Identity of the object
var UnderlayIdentity = bambou.Identity{
	Name:     "underlay",
	Category: "underlays",
}

// UnderlaysList represents a list of Underlays
type UnderlaysList []*Underlay

// UnderlaysAncestor is the interface that an ancestor of a Underlay must implement.
// An Ancestor is defined as an entity that has Underlay as a descendant.
// An Ancestor can get a list of its child Underlays, but not necessarily create one.
type UnderlaysAncestor interface {
	Underlays(*bambou.FetchingInfo) (UnderlaysList, *bambou.Error)
}

// UnderlaysParent is the interface that a parent of a Underlay must implement.
// A Parent is defined as an entity that has Underlay as a child.
// A Parent is an Ancestor which can create a Underlay.
type UnderlaysParent interface {
	UnderlaysAncestor
	CreateUnderlay(*Underlay) *bambou.Error
}

// Underlay represents the model of a underlay
type Underlay struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	UnderlayId    int    `json:"underlayId,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewUnderlay returns a new *Underlay
func NewUnderlay() *Underlay {

	return &Underlay{}
}

// Identity returns the Identity of the object.
func (o *Underlay) Identity() bambou.Identity {

	return UnderlayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Underlay) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Underlay) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Underlay from the server
func (o *Underlay) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Underlay into the server
func (o *Underlay) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Underlay from the server
func (o *Underlay) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Underlay
func (o *Underlay) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Underlay
func (o *Underlay) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Underlay
func (o *Underlay) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Underlay
func (o *Underlay) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
