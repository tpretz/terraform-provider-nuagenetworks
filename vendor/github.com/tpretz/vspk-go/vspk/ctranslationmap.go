/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// CTranslationMapIdentity represents the Identity of the object
var CTranslationMapIdentity = bambou.Identity{
	Name:     "ctranslationmap",
	Category: "ctranslationmaps",
}

// CTranslationMapsList represents a list of CTranslationMaps
type CTranslationMapsList []*CTranslationMap

// CTranslationMapsAncestor is the interface that an ancestor of a CTranslationMap must implement.
// An Ancestor is defined as an entity that has CTranslationMap as a descendant.
// An Ancestor can get a list of its child CTranslationMaps, but not necessarily create one.
type CTranslationMapsAncestor interface {
	CTranslationMaps(*bambou.FetchingInfo) (CTranslationMapsList, *bambou.Error)
}

// CTranslationMapsParent is the interface that a parent of a CTranslationMap must implement.
// A Parent is defined as an entity that has CTranslationMap as a child.
// A Parent is an Ancestor which can create a CTranslationMap.
type CTranslationMapsParent interface {
	CTranslationMapsAncestor
	CreateCTranslationMap(*CTranslationMap) *bambou.Error
}

// CTranslationMap represents the model of a ctranslationmap
type CTranslationMap struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	MappingType     string `json:"mappingType,omitempty"`
	CustomerAliasIP string `json:"customerAliasIP,omitempty"`
	CustomerIP      string `json:"customerIP,omitempty"`
}

// NewCTranslationMap returns a new *CTranslationMap
func NewCTranslationMap() *CTranslationMap {

	return &CTranslationMap{}
}

// Identity returns the Identity of the object.
func (o *CTranslationMap) Identity() bambou.Identity {

	return CTranslationMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CTranslationMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CTranslationMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CTranslationMap from the server
func (o *CTranslationMap) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CTranslationMap into the server
func (o *CTranslationMap) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CTranslationMap from the server
func (o *CTranslationMap) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
