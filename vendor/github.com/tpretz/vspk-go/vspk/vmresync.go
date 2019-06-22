/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VMResyncIdentity represents the Identity of the object
var VMResyncIdentity = bambou.Identity{
	Name:     "resync",
	Category: "resync",
}

// VMResyncsList represents a list of VMResyncs
type VMResyncsList []*VMResync

// VMResyncsAncestor is the interface that an ancestor of a VMResync must implement.
// An Ancestor is defined as an entity that has VMResync as a descendant.
// An Ancestor can get a list of its child VMResyncs, but not necessarily create one.
type VMResyncsAncestor interface {
	VMResyncs(*bambou.FetchingInfo) (VMResyncsList, *bambou.Error)
}

// VMResyncsParent is the interface that a parent of a VMResync must implement.
// A Parent is defined as an entity that has VMResync as a child.
// A Parent is an Ancestor which can create a VMResync.
type VMResyncsParent interface {
	VMResyncsAncestor
	CreateVMResync(*VMResync) *bambou.Error
}

// VMResync represents the model of a resync
type VMResync struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	LastRequestTimestamp    int    `json:"lastRequestTimestamp,omitempty"`
	LastTimeResyncInitiated int    `json:"lastTimeResyncInitiated,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	EntityScope             string `json:"entityScope,omitempty"`
	Status                  string `json:"status,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
}

// NewVMResync returns a new *VMResync
func NewVMResync() *VMResync {

	return &VMResync{}
}

// Identity returns the Identity of the object.
func (o *VMResync) Identity() bambou.Identity {

	return VMResyncIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VMResync) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VMResync) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VMResync from the server
func (o *VMResync) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VMResync into the server
func (o *VMResync) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VMResync from the server
func (o *VMResync) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VMResync
func (o *VMResync) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VMResync
func (o *VMResync) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VMResync
func (o *VMResync) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VMResync
func (o *VMResync) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
