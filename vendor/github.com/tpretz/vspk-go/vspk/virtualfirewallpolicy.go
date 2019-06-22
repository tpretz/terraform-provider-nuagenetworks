/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VirtualFirewallPolicyIdentity represents the Identity of the object
var VirtualFirewallPolicyIdentity = bambou.Identity{
	Name:     "virtualfirewallpolicy",
	Category: "virtualfirewallpolicies",
}

// VirtualFirewallPoliciesList represents a list of VirtualFirewallPolicies
type VirtualFirewallPoliciesList []*VirtualFirewallPolicy

// VirtualFirewallPoliciesAncestor is the interface that an ancestor of a VirtualFirewallPolicy must implement.
// An Ancestor is defined as an entity that has VirtualFirewallPolicy as a descendant.
// An Ancestor can get a list of its child VirtualFirewallPolicies, but not necessarily create one.
type VirtualFirewallPoliciesAncestor interface {
	VirtualFirewallPolicies(*bambou.FetchingInfo) (VirtualFirewallPoliciesList, *bambou.Error)
}

// VirtualFirewallPoliciesParent is the interface that a parent of a VirtualFirewallPolicy must implement.
// A Parent is defined as an entity that has VirtualFirewallPolicy as a child.
// A Parent is an Ancestor which can create a VirtualFirewallPolicy.
type VirtualFirewallPoliciesParent interface {
	VirtualFirewallPoliciesAncestor
	CreateVirtualFirewallPolicy(*VirtualFirewallPolicy) *bambou.Error
}

// VirtualFirewallPolicy represents the model of a virtualfirewallpolicy
type VirtualFirewallPolicy struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	Name                           string `json:"name,omitempty"`
	LastUpdatedBy                  string `json:"lastUpdatedBy,omitempty"`
	Active                         bool   `json:"active"`
	DefaultAllowIP                 bool   `json:"defaultAllowIP"`
	DefaultAllowNonIP              bool   `json:"defaultAllowNonIP"`
	DefaultInstallACLImplicitRules bool   `json:"defaultInstallACLImplicitRules"`
	Description                    string `json:"description,omitempty"`
	AllowAddressSpoof              bool   `json:"allowAddressSpoof"`
	EntityScope                    string `json:"entityScope,omitempty"`
	PolicyState                    string `json:"policyState,omitempty"`
	Priority                       int    `json:"priority,omitempty"`
	PriorityType                   string `json:"priorityType,omitempty"`
	AssociatedEgressTemplateID     string `json:"associatedEgressTemplateID,omitempty"`
	AssociatedIngressTemplateID    string `json:"associatedIngressTemplateID,omitempty"`
	AssociatedLiveEntityID         string `json:"associatedLiveEntityID,omitempty"`
	AutoGeneratePriority           bool   `json:"autoGeneratePriority"`
	ExternalID                     string `json:"externalID,omitempty"`
}

// NewVirtualFirewallPolicy returns a new *VirtualFirewallPolicy
func NewVirtualFirewallPolicy() *VirtualFirewallPolicy {

	return &VirtualFirewallPolicy{
		Active:                         false,
		DefaultAllowIP:                 false,
		DefaultAllowNonIP:              false,
		DefaultInstallACLImplicitRules: false,
		AllowAddressSpoof:              false,
		AutoGeneratePriority:           false,
	}
}

// Identity returns the Identity of the object.
func (o *VirtualFirewallPolicy) Identity() bambou.Identity {

	return VirtualFirewallPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VirtualFirewallPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VirtualFirewallPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VirtualFirewallPolicy from the server
func (o *VirtualFirewallPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VirtualFirewallPolicy into the server
func (o *VirtualFirewallPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VirtualFirewallPolicy from the server
func (o *VirtualFirewallPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VirtualFirewallPolicy
func (o *VirtualFirewallPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VirtualFirewallPolicy
func (o *VirtualFirewallPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VirtualFirewallRules retrieves the list of child VirtualFirewallRules of the VirtualFirewallPolicy
func (o *VirtualFirewallPolicy) VirtualFirewallRules(info *bambou.FetchingInfo) (VirtualFirewallRulesList, *bambou.Error) {

	var list VirtualFirewallRulesList
	err := bambou.CurrentSession().FetchChildren(o, VirtualFirewallRuleIdentity, &list, info)
	return list, err
}

// CreateVirtualFirewallRule creates a new child VirtualFirewallRule under the VirtualFirewallPolicy
func (o *VirtualFirewallPolicy) CreateVirtualFirewallRule(child *VirtualFirewallRule) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VirtualFirewallPolicy
func (o *VirtualFirewallPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VirtualFirewallPolicy
func (o *VirtualFirewallPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
