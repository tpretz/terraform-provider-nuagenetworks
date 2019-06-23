/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VNFMetadataIdentity represents the Identity of the object
var VNFMetadataIdentity = bambou.Identity{
	Name:     "vnfmetadata",
	Category: "vnfmetadatas",
}

// VNFMetadatasList represents a list of VNFMetadatas
type VNFMetadatasList []*VNFMetadata

// VNFMetadatasAncestor is the interface that an ancestor of a VNFMetadata must implement.
// An Ancestor is defined as an entity that has VNFMetadata as a descendant.
// An Ancestor can get a list of its child VNFMetadatas, but not necessarily create one.
type VNFMetadatasAncestor interface {
	VNFMetadatas(*bambou.FetchingInfo) (VNFMetadatasList, *bambou.Error)
}

// VNFMetadatasParent is the interface that a parent of a VNFMetadata must implement.
// A Parent is defined as an entity that has VNFMetadata as a child.
// A Parent is an Ancestor which can create a VNFMetadata.
type VNFMetadatasParent interface {
	VNFMetadatasAncestor
	CreateVNFMetadata(*VNFMetadata) *bambou.Error
}

// VNFMetadata represents the model of a vnfmetadata
type VNFMetadata struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Blob        string `json:"blob,omitempty"`
}

// NewVNFMetadata returns a new *VNFMetadata
func NewVNFMetadata() *VNFMetadata {

	return &VNFMetadata{}
}

// Identity returns the Identity of the object.
func (o *VNFMetadata) Identity() bambou.Identity {

	return VNFMetadataIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNFMetadata) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNFMetadata) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNFMetadata from the server
func (o *VNFMetadata) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNFMetadata into the server
func (o *VNFMetadata) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNFMetadata from the server
func (o *VNFMetadata) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
