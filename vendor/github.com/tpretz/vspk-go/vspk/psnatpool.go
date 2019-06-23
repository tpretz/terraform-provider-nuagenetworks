/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PSNATPoolIdentity represents the Identity of the object
var PSNATPoolIdentity = bambou.Identity{
	Name:     "psnatpool",
	Category: "psnatpools",
}

// PSNATPoolsList represents a list of PSNATPools
type PSNATPoolsList []*PSNATPool

// PSNATPoolsAncestor is the interface that an ancestor of a PSNATPool must implement.
// An Ancestor is defined as an entity that has PSNATPool as a descendant.
// An Ancestor can get a list of its child PSNATPools, but not necessarily create one.
type PSNATPoolsAncestor interface {
	PSNATPools(*bambou.FetchingInfo) (PSNATPoolsList, *bambou.Error)
}

// PSNATPoolsParent is the interface that a parent of a PSNATPool must implement.
// A Parent is defined as an entity that has PSNATPool as a child.
// A Parent is an Ancestor which can create a PSNATPool.
type PSNATPoolsParent interface {
	PSNATPoolsAncestor
	CreatePSNATPool(*PSNATPool) *bambou.Error
}

// PSNATPool represents the model of a psnatpool
type PSNATPool struct {
	ID           string `json:"ID,omitempty"`
	ParentID     string `json:"parentID,omitempty"`
	ParentType   string `json:"parentType,omitempty"`
	Owner        string `json:"owner,omitempty"`
	Name         string `json:"name,omitempty"`
	EndAddress   string `json:"endAddress,omitempty"`
	StartAddress string `json:"startAddress,omitempty"`
}

// NewPSNATPool returns a new *PSNATPool
func NewPSNATPool() *PSNATPool {

	return &PSNATPool{}
}

// Identity returns the Identity of the object.
func (o *PSNATPool) Identity() bambou.Identity {

	return PSNATPoolIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PSNATPool) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PSNATPool) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PSNATPool from the server
func (o *PSNATPool) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PSNATPool into the server
func (o *PSNATPool) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PSNATPool from the server
func (o *PSNATPool) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// PSPATMaps retrieves the list of child PSPATMaps of the PSNATPool
func (o *PSNATPool) PSPATMaps(info *bambou.FetchingInfo) (PSPATMapsList, *bambou.Error) {

	var list PSPATMapsList
	err := bambou.CurrentSession().FetchChildren(o, PSPATMapIdentity, &list, info)
	return list, err
}

// CreatePSPATMap creates a new child PSPATMap under the PSNATPool
func (o *PSNATPool) CreatePSPATMap(child *PSPATMap) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PTranslationMaps retrieves the list of child PTranslationMaps of the PSNATPool
func (o *PSNATPool) PTranslationMaps(info *bambou.FetchingInfo) (PTranslationMapsList, *bambou.Error) {

	var list PTranslationMapsList
	err := bambou.CurrentSession().FetchChildren(o, PTranslationMapIdentity, &list, info)
	return list, err
}

// CreatePTranslationMap creates a new child PTranslationMap under the PSNATPool
func (o *PSNATPool) CreatePTranslationMap(child *PTranslationMap) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
