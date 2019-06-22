/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKEGatewayConfigIdentity represents the Identity of the object
var IKEGatewayConfigIdentity = bambou.Identity{
	Name:     "ikegatewayconfig",
	Category: "ikegatewayconfig",
}

// IKEGatewayConfigsList represents a list of IKEGatewayConfigs
type IKEGatewayConfigsList []*IKEGatewayConfig

// IKEGatewayConfigsAncestor is the interface that an ancestor of a IKEGatewayConfig must implement.
// An Ancestor is defined as an entity that has IKEGatewayConfig as a descendant.
// An Ancestor can get a list of its child IKEGatewayConfigs, but not necessarily create one.
type IKEGatewayConfigsAncestor interface {
	IKEGatewayConfigs(*bambou.FetchingInfo) (IKEGatewayConfigsList, *bambou.Error)
}

// IKEGatewayConfigsParent is the interface that a parent of a IKEGatewayConfig must implement.
// A Parent is defined as an entity that has IKEGatewayConfig as a child.
// A Parent is an Ancestor which can create a IKEGatewayConfig.
type IKEGatewayConfigsParent interface {
	IKEGatewayConfigsAncestor
	CreateIKEGatewayConfig(*IKEGatewayConfig) *bambou.Error
}

// IKEGatewayConfig represents the model of a ikegatewayconfig
type IKEGatewayConfig struct {
	ID            string      `json:"ID,omitempty"`
	ParentID      string      `json:"parentID,omitempty"`
	ParentType    string      `json:"parentType,omitempty"`
	Owner         string      `json:"owner,omitempty"`
	LastUpdatedBy string      `json:"lastUpdatedBy,omitempty"`
	EntityScope   string      `json:"entityScope,omitempty"`
	Config        interface{} `json:"config,omitempty"`
	ExternalID    string      `json:"externalID,omitempty"`
}

// NewIKEGatewayConfig returns a new *IKEGatewayConfig
func NewIKEGatewayConfig() *IKEGatewayConfig {

	return &IKEGatewayConfig{}
}

// Identity returns the Identity of the object.
func (o *IKEGatewayConfig) Identity() bambou.Identity {

	return IKEGatewayConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEGatewayConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEGatewayConfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEGatewayConfig from the server
func (o *IKEGatewayConfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEGatewayConfig into the server
func (o *IKEGatewayConfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEGatewayConfig from the server
func (o *IKEGatewayConfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEGatewayConfig
func (o *IKEGatewayConfig) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEGatewayConfig
func (o *IKEGatewayConfig) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEGatewayConfig
func (o *IKEGatewayConfig) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEGatewayConfig
func (o *IKEGatewayConfig) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
