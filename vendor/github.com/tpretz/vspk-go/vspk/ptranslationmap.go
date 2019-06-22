/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PTranslationMapIdentity represents the Identity of the object
var PTranslationMapIdentity = bambou.Identity{
	Name:     "ptranslationmap",
	Category: "ptranslationmaps",
}

// PTranslationMapsList represents a list of PTranslationMaps
type PTranslationMapsList []*PTranslationMap

// PTranslationMapsAncestor is the interface that an ancestor of a PTranslationMap must implement.
// An Ancestor is defined as an entity that has PTranslationMap as a descendant.
// An Ancestor can get a list of its child PTranslationMaps, but not necessarily create one.
type PTranslationMapsAncestor interface {
	PTranslationMaps(*bambou.FetchingInfo) (PTranslationMapsList, *bambou.Error)
}

// PTranslationMapsParent is the interface that a parent of a PTranslationMap must implement.
// A Parent is defined as an entity that has PTranslationMap as a child.
// A Parent is an Ancestor which can create a PTranslationMap.
type PTranslationMapsParent interface {
	PTranslationMapsAncestor
	CreatePTranslationMap(*PTranslationMap) *bambou.Error
}

// PTranslationMap represents the model of a ptranslationmap
type PTranslationMap struct {
	ID              string        `json:"ID,omitempty"`
	ParentID        string        `json:"parentID,omitempty"`
	ParentType      string        `json:"parentType,omitempty"`
	Owner           string        `json:"owner,omitempty"`
	SPATSourceList  []interface{} `json:"SPATSourceList,omitempty"`
	MappingType     string        `json:"mappingType,omitempty"`
	ProviderAliasIP string        `json:"providerAliasIP,omitempty"`
	ProviderIP      string        `json:"providerIP,omitempty"`
}

// NewPTranslationMap returns a new *PTranslationMap
func NewPTranslationMap() *PTranslationMap {

	return &PTranslationMap{}
}

// Identity returns the Identity of the object.
func (o *PTranslationMap) Identity() bambou.Identity {

	return PTranslationMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PTranslationMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PTranslationMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PTranslationMap from the server
func (o *PTranslationMap) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PTranslationMap into the server
func (o *PTranslationMap) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PTranslationMap from the server
func (o *PTranslationMap) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
