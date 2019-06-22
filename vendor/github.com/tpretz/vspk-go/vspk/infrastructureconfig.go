/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// InfrastructureConfigIdentity represents the Identity of the object
var InfrastructureConfigIdentity = bambou.Identity{
	Name:     "infraconfig",
	Category: "infraconfig",
}

// InfrastructureConfigsList represents a list of InfrastructureConfigs
type InfrastructureConfigsList []*InfrastructureConfig

// InfrastructureConfigsAncestor is the interface that an ancestor of a InfrastructureConfig must implement.
// An Ancestor is defined as an entity that has InfrastructureConfig as a descendant.
// An Ancestor can get a list of its child InfrastructureConfigs, but not necessarily create one.
type InfrastructureConfigsAncestor interface {
	InfrastructureConfigs(*bambou.FetchingInfo) (InfrastructureConfigsList, *bambou.Error)
}

// InfrastructureConfigsParent is the interface that a parent of a InfrastructureConfig must implement.
// A Parent is defined as an entity that has InfrastructureConfig as a child.
// A Parent is an Ancestor which can create a InfrastructureConfig.
type InfrastructureConfigsParent interface {
	InfrastructureConfigsAncestor
	CreateInfrastructureConfig(*InfrastructureConfig) *bambou.Error
}

// InfrastructureConfig represents the model of a infraconfig
type InfrastructureConfig struct {
	ID            string      `json:"ID,omitempty"`
	ParentID      string      `json:"parentID,omitempty"`
	ParentType    string      `json:"parentType,omitempty"`
	Owner         string      `json:"owner,omitempty"`
	LastUpdatedBy string      `json:"lastUpdatedBy,omitempty"`
	EntityScope   string      `json:"entityScope,omitempty"`
	Config        interface{} `json:"config,omitempty"`
	ConfigStatus  string      `json:"configStatus,omitempty"`
	ExternalID    string      `json:"externalID,omitempty"`
}

// NewInfrastructureConfig returns a new *InfrastructureConfig
func NewInfrastructureConfig() *InfrastructureConfig {

	return &InfrastructureConfig{}
}

// Identity returns the Identity of the object.
func (o *InfrastructureConfig) Identity() bambou.Identity {

	return InfrastructureConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructureConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructureConfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the InfrastructureConfig from the server
func (o *InfrastructureConfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the InfrastructureConfig into the server
func (o *InfrastructureConfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the InfrastructureConfig from the server
func (o *InfrastructureConfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the InfrastructureConfig
func (o *InfrastructureConfig) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the InfrastructureConfig
func (o *InfrastructureConfig) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the InfrastructureConfig
func (o *InfrastructureConfig) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the InfrastructureConfig
func (o *InfrastructureConfig) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
