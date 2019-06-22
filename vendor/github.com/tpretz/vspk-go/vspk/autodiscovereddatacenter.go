/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AutodiscovereddatacenterIdentity represents the Identity of the object
var AutodiscovereddatacenterIdentity = bambou.Identity{
	Name:     "autodiscovereddatacenter",
	Category: "autodiscovereddatacenters",
}

// AutodiscovereddatacentersList represents a list of Autodiscovereddatacenters
type AutodiscovereddatacentersList []*Autodiscovereddatacenter

// AutodiscovereddatacentersAncestor is the interface that an ancestor of a Autodiscovereddatacenter must implement.
// An Ancestor is defined as an entity that has Autodiscovereddatacenter as a descendant.
// An Ancestor can get a list of its child Autodiscovereddatacenters, but not necessarily create one.
type AutodiscovereddatacentersAncestor interface {
	Autodiscovereddatacenters(*bambou.FetchingInfo) (AutodiscovereddatacentersList, *bambou.Error)
}

// AutodiscovereddatacentersParent is the interface that a parent of a Autodiscovereddatacenter must implement.
// A Parent is defined as an entity that has Autodiscovereddatacenter as a child.
// A Parent is an Ancestor which can create a Autodiscovereddatacenter.
type AutodiscovereddatacentersParent interface {
	AutodiscovereddatacentersAncestor
	CreateAutodiscovereddatacenter(*Autodiscovereddatacenter) *bambou.Error
}

// Autodiscovereddatacenter represents the model of a autodiscovereddatacenter
type Autodiscovereddatacenter struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	Name                string `json:"name,omitempty"`
	ManagedObjectID     string `json:"managedObjectID,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	AssociatedVCenterID string `json:"associatedVCenterID,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

// NewAutodiscovereddatacenter returns a new *Autodiscovereddatacenter
func NewAutodiscovereddatacenter() *Autodiscovereddatacenter {

	return &Autodiscovereddatacenter{}
}

// Identity returns the Identity of the object.
func (o *Autodiscovereddatacenter) Identity() bambou.Identity {

	return AutodiscovereddatacenterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Autodiscovereddatacenter) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Autodiscovereddatacenter) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Autodiscovereddatacenter from the server
func (o *Autodiscovereddatacenter) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Autodiscovereddatacenter into the server
func (o *Autodiscovereddatacenter) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Autodiscovereddatacenter from the server
func (o *Autodiscovereddatacenter) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
