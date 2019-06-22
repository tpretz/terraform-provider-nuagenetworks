/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EgressQOSPolicyIdentity represents the Identity of the object
var EgressQOSPolicyIdentity = bambou.Identity{
	Name:     "egressqospolicy",
	Category: "egressqospolicies",
}

// EgressQOSPoliciesList represents a list of EgressQOSPolicies
type EgressQOSPoliciesList []*EgressQOSPolicy

// EgressQOSPoliciesAncestor is the interface that an ancestor of a EgressQOSPolicy must implement.
// An Ancestor is defined as an entity that has EgressQOSPolicy as a descendant.
// An Ancestor can get a list of its child EgressQOSPolicies, but not necessarily create one.
type EgressQOSPoliciesAncestor interface {
	EgressQOSPolicies(*bambou.FetchingInfo) (EgressQOSPoliciesList, *bambou.Error)
}

// EgressQOSPoliciesParent is the interface that a parent of a EgressQOSPolicy must implement.
// A Parent is defined as an entity that has EgressQOSPolicy as a child.
// A Parent is an Ancestor which can create a EgressQOSPolicy.
type EgressQOSPoliciesParent interface {
	EgressQOSPoliciesAncestor
	CreateEgressQOSPolicy(*EgressQOSPolicy) *bambou.Error
}

// EgressQOSPolicy represents the model of a egressqospolicy
type EgressQOSPolicy struct {
	ID                                   string        `json:"ID,omitempty"`
	ParentID                             string        `json:"parentID,omitempty"`
	ParentType                           string        `json:"parentType,omitempty"`
	Owner                                string        `json:"owner,omitempty"`
	Name                                 string        `json:"name,omitempty"`
	ParentQueueAssociatedRateLimiterID   string        `json:"parentQueueAssociatedRateLimiterID,omitempty"`
	LastUpdatedBy                        string        `json:"lastUpdatedBy,omitempty"`
	DefaultServiceClass                  string        `json:"defaultServiceClass,omitempty"`
	Description                          string        `json:"description,omitempty"`
	EntityScope                          string        `json:"entityScope,omitempty"`
	AssocEgressQosId                     string        `json:"assocEgressQosId,omitempty"`
	AssociatedCOSRemarkingPolicyTableID  string        `json:"associatedCOSRemarkingPolicyTableID,omitempty"`
	AssociatedDSCPRemarkingPolicyTableID string        `json:"associatedDSCPRemarkingPolicyTableID,omitempty"`
	Queue1AssociatedRateLimiterID        string        `json:"queue1AssociatedRateLimiterID,omitempty"`
	Queue1ForwardingClasses              []interface{} `json:"queue1ForwardingClasses,omitempty"`
	Queue2AssociatedRateLimiterID        string        `json:"queue2AssociatedRateLimiterID,omitempty"`
	Queue2ForwardingClasses              []interface{} `json:"queue2ForwardingClasses,omitempty"`
	Queue3AssociatedRateLimiterID        string        `json:"queue3AssociatedRateLimiterID,omitempty"`
	Queue3ForwardingClasses              []interface{} `json:"queue3ForwardingClasses,omitempty"`
	Queue4AssociatedRateLimiterID        string        `json:"queue4AssociatedRateLimiterID,omitempty"`
	Queue4ForwardingClasses              []interface{} `json:"queue4ForwardingClasses,omitempty"`
	ExternalID                           string        `json:"externalID,omitempty"`
}

// NewEgressQOSPolicy returns a new *EgressQOSPolicy
func NewEgressQOSPolicy() *EgressQOSPolicy {

	return &EgressQOSPolicy{}
}

// Identity returns the Identity of the object.
func (o *EgressQOSPolicy) Identity() bambou.Identity {

	return EgressQOSPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressQOSPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressQOSPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressQOSPolicy from the server
func (o *EgressQOSPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressQOSPolicy into the server
func (o *EgressQOSPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressQOSPolicy from the server
func (o *EgressQOSPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EgressQOSPolicy
func (o *EgressQOSPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressQOSPolicy
func (o *EgressQOSPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressQOSPolicy
func (o *EgressQOSPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressQOSPolicy
func (o *EgressQOSPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
