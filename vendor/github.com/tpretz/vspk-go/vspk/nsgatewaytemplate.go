/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSGatewayTemplateIdentity represents the Identity of the object
var NSGatewayTemplateIdentity = bambou.Identity{
	Name:     "nsgatewaytemplate",
	Category: "nsgatewaytemplates",
}

// NSGatewayTemplatesList represents a list of NSGatewayTemplates
type NSGatewayTemplatesList []*NSGatewayTemplate

// NSGatewayTemplatesAncestor is the interface that an ancestor of a NSGatewayTemplate must implement.
// An Ancestor is defined as an entity that has NSGatewayTemplate as a descendant.
// An Ancestor can get a list of its child NSGatewayTemplates, but not necessarily create one.
type NSGatewayTemplatesAncestor interface {
	NSGatewayTemplates(*bambou.FetchingInfo) (NSGatewayTemplatesList, *bambou.Error)
}

// NSGatewayTemplatesParent is the interface that a parent of a NSGatewayTemplate must implement.
// A Parent is defined as an entity that has NSGatewayTemplate as a child.
// A Parent is an Ancestor which can create a NSGatewayTemplate.
type NSGatewayTemplatesParent interface {
	NSGatewayTemplatesAncestor
	CreateNSGatewayTemplate(*NSGatewayTemplate) *bambou.Error
}

// NSGatewayTemplate represents the model of a nsgatewaytemplate
type NSGatewayTemplate struct {
	ID                            string `json:"ID,omitempty"`
	ParentID                      string `json:"parentID,omitempty"`
	ParentType                    string `json:"parentType,omitempty"`
	Owner                         string `json:"owner,omitempty"`
	SSHService                    string `json:"SSHService,omitempty"`
	Name                          string `json:"name,omitempty"`
	LastUpdatedBy                 string `json:"lastUpdatedBy,omitempty"`
	Personality                   string `json:"personality,omitempty"`
	Description                   string `json:"description,omitempty"`
	InfrastructureAccessProfileID string `json:"infrastructureAccessProfileID,omitempty"`
	InfrastructureProfileID       string `json:"infrastructureProfileID,omitempty"`
	InstanceSSHOverride           string `json:"instanceSSHOverride,omitempty"`
	EnterpriseID                  string `json:"enterpriseID,omitempty"`
	EntityScope                   string `json:"entityScope,omitempty"`
	ExternalID                    string `json:"externalID,omitempty"`
}

// NewNSGatewayTemplate returns a new *NSGatewayTemplate
func NewNSGatewayTemplate() *NSGatewayTemplate {

	return &NSGatewayTemplate{
		SSHService:          "ENABLED",
		InstanceSSHOverride: "DISALLOWED",
	}
}

// Identity returns the Identity of the object.
func (o *NSGatewayTemplate) Identity() bambou.Identity {

	return NSGatewayTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGatewayTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGatewayTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGatewayTemplate from the server
func (o *NSGatewayTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGatewayTemplate into the server
func (o *NSGatewayTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGatewayTemplate from the server
func (o *NSGatewayTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NSGatewayTemplate
func (o *NSGatewayTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSGatewayTemplate
func (o *NSGatewayTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSGatewayTemplate
func (o *NSGatewayTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSGatewayTemplate
func (o *NSGatewayTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NSPortTemplates retrieves the list of child NSPortTemplates of the NSGatewayTemplate
func (o *NSGatewayTemplate) NSPortTemplates(info *bambou.FetchingInfo) (NSPortTemplatesList, *bambou.Error) {

	var list NSPortTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, NSPortTemplateIdentity, &list, info)
	return list, err
}

// CreateNSPortTemplate creates a new child NSPortTemplate under the NSGatewayTemplate
func (o *NSGatewayTemplate) CreateNSPortTemplate(child *NSPortTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
