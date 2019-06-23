/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VNFCatalogIdentity represents the Identity of the object
var VNFCatalogIdentity = bambou.Identity{
	Name:     "vnfcatalog",
	Category: "vnfcatalogs",
}

// VNFCatalogsList represents a list of VNFCatalogs
type VNFCatalogsList []*VNFCatalog

// VNFCatalogsAncestor is the interface that an ancestor of a VNFCatalog must implement.
// An Ancestor is defined as an entity that has VNFCatalog as a descendant.
// An Ancestor can get a list of its child VNFCatalogs, but not necessarily create one.
type VNFCatalogsAncestor interface {
	VNFCatalogs(*bambou.FetchingInfo) (VNFCatalogsList, *bambou.Error)
}

// VNFCatalogsParent is the interface that a parent of a VNFCatalog must implement.
// A Parent is defined as an entity that has VNFCatalog as a child.
// A Parent is an Ancestor which can create a VNFCatalog.
type VNFCatalogsParent interface {
	VNFCatalogsAncestor
	CreateVNFCatalog(*VNFCatalog) *bambou.Error
}

// VNFCatalog represents the model of a vnfcatalog
type VNFCatalog struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// NewVNFCatalog returns a new *VNFCatalog
func NewVNFCatalog() *VNFCatalog {

	return &VNFCatalog{}
}

// Identity returns the Identity of the object.
func (o *VNFCatalog) Identity() bambou.Identity {

	return VNFCatalogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VNFCatalog) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VNFCatalog) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VNFCatalog from the server
func (o *VNFCatalog) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VNFCatalog into the server
func (o *VNFCatalog) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VNFCatalog from the server
func (o *VNFCatalog) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VNFDescriptors retrieves the list of child VNFDescriptors of the VNFCatalog
func (o *VNFCatalog) VNFDescriptors(info *bambou.FetchingInfo) (VNFDescriptorsList, *bambou.Error) {

	var list VNFDescriptorsList
	err := bambou.CurrentSession().FetchChildren(o, VNFDescriptorIdentity, &list, info)
	return list, err
}

// CreateVNFDescriptor creates a new child VNFDescriptor under the VNFCatalog
func (o *VNFCatalog) CreateVNFDescriptor(child *VNFDescriptor) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
