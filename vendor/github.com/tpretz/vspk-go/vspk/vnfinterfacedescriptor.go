/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

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
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
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
