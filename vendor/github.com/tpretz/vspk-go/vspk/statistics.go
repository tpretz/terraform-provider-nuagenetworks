/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// StatisticsIdentity represents the Identity of the object
var StatisticsIdentity = bambou.Identity{
	Name:     "statistics",
	Category: "statistics",
}

// StatisticsList represents a list of Statistics
type StatisticsList []*Statistics

// StatisticsAncestor is the interface that an ancestor of a Statistics must implement.
// An Ancestor is defined as an entity that has Statistics as a descendant.
// An Ancestor can get a list of its child Statistics, but not necessarily create one.
type StatisticsAncestor interface {
	Statistics(*bambou.FetchingInfo) (StatisticsList, *bambou.Error)
}

// StatisticsParent is the interface that a parent of a Statistics must implement.
// A Parent is defined as an entity that has Statistics as a child.
// A Parent is an Ancestor which can create a Statistics.
type StatisticsParent interface {
	StatisticsAncestor
	CreateStatistics(*Statistics) *bambou.Error
}

// Statistics represents the model of a statistics
type Statistics struct {
	ID                 string      `json:"ID,omitempty"`
	ParentID           string      `json:"parentID,omitempty"`
	ParentType         string      `json:"parentType,omitempty"`
	Owner              string      `json:"owner,omitempty"`
	Version            int         `json:"version,omitempty"`
	EndTime            int         `json:"endTime,omitempty"`
	StartTime          int         `json:"startTime,omitempty"`
	StatsData          interface{} `json:"statsData,omitempty"`
	NumberOfDataPoints int         `json:"numberOfDataPoints,omitempty"`
}

// NewStatistics returns a new *Statistics
func NewStatistics() *Statistics {

	return &Statistics{}
}

// Identity returns the Identity of the object.
func (o *Statistics) Identity() bambou.Identity {

	return StatisticsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Statistics) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Statistics) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Statistics from the server
func (o *Statistics) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Statistics into the server
func (o *Statistics) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Statistics from the server
func (o *Statistics) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Statistics
func (o *Statistics) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Statistics
func (o *Statistics) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Statistics
func (o *Statistics) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Statistics
func (o *Statistics) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
