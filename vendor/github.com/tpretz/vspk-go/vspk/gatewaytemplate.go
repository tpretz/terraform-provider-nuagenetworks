/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GatewayTemplateIdentity represents the Identity of the object
var GatewayTemplateIdentity = bambou.Identity{
	Name:     "gatewaytemplate",
	Category: "gatewaytemplates",
}

// GatewayTemplatesList represents a list of GatewayTemplates
type GatewayTemplatesList []*GatewayTemplate

// GatewayTemplatesAncestor is the interface that an ancestor of a GatewayTemplate must implement.
// An Ancestor is defined as an entity that has GatewayTemplate as a descendant.
// An Ancestor can get a list of its child GatewayTemplates, but not necessarily create one.
type GatewayTemplatesAncestor interface {
	GatewayTemplates(*bambou.FetchingInfo) (GatewayTemplatesList, *bambou.Error)
}

// GatewayTemplatesParent is the interface that a parent of a GatewayTemplate must implement.
// A Parent is defined as an entity that has GatewayTemplate as a child.
// A Parent is an Ancestor which can create a GatewayTemplate.
type GatewayTemplatesParent interface {
	GatewayTemplatesAncestor
	CreateGatewayTemplate(*GatewayTemplate) *bambou.Error
}

// GatewayTemplate represents the model of a gatewaytemplate
type GatewayTemplate struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	Name                    string `json:"name,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	Personality             string `json:"personality,omitempty"`
	Description             string `json:"description,omitempty"`
	InfrastructureProfileID string `json:"infrastructureProfileID,omitempty"`
	EnterpriseID            string `json:"enterpriseID,omitempty"`
	EntityScope             string `json:"entityScope,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
}

// NewGatewayTemplate returns a new *GatewayTemplate
func NewGatewayTemplate() *GatewayTemplate {

	return &GatewayTemplate{}
}

// Identity returns the Identity of the object.
func (o *GatewayTemplate) Identity() bambou.Identity {

	return GatewayTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewayTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewayTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewayTemplate from the server
func (o *GatewayTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewayTemplate into the server
func (o *GatewayTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewayTemplate from the server
func (o *GatewayTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GatewayTemplate
func (o *GatewayTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GatewayTemplate
func (o *GatewayTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GatewayTemplate
func (o *GatewayTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GatewayTemplate
func (o *GatewayTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PortTemplates retrieves the list of child PortTemplates of the GatewayTemplate
func (o *GatewayTemplate) PortTemplates(info *bambou.FetchingInfo) (PortTemplatesList, *bambou.Error) {

	var list PortTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, PortTemplateIdentity, &list, info)
	return list, err
}

// CreatePortTemplate creates a new child PortTemplate under the GatewayTemplate
func (o *GatewayTemplate) CreatePortTemplate(child *PortTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
