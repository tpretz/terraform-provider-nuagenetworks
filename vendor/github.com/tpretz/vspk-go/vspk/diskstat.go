/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DiskStatIdentity represents the Identity of the object
var DiskStatIdentity = bambou.Identity{
	Name:     "diskstat",
	Category: "diskstats",
}

// DiskStatsList represents a list of DiskStats
type DiskStatsList []*DiskStat

// DiskStatsAncestor is the interface that an ancestor of a DiskStat must implement.
// An Ancestor is defined as an entity that has DiskStat as a descendant.
// An Ancestor can get a list of its child DiskStats, but not necessarily create one.
type DiskStatsAncestor interface {
	DiskStats(*bambou.FetchingInfo) (DiskStatsList, *bambou.Error)
}

// DiskStatsParent is the interface that a parent of a DiskStat must implement.
// A Parent is defined as an entity that has DiskStat as a child.
// A Parent is an Ancestor which can create a DiskStat.
type DiskStatsParent interface {
	DiskStatsAncestor
	CreateDiskStat(*DiskStat) *bambou.Error
}

// DiskStat represents the model of a diskstat
type DiskStat struct {
	ID          string  `json:"ID,omitempty"`
	ParentID    string  `json:"parentID,omitempty"`
	ParentType  string  `json:"parentType,omitempty"`
	Owner       string  `json:"owner,omitempty"`
	Name        string  `json:"name,omitempty"`
	Size        float64 `json:"size,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	EntityScope string  `json:"entityScope,omitempty"`
	Used        float64 `json:"used,omitempty"`
	Available   float64 `json:"available,omitempty"`
	ExternalID  string  `json:"externalID,omitempty"`
}

// NewDiskStat returns a new *DiskStat
func NewDiskStat() *DiskStat {

	return &DiskStat{}
}

// Identity returns the Identity of the object.
func (o *DiskStat) Identity() bambou.Identity {

	return DiskStatIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DiskStat) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DiskStat) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DiskStat from the server
func (o *DiskStat) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DiskStat into the server
func (o *DiskStat) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DiskStat from the server
func (o *DiskStat) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DiskStat
func (o *DiskStat) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DiskStat
func (o *DiskStat) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DiskStat
func (o *DiskStat) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DiskStat
func (o *DiskStat) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
