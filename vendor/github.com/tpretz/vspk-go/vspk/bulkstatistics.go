/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// BulkStatisticsIdentity represents the Identity of the object
var BulkStatisticsIdentity = bambou.Identity{
	Name:     "bulkstatistics",
	Category: "bulkstatistics",
}

// BulkStatisticsList represents a list of BulkStatistics
type BulkStatisticsList []*BulkStatistics

// BulkStatisticsAncestor is the interface that an ancestor of a BulkStatistics must implement.
// An Ancestor is defined as an entity that has BulkStatistics as a descendant.
// An Ancestor can get a list of its child BulkStatistics, but not necessarily create one.
type BulkStatisticsAncestor interface {
	BulkStatistics(*bambou.FetchingInfo) (BulkStatisticsList, *bambou.Error)
}

// BulkStatisticsParent is the interface that a parent of a BulkStatistics must implement.
// A Parent is defined as an entity that has BulkStatistics as a child.
// A Parent is an Ancestor which can create a BulkStatistics.
type BulkStatisticsParent interface {
	BulkStatisticsAncestor
	CreateBulkStatistics(*BulkStatistics) *bambou.Error
}

// BulkStatistics represents the model of a bulkstatistics
type BulkStatistics struct {
	ID                 string        `json:"ID,omitempty"`
	ParentID           string        `json:"parentID,omitempty"`
	ParentType         string        `json:"parentType,omitempty"`
	Owner              string        `json:"owner,omitempty"`
	Data               []interface{} `json:"data,omitempty"`
	Version            int           `json:"version,omitempty"`
	EndTime            int           `json:"endTime,omitempty"`
	StartTime          int           `json:"startTime,omitempty"`
	NumberOfDataPoints int           `json:"numberOfDataPoints,omitempty"`
}

// NewBulkStatistics returns a new *BulkStatistics
func NewBulkStatistics() *BulkStatistics {

	return &BulkStatistics{}
}

// Identity returns the Identity of the object.
func (o *BulkStatistics) Identity() bambou.Identity {

	return BulkStatisticsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BulkStatistics) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BulkStatistics) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BulkStatistics from the server
func (o *BulkStatistics) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BulkStatistics into the server
func (o *BulkStatistics) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BulkStatistics from the server
func (o *BulkStatistics) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BulkStatistics
func (o *BulkStatistics) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BulkStatistics
func (o *BulkStatistics) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BulkStatistics
func (o *BulkStatistics) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BulkStatistics
func (o *BulkStatistics) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
