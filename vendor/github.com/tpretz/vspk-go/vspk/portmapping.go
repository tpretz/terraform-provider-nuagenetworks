/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PortMappingIdentity represents the Identity of the object
var PortMappingIdentity = bambou.Identity{
	Name:     "portmapping",
	Category: "portmappings",
}

// PortMappingsList represents a list of PortMappings
type PortMappingsList []*PortMapping

// PortMappingsAncestor is the interface that an ancestor of a PortMapping must implement.
// An Ancestor is defined as an entity that has PortMapping as a descendant.
// An Ancestor can get a list of its child PortMappings, but not necessarily create one.
type PortMappingsAncestor interface {
	PortMappings(*bambou.FetchingInfo) (PortMappingsList, *bambou.Error)
}

// PortMappingsParent is the interface that a parent of a PortMapping must implement.
// A Parent is defined as an entity that has PortMapping as a child.
// A Parent is an Ancestor which can create a PortMapping.
type PortMappingsParent interface {
	PortMappingsAncestor
	CreatePortMapping(*PortMapping) *bambou.Error
}

// PortMapping represents the model of a portmapping
type PortMapping struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	PrivatePort   string `json:"privatePort,omitempty"`
	PublicPort    string `json:"publicPort,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPortMapping returns a new *PortMapping
func NewPortMapping() *PortMapping {

	return &PortMapping{}
}

// Identity returns the Identity of the object.
func (o *PortMapping) Identity() bambou.Identity {

	return PortMappingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PortMapping) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PortMapping) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PortMapping from the server
func (o *PortMapping) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PortMapping into the server
func (o *PortMapping) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PortMapping from the server
func (o *PortMapping) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
