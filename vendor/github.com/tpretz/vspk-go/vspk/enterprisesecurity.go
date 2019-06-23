/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// EnterpriseSecurityIdentity represents the Identity of the object
var EnterpriseSecurityIdentity = bambou.Identity{
	Name:     "enterprisesecurity",
	Category: "enterprisesecurities",
}

// EnterpriseSecuritiesList represents a list of EnterpriseSecurities
type EnterpriseSecuritiesList []*EnterpriseSecurity

// EnterpriseSecuritiesAncestor is the interface that an ancestor of a EnterpriseSecurity must implement.
// An Ancestor is defined as an entity that has EnterpriseSecurity as a descendant.
// An Ancestor can get a list of its child EnterpriseSecurities, but not necessarily create one.
type EnterpriseSecuritiesAncestor interface {
	EnterpriseSecurities(*bambou.FetchingInfo) (EnterpriseSecuritiesList, *bambou.Error)
}

// EnterpriseSecuritiesParent is the interface that a parent of a EnterpriseSecurity must implement.
// A Parent is defined as an entity that has EnterpriseSecurity as a child.
// A Parent is an Ancestor which can create a EnterpriseSecurity.
type EnterpriseSecuritiesParent interface {
	EnterpriseSecuritiesAncestor
	CreateEnterpriseSecurity(*EnterpriseSecurity) *bambou.Error
}

// EnterpriseSecurity represents the model of a enterprisesecurity
type EnterpriseSecurity struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	GatewaySecurityRevision int    `json:"gatewaySecurityRevision,omitempty"`
	Revision                int    `json:"revision,omitempty"`
	EnterpriseID            string `json:"enterpriseID,omitempty"`
	EntityScope             string `json:"entityScope,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
}

// NewEnterpriseSecurity returns a new *EnterpriseSecurity
func NewEnterpriseSecurity() *EnterpriseSecurity {

	return &EnterpriseSecurity{}
}

// Identity returns the Identity of the object.
func (o *EnterpriseSecurity) Identity() bambou.Identity {

	return EnterpriseSecurityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnterpriseSecurity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnterpriseSecurity) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EnterpriseSecurity from the server
func (o *EnterpriseSecurity) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EnterpriseSecurity into the server
func (o *EnterpriseSecurity) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EnterpriseSecurity from the server
func (o *EnterpriseSecurity) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EnterpriseSecurity
func (o *EnterpriseSecurity) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EnterpriseSecurity
func (o *EnterpriseSecurity) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EnterpriseSecurity
func (o *EnterpriseSecurity) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EnterpriseSecurity
func (o *EnterpriseSecurity) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterpriseSecuredDatas retrieves the list of child EnterpriseSecuredDatas of the EnterpriseSecurity
func (o *EnterpriseSecurity) EnterpriseSecuredDatas(info *bambou.FetchingInfo) (EnterpriseSecuredDatasList, *bambou.Error) {

	var list EnterpriseSecuredDatasList
	err := bambou.CurrentSession().FetchChildren(o, EnterpriseSecuredDataIdentity, &list, info)
	return list, err
}

// CreateEnterpriseSecuredData creates a new child EnterpriseSecuredData under the EnterpriseSecurity
func (o *EnterpriseSecurity) CreateEnterpriseSecuredData(child *EnterpriseSecuredData) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
