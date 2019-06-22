/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VNFDescriptorIdentity represents the Identity of the object
var VNFDescriptorIdentity = bambou.Identity{
	Name:     "vnfdescriptor",
	Category: "vnfdescriptors",
}

// VNFDescriptorsList represents a list of VNFDescriptors
type VNFDescriptorsList []*VNFDescriptor

// VNFDescriptorsAncestor is the interface that an ancestor of a VNFDescriptor must implement.
// An Ancestor is defined as an entity that has VNFDescriptor as a descendant.
// An Ancestor can get a list of its child VNFDescriptors, but not necessarily create one.
type VNFDescriptorsAncestor interface {
	VNFDescriptors(*bambou.FetchingInfo) (VNFDescriptorsList, *bambou.Error)
}

// VNFDescriptorsParent is the interface that a parent of a VNFDescriptor must implement.
// A Parent is defined as an entity that has VNFDescriptor as a child.
// A Parent is an Ancestor which can create a VNFDescriptor.
type VNFDescriptorsParent interface {
	VNFDescriptorsAncestor
	CreateVNFDescriptor(*VNFDescriptor) *bambou.Error
}

// VNFDescriptor represents the model of a vnfdescriptor
type VNFDescriptor struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	CPUCount                       int    `json:"CPUCount,omitempty"`
	Name                           string `json:"name,omitempty"`
	MemoryMB                       int    `json:"memoryMB,omitempty"`
	Vendor                         string `json:"vendor,omitempty"`
	Description                    string `json:"description,omitempty"`
	MetadataID                     string `json:"metadataID,omitempty"`
	Visible                        bool   `json:"visible"`
	EntityScope                    string `json:"entityScope,omitempty"`
	AssociatedVNFThresholdPolicyID string `json:"associatedVNFThresholdPolicyID,omitempty"`
	StorageGB                      int    `json:"storageGB,omitempty"`
	ExternalID                     string `json:"externalID,omitempty"`
	Type                           string `json:"type,omitempty"`
}

// NewVNFDescriptor returns a new *VNFDescriptor
func NewVNFDescriptor() *VNFDescriptor {

	return &VNFDescriptor{
		Visible: true,
		Type:    "FIREWALL",
	}
}

// Identity returns the Identity of the object.
func (o *VNFDescriptor) Identity() bambou.Identity {

	return VNFDescriptorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNFDescriptor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNFDescriptor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNFDescriptor from the server
func (o *VNFDescriptor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNFDescriptor into the server
func (o *VNFDescriptor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNFDescriptor from the server
func (o *VNFDescriptor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VNFDescriptor
func (o *VNFDescriptor) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VNFDescriptor
func (o *VNFDescriptor) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VNFDescriptor
func (o *VNFDescriptor) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VNFDescriptor
func (o *VNFDescriptor) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VNFInterfaceDescriptors retrieves the list of child VNFInterfaceDescriptors of the VNFDescriptor
func (o *VNFDescriptor) VNFInterfaceDescriptors(info *bambou.FetchingInfo) (VNFInterfaceDescriptorsList, *bambou.Error) {

	var list VNFInterfaceDescriptorsList
	err := bambou.CurrentSession().FetchChildren(o, VNFInterfaceDescriptorIdentity, &list, info)
	return list, err
}

// CreateVNFInterfaceDescriptor creates a new child VNFInterfaceDescriptor under the VNFDescriptor
func (o *VNFDescriptor) CreateVNFInterfaceDescriptor(child *VNFInterfaceDescriptor) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
