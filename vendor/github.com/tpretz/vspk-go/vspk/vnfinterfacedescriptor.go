/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VNFInterfaceDescriptorIdentity represents the Identity of the object
var VNFInterfaceDescriptorIdentity = bambou.Identity{
	Name:     "vnfinterfacedescriptor",
	Category: "vnfinterfacedescriptors",
}

// VNFInterfaceDescriptorsList represents a list of VNFInterfaceDescriptors
type VNFInterfaceDescriptorsList []*VNFInterfaceDescriptor

// VNFInterfaceDescriptorsAncestor is the interface that an ancestor of a VNFInterfaceDescriptor must implement.
// An Ancestor is defined as an entity that has VNFInterfaceDescriptor as a descendant.
// An Ancestor can get a list of its child VNFInterfaceDescriptors, but not necessarily create one.
type VNFInterfaceDescriptorsAncestor interface {
	VNFInterfaceDescriptors(*bambou.FetchingInfo) (VNFInterfaceDescriptorsList, *bambou.Error)
}

// VNFInterfaceDescriptorsParent is the interface that a parent of a VNFInterfaceDescriptor must implement.
// A Parent is defined as an entity that has VNFInterfaceDescriptor as a child.
// A Parent is an Ancestor which can create a VNFInterfaceDescriptor.
type VNFInterfaceDescriptorsParent interface {
	VNFInterfaceDescriptorsAncestor
	CreateVNFInterfaceDescriptor(*VNFInterfaceDescriptor) *bambou.Error
}

// VNFInterfaceDescriptor represents the model of a vnfinterfacedescriptor
type VNFInterfaceDescriptor struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	EntityScope string `json:"entityScope,omitempty"`
	ExternalID  string `json:"externalID,omitempty"`
	Type        string `json:"type,omitempty"`
}

// NewVNFInterfaceDescriptor returns a new *VNFInterfaceDescriptor
func NewVNFInterfaceDescriptor() *VNFInterfaceDescriptor {

	return &VNFInterfaceDescriptor{
		Type: "MANAGEMENT",
	}
}

// Identity returns the Identity of the object.
func (o *VNFInterfaceDescriptor) Identity() bambou.Identity {

	return VNFInterfaceDescriptorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNFInterfaceDescriptor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNFInterfaceDescriptor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNFInterfaceDescriptor from the server
func (o *VNFInterfaceDescriptor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNFInterfaceDescriptor into the server
func (o *VNFInterfaceDescriptor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNFInterfaceDescriptor from the server
func (o *VNFInterfaceDescriptor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VNFInterfaceDescriptor
func (o *VNFInterfaceDescriptor) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VNFInterfaceDescriptor
func (o *VNFInterfaceDescriptor) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VNFInterfaceDescriptor
func (o *VNFInterfaceDescriptor) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VNFInterfaceDescriptor
func (o *VNFInterfaceDescriptor) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
