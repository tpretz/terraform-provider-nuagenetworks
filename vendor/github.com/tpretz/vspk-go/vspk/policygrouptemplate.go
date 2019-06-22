/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PolicyGroupTemplateIdentity represents the Identity of the object
var PolicyGroupTemplateIdentity = bambou.Identity{
	Name:     "policygrouptemplate",
	Category: "policygrouptemplates",
}

// PolicyGroupTemplatesList represents a list of PolicyGroupTemplates
type PolicyGroupTemplatesList []*PolicyGroupTemplate

// PolicyGroupTemplatesAncestor is the interface that an ancestor of a PolicyGroupTemplate must implement.
// An Ancestor is defined as an entity that has PolicyGroupTemplate as a descendant.
// An Ancestor can get a list of its child PolicyGroupTemplates, but not necessarily create one.
type PolicyGroupTemplatesAncestor interface {
	PolicyGroupTemplates(*bambou.FetchingInfo) (PolicyGroupTemplatesList, *bambou.Error)
}

// PolicyGroupTemplatesParent is the interface that a parent of a PolicyGroupTemplate must implement.
// A Parent is defined as an entity that has PolicyGroupTemplate as a child.
// A Parent is an Ancestor which can create a PolicyGroupTemplate.
type PolicyGroupTemplatesParent interface {
	PolicyGroupTemplatesAncestor
	CreatePolicyGroupTemplate(*PolicyGroupTemplate) *bambou.Error
}

// PolicyGroupTemplate represents the model of a policygrouptemplate
type PolicyGroupTemplate struct {
	ID               string `json:"ID,omitempty"`
	ParentID         string `json:"parentID,omitempty"`
	ParentType       string `json:"parentType,omitempty"`
	Owner            string `json:"owner,omitempty"`
	EVPNCommunityTag string `json:"EVPNCommunityTag,omitempty"`
	Name             string `json:"name,omitempty"`
	LastUpdatedBy    string `json:"lastUpdatedBy,omitempty"`
	Description      string `json:"description,omitempty"`
	EntityScope      string `json:"entityScope,omitempty"`
	External         bool   `json:"external"`
	ExternalID       string `json:"externalID,omitempty"`
	Type             string `json:"type,omitempty"`
}

// NewPolicyGroupTemplate returns a new *PolicyGroupTemplate
func NewPolicyGroupTemplate() *PolicyGroupTemplate {

	return &PolicyGroupTemplate{}
}

// Identity returns the Identity of the object.
func (o *PolicyGroupTemplate) Identity() bambou.Identity {

	return PolicyGroupTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyGroupTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyGroupTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyGroupTemplate from the server
func (o *PolicyGroupTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyGroupTemplate into the server
func (o *PolicyGroupTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyGroupTemplate from the server
func (o *PolicyGroupTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PolicyGroupTemplate
func (o *PolicyGroupTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyGroupTemplate
func (o *PolicyGroupTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyGroupTemplate
func (o *PolicyGroupTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyGroupTemplate
func (o *PolicyGroupTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the PolicyGroupTemplate
func (o *PolicyGroupTemplate) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
