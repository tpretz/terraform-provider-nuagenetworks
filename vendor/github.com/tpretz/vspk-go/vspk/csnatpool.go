/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CSNATPoolIdentity represents the Identity of the object
var CSNATPoolIdentity = bambou.Identity{
	Name:     "csnatpool",
	Category: "csnatpools",
}

// CSNATPoolsList represents a list of CSNATPools
type CSNATPoolsList []*CSNATPool

// CSNATPoolsAncestor is the interface that an ancestor of a CSNATPool must implement.
// An Ancestor is defined as an entity that has CSNATPool as a descendant.
// An Ancestor can get a list of its child CSNATPools, but not necessarily create one.
type CSNATPoolsAncestor interface {
	CSNATPools(*bambou.FetchingInfo) (CSNATPoolsList, *bambou.Error)
}

// CSNATPoolsParent is the interface that a parent of a CSNATPool must implement.
// A Parent is defined as an entity that has CSNATPool as a child.
// A Parent is an Ancestor which can create a CSNATPool.
type CSNATPoolsParent interface {
	CSNATPoolsAncestor
	CreateCSNATPool(*CSNATPool) *bambou.Error
}

// CSNATPool represents the model of a csnatpool
type CSNATPool struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	EndAddress    string `json:"endAddress,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	StartAddress  string `json:"startAddress,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewCSNATPool returns a new *CSNATPool
func NewCSNATPool() *CSNATPool {

	return &CSNATPool{}
}

// Identity returns the Identity of the object.
func (o *CSNATPool) Identity() bambou.Identity {

	return CSNATPoolIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CSNATPool) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CSNATPool) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CSNATPool from the server
func (o *CSNATPool) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CSNATPool into the server
func (o *CSNATPool) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CSNATPool from the server
func (o *CSNATPool) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the CSNATPool
func (o *CSNATPool) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the CSNATPool
func (o *CSNATPool) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the CSNATPool
func (o *CSNATPool) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the CSNATPool
func (o *CSNATPool) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CTranslationMaps retrieves the list of child CTranslationMaps of the CSNATPool
func (o *CSNATPool) CTranslationMaps(info *bambou.FetchingInfo) (CTranslationMapsList, *bambou.Error) {

	var list CTranslationMapsList
	err := bambou.CurrentSession().FetchChildren(o, CTranslationMapIdentity, &list, info)
	return list, err
}

// CreateCTranslationMap creates a new child CTranslationMap under the CSNATPool
func (o *CSNATPool) CreateCTranslationMap(child *CTranslationMap) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
