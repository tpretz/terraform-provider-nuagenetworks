/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// DUCGroupBindingIdentity represents the Identity of the object
var DUCGroupBindingIdentity = bambou.Identity{
	Name:     "ducgroupbinding",
	Category: "ducgroupbindings",
}

// DUCGroupBindingsList represents a list of DUCGroupBindings
type DUCGroupBindingsList []*DUCGroupBinding

// DUCGroupBindingsAncestor is the interface that an ancestor of a DUCGroupBinding must implement.
// An Ancestor is defined as an entity that has DUCGroupBinding as a descendant.
// An Ancestor can get a list of its child DUCGroupBindings, but not necessarily create one.
type DUCGroupBindingsAncestor interface {
	DUCGroupBindings(*bambou.FetchingInfo) (DUCGroupBindingsList, *bambou.Error)
}

// DUCGroupBindingsParent is the interface that a parent of a DUCGroupBinding must implement.
// A Parent is defined as an entity that has DUCGroupBinding as a child.
// A Parent is an Ancestor which can create a DUCGroupBinding.
type DUCGroupBindingsParent interface {
	DUCGroupBindingsAncestor
	CreateDUCGroupBinding(*DUCGroupBinding) *bambou.Error
}

// DUCGroupBinding represents the model of a ducgroupbinding
type DUCGroupBinding struct {
	ID                   string `json:"ID,omitempty"`
	ParentID             string `json:"parentID,omitempty"`
	ParentType           string `json:"parentType,omitempty"`
	Owner                string `json:"owner,omitempty"`
	OneWayDelay          int    `json:"oneWayDelay,omitempty"`
	Priority             int    `json:"priority,omitempty"`
	AssociatedDUCGroupID string `json:"associatedDUCGroupID,omitempty"`
}

// NewDUCGroupBinding returns a new *DUCGroupBinding
func NewDUCGroupBinding() *DUCGroupBinding {

	return &DUCGroupBinding{
		OneWayDelay: 50,
	}
}

// Identity returns the Identity of the object.
func (o *DUCGroupBinding) Identity() bambou.Identity {

	return DUCGroupBindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DUCGroupBinding) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DUCGroupBinding) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DUCGroupBinding from the server
func (o *DUCGroupBinding) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DUCGroupBinding into the server
func (o *DUCGroupBinding) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DUCGroupBinding from the server
func (o *DUCGroupBinding) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
