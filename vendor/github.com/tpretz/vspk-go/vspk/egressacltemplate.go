/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EgressACLTemplateIdentity represents the Identity of the object
var EgressACLTemplateIdentity = bambou.Identity{
	Name:     "egressacltemplate",
	Category: "egressacltemplates",
}

// EgressACLTemplatesList represents a list of EgressACLTemplates
type EgressACLTemplatesList []*EgressACLTemplate

// EgressACLTemplatesAncestor is the interface that an ancestor of a EgressACLTemplate must implement.
// An Ancestor is defined as an entity that has EgressACLTemplate as a descendant.
// An Ancestor can get a list of its child EgressACLTemplates, but not necessarily create one.
type EgressACLTemplatesAncestor interface {
	EgressACLTemplates(*bambou.FetchingInfo) (EgressACLTemplatesList, *bambou.Error)
}

// EgressACLTemplatesParent is the interface that a parent of a EgressACLTemplate must implement.
// A Parent is defined as an entity that has EgressACLTemplate as a child.
// A Parent is an Ancestor which can create a EgressACLTemplate.
type EgressACLTemplatesParent interface {
	EgressACLTemplatesAncestor
	CreateEgressACLTemplate(*EgressACLTemplate) *bambou.Error
}

// EgressACLTemplate represents the model of a egressacltemplate
type EgressACLTemplate struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	Name                              string `json:"name,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	Active                            bool   `json:"active"`
	DefaultAllowIP                    bool   `json:"defaultAllowIP"`
	DefaultAllowNonIP                 bool   `json:"defaultAllowNonIP"`
	DefaultInstallACLImplicitRules    bool   `json:"defaultInstallACLImplicitRules"`
	Description                       string `json:"description,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	PolicyState                       string `json:"policyState,omitempty"`
	Priority                          int    `json:"priority,omitempty"`
	PriorityType                      string `json:"priorityType,omitempty"`
	AssociatedLiveEntityID            string `json:"associatedLiveEntityID,omitempty"`
	AssociatedVirtualFirewallPolicyID string `json:"associatedVirtualFirewallPolicyID,omitempty"`
	AutoGeneratePriority              bool   `json:"autoGeneratePriority"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewEgressACLTemplate returns a new *EgressACLTemplate
func NewEgressACLTemplate() *EgressACLTemplate {

	return &EgressACLTemplate{}
}

// Identity returns the Identity of the object.
func (o *EgressACLTemplate) Identity() bambou.Identity {

	return EgressACLTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressACLTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressACLTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressACLTemplate from the server
func (o *EgressACLTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressACLTemplate into the server
func (o *EgressACLTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressACLTemplate from the server
func (o *EgressACLTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EgressACLTemplate
func (o *EgressACLTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressACLTemplate
func (o *EgressACLTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EgressACLEntryTemplates retrieves the list of child EgressACLEntryTemplates of the EgressACLTemplate
func (o *EgressACLTemplate) EgressACLEntryTemplates(info *bambou.FetchingInfo) (EgressACLEntryTemplatesList, *bambou.Error) {

	var list EgressACLEntryTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, EgressACLEntryTemplateIdentity, &list, info)
	return list, err
}

// CreateEgressACLEntryTemplate creates a new child EgressACLEntryTemplate under the EgressACLTemplate
func (o *EgressACLTemplate) CreateEgressACLEntryTemplate(child *EgressACLEntryTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressACLTemplate
func (o *EgressACLTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressACLTemplate
func (o *EgressACLTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VMs retrieves the list of child VMs of the EgressACLTemplate
func (o *EgressACLTemplate) VMs(info *bambou.FetchingInfo) (VMsList, *bambou.Error) {

	var list VMsList
	err := bambou.CurrentSession().FetchChildren(o, VMIdentity, &list, info)
	return list, err
}

// Jobs retrieves the list of child Jobs of the EgressACLTemplate
func (o *EgressACLTemplate) Jobs(info *bambou.FetchingInfo) (JobsList, *bambou.Error) {

	var list JobsList
	err := bambou.CurrentSession().FetchChildren(o, JobIdentity, &list, info)
	return list, err
}

// CreateJob creates a new child Job under the EgressACLTemplate
func (o *EgressACLTemplate) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Containers retrieves the list of child Containers of the EgressACLTemplate
func (o *EgressACLTemplate) Containers(info *bambou.FetchingInfo) (ContainersList, *bambou.Error) {

	var list ContainersList
	err := bambou.CurrentSession().FetchChildren(o, ContainerIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the EgressACLTemplate
func (o *EgressACLTemplate) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
