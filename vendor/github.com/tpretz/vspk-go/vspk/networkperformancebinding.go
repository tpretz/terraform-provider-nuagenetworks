/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NetworkPerformanceBindingIdentity represents the Identity of the object
var NetworkPerformanceBindingIdentity = bambou.Identity{
	Name:     "networkperformancebinding",
	Category: "networkperformancebindings",
}

// NetworkPerformanceBindingsList represents a list of NetworkPerformanceBindings
type NetworkPerformanceBindingsList []*NetworkPerformanceBinding

// NetworkPerformanceBindingsAncestor is the interface that an ancestor of a NetworkPerformanceBinding must implement.
// An Ancestor is defined as an entity that has NetworkPerformanceBinding as a descendant.
// An Ancestor can get a list of its child NetworkPerformanceBindings, but not necessarily create one.
type NetworkPerformanceBindingsAncestor interface {
	NetworkPerformanceBindings(*bambou.FetchingInfo) (NetworkPerformanceBindingsList, *bambou.Error)
}

// NetworkPerformanceBindingsParent is the interface that a parent of a NetworkPerformanceBinding must implement.
// A Parent is defined as an entity that has NetworkPerformanceBinding as a child.
// A Parent is an Ancestor which can create a NetworkPerformanceBinding.
type NetworkPerformanceBindingsParent interface {
	NetworkPerformanceBindingsAncestor
	CreateNetworkPerformanceBinding(*NetworkPerformanceBinding) *bambou.Error
}

// NetworkPerformanceBinding represents the model of a networkperformancebinding
type NetworkPerformanceBinding struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	ReadOnly                       bool   `json:"readOnly"`
	Priority                       int    `json:"priority,omitempty"`
	AssociatedNetworkMeasurementID string `json:"associatedNetworkMeasurementID,omitempty"`
}

// NewNetworkPerformanceBinding returns a new *NetworkPerformanceBinding
func NewNetworkPerformanceBinding() *NetworkPerformanceBinding {

	return &NetworkPerformanceBinding{
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *NetworkPerformanceBinding) Identity() bambou.Identity {

	return NetworkPerformanceBindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkPerformanceBinding) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkPerformanceBinding) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetworkPerformanceBinding from the server
func (o *NetworkPerformanceBinding) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetworkPerformanceBinding into the server
func (o *NetworkPerformanceBinding) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetworkPerformanceBinding from the server
func (o *NetworkPerformanceBinding) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
