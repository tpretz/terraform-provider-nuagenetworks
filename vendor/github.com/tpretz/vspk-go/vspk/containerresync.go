/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// ContainerResyncIdentity represents the Identity of the object
var ContainerResyncIdentity = bambou.Identity{
	Name:     "containerresync",
	Category: "containerresync",
}

// ContainerResyncsList represents a list of ContainerResyncs
type ContainerResyncsList []*ContainerResync

// ContainerResyncsAncestor is the interface that an ancestor of a ContainerResync must implement.
// An Ancestor is defined as an entity that has ContainerResync as a descendant.
// An Ancestor can get a list of its child ContainerResyncs, but not necessarily create one.
type ContainerResyncsAncestor interface {
	ContainerResyncs(*bambou.FetchingInfo) (ContainerResyncsList, *bambou.Error)
}

// ContainerResyncsParent is the interface that a parent of a ContainerResync must implement.
// A Parent is defined as an entity that has ContainerResync as a child.
// A Parent is an Ancestor which can create a ContainerResync.
type ContainerResyncsParent interface {
	ContainerResyncsAncestor
	CreateContainerResync(*ContainerResync) *bambou.Error
}

// ContainerResync represents the model of a containerresync
type ContainerResync struct {
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

// NewContainerResync returns a new *ContainerResync
func NewContainerResync() *ContainerResync {

	return &ContainerResync{}
}

// Identity returns the Identity of the object.
func (o *ContainerResync) Identity() bambou.Identity {

	return ContainerResyncIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ContainerResync) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ContainerResync) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ContainerResync from the server
func (o *ContainerResync) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ContainerResync into the server
func (o *ContainerResync) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ContainerResync from the server
func (o *ContainerResync) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ContainerResync
func (o *ContainerResync) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ContainerResync
func (o *ContainerResync) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ContainerResync
func (o *ContainerResync) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ContainerResync
func (o *ContainerResync) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
