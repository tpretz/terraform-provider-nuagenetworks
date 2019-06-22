/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SaaSApplicationTypeIdentity represents the Identity of the object
var SaaSApplicationTypeIdentity = bambou.Identity{
	Name:     "saasapplicationtype",
	Category: "saasapplicationtypes",
}

// SaaSApplicationTypesList represents a list of SaaSApplicationTypes
type SaaSApplicationTypesList []*SaaSApplicationType

// SaaSApplicationTypesAncestor is the interface that an ancestor of a SaaSApplicationType must implement.
// An Ancestor is defined as an entity that has SaaSApplicationType as a descendant.
// An Ancestor can get a list of its child SaaSApplicationTypes, but not necessarily create one.
type SaaSApplicationTypesAncestor interface {
	SaaSApplicationTypes(*bambou.FetchingInfo) (SaaSApplicationTypesList, *bambou.Error)
}

// SaaSApplicationTypesParent is the interface that a parent of a SaaSApplicationType must implement.
// A Parent is defined as an entity that has SaaSApplicationType as a child.
// A Parent is an Ancestor which can create a SaaSApplicationType.
type SaaSApplicationTypesParent interface {
	SaaSApplicationTypesAncestor
	CreateSaaSApplicationType(*SaaSApplicationType) *bambou.Error
}

// SaaSApplicationType represents the model of a saasapplicationtype
type SaaSApplicationType struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	ReadOnly      bool   `json:"readOnly"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewSaaSApplicationType returns a new *SaaSApplicationType
func NewSaaSApplicationType() *SaaSApplicationType {

	return &SaaSApplicationType{
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *SaaSApplicationType) Identity() bambou.Identity {

	return SaaSApplicationTypeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SaaSApplicationType) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SaaSApplicationType) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SaaSApplicationType from the server
func (o *SaaSApplicationType) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SaaSApplicationType into the server
func (o *SaaSApplicationType) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SaaSApplicationType from the server
func (o *SaaSApplicationType) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the SaaSApplicationType
func (o *SaaSApplicationType) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SaaSApplicationType
func (o *SaaSApplicationType) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SaaSApplicationType
func (o *SaaSApplicationType) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SaaSApplicationType
func (o *SaaSApplicationType) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterpriseNetworks retrieves the list of child EnterpriseNetworks of the SaaSApplicationType
func (o *SaaSApplicationType) EnterpriseNetworks(info *bambou.FetchingInfo) (EnterpriseNetworksList, *bambou.Error) {

	var list EnterpriseNetworksList
	err := bambou.CurrentSession().FetchChildren(o, EnterpriseNetworkIdentity, &list, info)
	return list, err
}

// CreateEnterpriseNetwork creates a new child EnterpriseNetwork under the SaaSApplicationType
func (o *SaaSApplicationType) CreateEnterpriseNetwork(child *EnterpriseNetwork) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
