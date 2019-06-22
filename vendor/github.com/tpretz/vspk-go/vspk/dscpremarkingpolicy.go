/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DSCPRemarkingPolicyIdentity represents the Identity of the object
var DSCPRemarkingPolicyIdentity = bambou.Identity{
	Name:     "dscpremarkingpolicy",
	Category: "dscpremarkingpolicies",
}

// DSCPRemarkingPoliciesList represents a list of DSCPRemarkingPolicies
type DSCPRemarkingPoliciesList []*DSCPRemarkingPolicy

// DSCPRemarkingPoliciesAncestor is the interface that an ancestor of a DSCPRemarkingPolicy must implement.
// An Ancestor is defined as an entity that has DSCPRemarkingPolicy as a descendant.
// An Ancestor can get a list of its child DSCPRemarkingPolicies, but not necessarily create one.
type DSCPRemarkingPoliciesAncestor interface {
	DSCPRemarkingPolicies(*bambou.FetchingInfo) (DSCPRemarkingPoliciesList, *bambou.Error)
}

// DSCPRemarkingPoliciesParent is the interface that a parent of a DSCPRemarkingPolicy must implement.
// A Parent is defined as an entity that has DSCPRemarkingPolicy as a child.
// A Parent is an Ancestor which can create a DSCPRemarkingPolicy.
type DSCPRemarkingPoliciesParent interface {
	DSCPRemarkingPoliciesAncestor
	CreateDSCPRemarkingPolicy(*DSCPRemarkingPolicy) *bambou.Error
}

// DSCPRemarkingPolicy represents the model of a dscpremarkingpolicy
type DSCPRemarkingPolicy struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	DSCP            string `json:"DSCP,omitempty"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	ForwardingClass string `json:"forwardingClass,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewDSCPRemarkingPolicy returns a new *DSCPRemarkingPolicy
func NewDSCPRemarkingPolicy() *DSCPRemarkingPolicy {

	return &DSCPRemarkingPolicy{}
}

// Identity returns the Identity of the object.
func (o *DSCPRemarkingPolicy) Identity() bambou.Identity {

	return DSCPRemarkingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DSCPRemarkingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DSCPRemarkingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DSCPRemarkingPolicy from the server
func (o *DSCPRemarkingPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DSCPRemarkingPolicy into the server
func (o *DSCPRemarkingPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DSCPRemarkingPolicy from the server
func (o *DSCPRemarkingPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DSCPRemarkingPolicy
func (o *DSCPRemarkingPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DSCPRemarkingPolicy
func (o *DSCPRemarkingPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DSCPRemarkingPolicy
func (o *DSCPRemarkingPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DSCPRemarkingPolicy
func (o *DSCPRemarkingPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
