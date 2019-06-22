/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DSCPForwardingClassMappingIdentity represents the Identity of the object
var DSCPForwardingClassMappingIdentity = bambou.Identity{
	Name:     "dscpforwardingclassmapping",
	Category: "dscpforwardingclassmappings",
}

// DSCPForwardingClassMappingsList represents a list of DSCPForwardingClassMappings
type DSCPForwardingClassMappingsList []*DSCPForwardingClassMapping

// DSCPForwardingClassMappingsAncestor is the interface that an ancestor of a DSCPForwardingClassMapping must implement.
// An Ancestor is defined as an entity that has DSCPForwardingClassMapping as a descendant.
// An Ancestor can get a list of its child DSCPForwardingClassMappings, but not necessarily create one.
type DSCPForwardingClassMappingsAncestor interface {
	DSCPForwardingClassMappings(*bambou.FetchingInfo) (DSCPForwardingClassMappingsList, *bambou.Error)
}

// DSCPForwardingClassMappingsParent is the interface that a parent of a DSCPForwardingClassMapping must implement.
// A Parent is defined as an entity that has DSCPForwardingClassMapping as a child.
// A Parent is an Ancestor which can create a DSCPForwardingClassMapping.
type DSCPForwardingClassMappingsParent interface {
	DSCPForwardingClassMappingsAncestor
	CreateDSCPForwardingClassMapping(*DSCPForwardingClassMapping) *bambou.Error
}

// DSCPForwardingClassMapping represents the model of a dscpforwardingclassmapping
type DSCPForwardingClassMapping struct {
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

// NewDSCPForwardingClassMapping returns a new *DSCPForwardingClassMapping
func NewDSCPForwardingClassMapping() *DSCPForwardingClassMapping {

	return &DSCPForwardingClassMapping{}
}

// Identity returns the Identity of the object.
func (o *DSCPForwardingClassMapping) Identity() bambou.Identity {

	return DSCPForwardingClassMappingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DSCPForwardingClassMapping) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DSCPForwardingClassMapping) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DSCPForwardingClassMapping from the server
func (o *DSCPForwardingClassMapping) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DSCPForwardingClassMapping into the server
func (o *DSCPForwardingClassMapping) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DSCPForwardingClassMapping from the server
func (o *DSCPForwardingClassMapping) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DSCPForwardingClassMapping
func (o *DSCPForwardingClassMapping) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DSCPForwardingClassMapping
func (o *DSCPForwardingClassMapping) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DSCPForwardingClassMapping
func (o *DSCPForwardingClassMapping) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DSCPForwardingClassMapping
func (o *DSCPForwardingClassMapping) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
