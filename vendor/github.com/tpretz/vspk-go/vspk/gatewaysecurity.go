/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GatewaySecurityIdentity represents the Identity of the object
var GatewaySecurityIdentity = bambou.Identity{
	Name:     "gatewaysecurity",
	Category: "gatewaysecurities",
}

// GatewaySecuritiesList represents a list of GatewaySecurities
type GatewaySecuritiesList []*GatewaySecurity

// GatewaySecuritiesAncestor is the interface that an ancestor of a GatewaySecurity must implement.
// An Ancestor is defined as an entity that has GatewaySecurity as a descendant.
// An Ancestor can get a list of its child GatewaySecurities, but not necessarily create one.
type GatewaySecuritiesAncestor interface {
	GatewaySecurities(*bambou.FetchingInfo) (GatewaySecuritiesList, *bambou.Error)
}

// GatewaySecuritiesParent is the interface that a parent of a GatewaySecurity must implement.
// A Parent is defined as an entity that has GatewaySecurity as a child.
// A Parent is an Ancestor which can create a GatewaySecurity.
type GatewaySecuritiesParent interface {
	GatewaySecuritiesAncestor
	CreateGatewaySecurity(*GatewaySecurity) *bambou.Error
}

// GatewaySecurity represents the model of a gatewaysecurity
type GatewaySecurity struct {
	ID                   string `json:"ID,omitempty"`
	ParentID             string `json:"parentID,omitempty"`
	ParentType           string `json:"parentType,omitempty"`
	Owner                string `json:"owner,omitempty"`
	LastUpdatedBy        string `json:"lastUpdatedBy,omitempty"`
	GatewayID            string `json:"gatewayID,omitempty"`
	Revision             int    `json:"revision,omitempty"`
	EntityScope          string `json:"entityScope,omitempty"`
	AssociatedEntityType string `json:"associatedEntityType,omitempty"`
	ExternalID           string `json:"externalID,omitempty"`
}

// NewGatewaySecurity returns a new *GatewaySecurity
func NewGatewaySecurity() *GatewaySecurity {

	return &GatewaySecurity{}
}

// Identity returns the Identity of the object.
func (o *GatewaySecurity) Identity() bambou.Identity {

	return GatewaySecurityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewaySecurity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewaySecurity) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewaySecurity from the server
func (o *GatewaySecurity) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewaySecurity into the server
func (o *GatewaySecurity) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewaySecurity from the server
func (o *GatewaySecurity) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// GatewaySecuredDatas retrieves the list of child GatewaySecuredDatas of the GatewaySecurity
func (o *GatewaySecurity) GatewaySecuredDatas(info *bambou.FetchingInfo) (GatewaySecuredDatasList, *bambou.Error) {

	var list GatewaySecuredDatasList
	err := bambou.CurrentSession().FetchChildren(o, GatewaySecuredDataIdentity, &list, info)
	return list, err
}

// CreateGatewaySecuredData creates a new child GatewaySecuredData under the GatewaySecurity
func (o *GatewaySecurity) CreateGatewaySecuredData(child *GatewaySecuredData) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the GatewaySecurity
func (o *GatewaySecurity) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GatewaySecurity
func (o *GatewaySecurity) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GatewaySecurity
func (o *GatewaySecurity) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GatewaySecurity
func (o *GatewaySecurity) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
