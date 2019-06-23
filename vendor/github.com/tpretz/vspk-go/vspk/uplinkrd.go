/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// UplinkRDIdentity represents the Identity of the object
var UplinkRDIdentity = bambou.Identity{
	Name:     "uplinkroutedistinguisher",
	Category: "uplinkroutedistinguishers",
}

// UplinkRDsList represents a list of UplinkRDs
type UplinkRDsList []*UplinkRD

// UplinkRDsAncestor is the interface that an ancestor of a UplinkRD must implement.
// An Ancestor is defined as an entity that has UplinkRD as a descendant.
// An Ancestor can get a list of its child UplinkRDs, but not necessarily create one.
type UplinkRDsAncestor interface {
	UplinkRDs(*bambou.FetchingInfo) (UplinkRDsList, *bambou.Error)
}

// UplinkRDsParent is the interface that a parent of a UplinkRD must implement.
// A Parent is defined as an entity that has UplinkRD as a child.
// A Parent is an Ancestor which can create a UplinkRD.
type UplinkRDsParent interface {
	UplinkRDsAncestor
	CreateUplinkRD(*UplinkRD) *bambou.Error
}

// UplinkRD represents the model of a uplinkroutedistinguisher
type UplinkRD struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`
	UplinkType         string `json:"uplinkType,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewUplinkRD returns a new *UplinkRD
func NewUplinkRD() *UplinkRD {

	return &UplinkRD{
		UplinkType: "RD_PRIMARY1",
	}
}

// Identity returns the Identity of the object.
func (o *UplinkRD) Identity() bambou.Identity {

	return UplinkRDIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UplinkRD) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UplinkRD) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the UplinkRD from the server
func (o *UplinkRD) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the UplinkRD into the server
func (o *UplinkRD) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the UplinkRD from the server
func (o *UplinkRD) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the UplinkRD
func (o *UplinkRD) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the UplinkRD
func (o *UplinkRD) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the UplinkRD
func (o *UplinkRD) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the UplinkRD
func (o *UplinkRD) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
