/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// COSRemarkingPolicyIdentity represents the Identity of the object
var COSRemarkingPolicyIdentity = bambou.Identity{
	Name:     "cosremarkingpolicy",
	Category: "cosremarkingpolicies",
}

// COSRemarkingPoliciesList represents a list of COSRemarkingPolicies
type COSRemarkingPoliciesList []*COSRemarkingPolicy

// COSRemarkingPoliciesAncestor is the interface that an ancestor of a COSRemarkingPolicy must implement.
// An Ancestor is defined as an entity that has COSRemarkingPolicy as a descendant.
// An Ancestor can get a list of its child COSRemarkingPolicies, but not necessarily create one.
type COSRemarkingPoliciesAncestor interface {
	COSRemarkingPolicies(*bambou.FetchingInfo) (COSRemarkingPoliciesList, *bambou.Error)
}

// COSRemarkingPoliciesParent is the interface that a parent of a COSRemarkingPolicy must implement.
// A Parent is defined as an entity that has COSRemarkingPolicy as a child.
// A Parent is an Ancestor which can create a COSRemarkingPolicy.
type COSRemarkingPoliciesParent interface {
	COSRemarkingPoliciesAncestor
	CreateCOSRemarkingPolicy(*COSRemarkingPolicy) *bambou.Error
}

// COSRemarkingPolicy represents the model of a cosremarkingpolicy
type COSRemarkingPolicy struct {
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

// NewCOSRemarkingPolicy returns a new *COSRemarkingPolicy
func NewCOSRemarkingPolicy() *COSRemarkingPolicy {

	return &COSRemarkingPolicy{}
}

// Identity returns the Identity of the object.
func (o *COSRemarkingPolicy) Identity() bambou.Identity {

	return COSRemarkingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *COSRemarkingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *COSRemarkingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the COSRemarkingPolicy from the server
func (o *COSRemarkingPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the COSRemarkingPolicy into the server
func (o *COSRemarkingPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the COSRemarkingPolicy from the server
func (o *COSRemarkingPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the COSRemarkingPolicy
func (o *COSRemarkingPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the COSRemarkingPolicy
func (o *COSRemarkingPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the COSRemarkingPolicy
func (o *COSRemarkingPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the COSRemarkingPolicy
func (o *COSRemarkingPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
