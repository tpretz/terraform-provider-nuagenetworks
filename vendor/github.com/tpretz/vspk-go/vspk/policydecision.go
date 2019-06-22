/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PolicyDecisionIdentity represents the Identity of the object
var PolicyDecisionIdentity = bambou.Identity{
	Name:     "policydecision",
	Category: "policydecisions",
}

// PolicyDecisionsList represents a list of PolicyDecisions
type PolicyDecisionsList []*PolicyDecision

// PolicyDecisionsAncestor is the interface that an ancestor of a PolicyDecision must implement.
// An Ancestor is defined as an entity that has PolicyDecision as a descendant.
// An Ancestor can get a list of its child PolicyDecisions, but not necessarily create one.
type PolicyDecisionsAncestor interface {
	PolicyDecisions(*bambou.FetchingInfo) (PolicyDecisionsList, *bambou.Error)
}

// PolicyDecisionsParent is the interface that a parent of a PolicyDecision must implement.
// A Parent is defined as an entity that has PolicyDecision as a child.
// A Parent is an Ancestor which can create a PolicyDecision.
type PolicyDecisionsParent interface {
	PolicyDecisionsAncestor
	CreatePolicyDecision(*PolicyDecision) *bambou.Error
}

// PolicyDecision represents the model of a policydecision
type PolicyDecision struct {
	ID                         string        `json:"ID,omitempty"`
	ParentID                   string        `json:"parentID,omitempty"`
	ParentType                 string        `json:"parentType,omitempty"`
	Owner                      string        `json:"owner,omitempty"`
	LastUpdatedBy              string        `json:"lastUpdatedBy,omitempty"`
	EgressACLs                 []interface{} `json:"egressACLs,omitempty"`
	EgressQos                  interface{}   `json:"egressQos,omitempty"`
	FipACLs                    []interface{} `json:"fipACLs,omitempty"`
	IngressACLs                []interface{} `json:"ingressACLs,omitempty"`
	IngressAdvFwd              []interface{} `json:"ingressAdvFwd,omitempty"`
	IngressExternalServiceACLs []interface{} `json:"ingressExternalServiceACLs,omitempty"`
	EntityScope                string        `json:"entityScope,omitempty"`
	Qos                        interface{}   `json:"qos,omitempty"`
	Stats                      interface{}   `json:"stats,omitempty"`
	ExternalID                 string        `json:"externalID,omitempty"`
}

// NewPolicyDecision returns a new *PolicyDecision
func NewPolicyDecision() *PolicyDecision {

	return &PolicyDecision{}
}

// Identity returns the Identity of the object.
func (o *PolicyDecision) Identity() bambou.Identity {

	return PolicyDecisionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyDecision) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyDecision) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyDecision from the server
func (o *PolicyDecision) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyDecision into the server
func (o *PolicyDecision) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyDecision from the server
func (o *PolicyDecision) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PolicyDecision
func (o *PolicyDecision) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyDecision
func (o *PolicyDecision) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyDecision
func (o *PolicyDecision) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyDecision
func (o *PolicyDecision) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// QOSs retrieves the list of child QOSs of the PolicyDecision
func (o *PolicyDecision) QOSs(info *bambou.FetchingInfo) (QOSsList, *bambou.Error) {

	var list QOSsList
	err := bambou.CurrentSession().FetchChildren(o, QOSIdentity, &list, info)
	return list, err
}
