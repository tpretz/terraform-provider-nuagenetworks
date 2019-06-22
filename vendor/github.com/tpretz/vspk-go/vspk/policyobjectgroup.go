/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PolicyObjectGroupIdentity represents the Identity of the object
var PolicyObjectGroupIdentity = bambou.Identity{
	Name:     "policyobjectgroup",
	Category: "policyobjectgroups",
}

// PolicyObjectGroupsList represents a list of PolicyObjectGroups
type PolicyObjectGroupsList []*PolicyObjectGroup

// PolicyObjectGroupsAncestor is the interface that an ancestor of a PolicyObjectGroup must implement.
// An Ancestor is defined as an entity that has PolicyObjectGroup as a descendant.
// An Ancestor can get a list of its child PolicyObjectGroups, but not necessarily create one.
type PolicyObjectGroupsAncestor interface {
	PolicyObjectGroups(*bambou.FetchingInfo) (PolicyObjectGroupsList, *bambou.Error)
}

// PolicyObjectGroupsParent is the interface that a parent of a PolicyObjectGroup must implement.
// A Parent is defined as an entity that has PolicyObjectGroup as a child.
// A Parent is an Ancestor which can create a PolicyObjectGroup.
type PolicyObjectGroupsParent interface {
	PolicyObjectGroupsAncestor
	CreatePolicyObjectGroup(*PolicyObjectGroup) *bambou.Error
}

// PolicyObjectGroup represents the model of a policyobjectgroup
type PolicyObjectGroup struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

// NewPolicyObjectGroup returns a new *PolicyObjectGroup
func NewPolicyObjectGroup() *PolicyObjectGroup {

	return &PolicyObjectGroup{}
}

// Identity returns the Identity of the object.
func (o *PolicyObjectGroup) Identity() bambou.Identity {

	return PolicyObjectGroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyObjectGroup) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyObjectGroup) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyObjectGroup from the server
func (o *PolicyObjectGroup) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyObjectGroup into the server
func (o *PolicyObjectGroup) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyObjectGroup from the server
func (o *PolicyObjectGroup) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// NSGateways retrieves the list of child NSGateways of the PolicyObjectGroup
func (o *PolicyObjectGroup) NSGateways(info *bambou.FetchingInfo) (NSGatewaysList, *bambou.Error) {

	var list NSGatewaysList
	err := bambou.CurrentSession().FetchChildren(o, NSGatewayIdentity, &list, info)
	return list, err
}

// AssignNSGateways assigns the list of NSGateways to the PolicyObjectGroup
func (o *PolicyObjectGroup) AssignNSGateways(children NSGatewaysList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, NSGatewayIdentity)
}
