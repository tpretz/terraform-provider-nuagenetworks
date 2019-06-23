/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// IKESubnetIdentity represents the Identity of the object
var IKESubnetIdentity = bambou.Identity{
	Name:     "ikesubnet",
	Category: "ikesubnets",
}

// IKESubnetsList represents a list of IKESubnets
type IKESubnetsList []*IKESubnet

// IKESubnetsAncestor is the interface that an ancestor of a IKESubnet must implement.
// An Ancestor is defined as an entity that has IKESubnet as a descendant.
// An Ancestor can get a list of its child IKESubnets, but not necessarily create one.
type IKESubnetsAncestor interface {
	IKESubnets(*bambou.FetchingInfo) (IKESubnetsList, *bambou.Error)
}

// IKESubnetsParent is the interface that a parent of a IKESubnet must implement.
// A Parent is defined as an entity that has IKESubnet as a child.
// A Parent is an Ancestor which can create a IKESubnet.
type IKESubnetsParent interface {
	IKESubnetsAncestor
	CreateIKESubnet(*IKESubnet) *bambou.Error
}

// IKESubnet represents the model of a ikesubnet
type IKESubnet struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	Prefix                 string `json:"prefix,omitempty"`
	AssociatedIKEGatewayID string `json:"associatedIKEGatewayID,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewIKESubnet returns a new *IKESubnet
func NewIKESubnet() *IKESubnet {

	return &IKESubnet{}
}

// Identity returns the Identity of the object.
func (o *IKESubnet) Identity() bambou.Identity {

	return IKESubnetIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKESubnet) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKESubnet) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKESubnet from the server
func (o *IKESubnet) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKESubnet into the server
func (o *IKESubnet) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKESubnet from the server
func (o *IKESubnet) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKESubnet
func (o *IKESubnet) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKESubnet
func (o *IKESubnet) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKESubnet
func (o *IKESubnet) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKESubnet
func (o *IKESubnet) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
