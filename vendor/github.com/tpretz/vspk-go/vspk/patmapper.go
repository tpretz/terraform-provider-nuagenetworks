/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PATMapperIdentity represents the Identity of the object
var PATMapperIdentity = bambou.Identity{
	Name:     "patmapper",
	Category: "patmappers",
}

// PATMappersList represents a list of PATMappers
type PATMappersList []*PATMapper

// PATMappersAncestor is the interface that an ancestor of a PATMapper must implement.
// An Ancestor is defined as an entity that has PATMapper as a descendant.
// An Ancestor can get a list of its child PATMappers, but not necessarily create one.
type PATMappersAncestor interface {
	PATMappers(*bambou.FetchingInfo) (PATMappersList, *bambou.Error)
}

// PATMappersParent is the interface that a parent of a PATMapper must implement.
// A Parent is defined as an entity that has PATMapper as a child.
// A Parent is an Ancestor which can create a PATMapper.
type PATMappersParent interface {
	PATMappersAncestor
	CreatePATMapper(*PATMapper) *bambou.Error
}

// PATMapper represents the model of a patmapper
type PATMapper struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPATMapper returns a new *PATMapper
func NewPATMapper() *PATMapper {

	return &PATMapper{}
}

// Identity returns the Identity of the object.
func (o *PATMapper) Identity() bambou.Identity {

	return PATMapperIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PATMapper) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PATMapper) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PATMapper from the server
func (o *PATMapper) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PATMapper into the server
func (o *PATMapper) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PATMapper from the server
func (o *PATMapper) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// SharedNetworkResources retrieves the list of child SharedNetworkResources of the PATMapper
func (o *PATMapper) SharedNetworkResources(info *bambou.FetchingInfo) (SharedNetworkResourcesList, *bambou.Error) {

	var list SharedNetworkResourcesList
	err := bambou.CurrentSession().FetchChildren(o, SharedNetworkResourceIdentity, &list, info)
	return list, err
}

// Subnets retrieves the list of child Subnets of the PATMapper
func (o *PATMapper) Subnets(info *bambou.FetchingInfo) (SubnetsList, *bambou.Error) {

	var list SubnetsList
	err := bambou.CurrentSession().FetchChildren(o, SubnetIdentity, &list, info)
	return list, err
}
