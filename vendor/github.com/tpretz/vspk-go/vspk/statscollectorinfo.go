/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// StatsCollectorInfoIdentity represents the Identity of the object
var StatsCollectorInfoIdentity = bambou.Identity{
	Name:     "statisticscollector",
	Category: "statisticscollector",
}

// StatsCollectorInfosList represents a list of StatsCollectorInfos
type StatsCollectorInfosList []*StatsCollectorInfo

// StatsCollectorInfosAncestor is the interface that an ancestor of a StatsCollectorInfo must implement.
// An Ancestor is defined as an entity that has StatsCollectorInfo as a descendant.
// An Ancestor can get a list of its child StatsCollectorInfos, but not necessarily create one.
type StatsCollectorInfosAncestor interface {
	StatsCollectorInfos(*bambou.FetchingInfo) (StatsCollectorInfosList, *bambou.Error)
}

// StatsCollectorInfosParent is the interface that a parent of a StatsCollectorInfo must implement.
// A Parent is defined as an entity that has StatsCollectorInfo as a child.
// A Parent is an Ancestor which can create a StatsCollectorInfo.
type StatsCollectorInfosParent interface {
	StatsCollectorInfosAncestor
	CreateStatsCollectorInfo(*StatsCollectorInfo) *bambou.Error
}

// StatsCollectorInfo represents the model of a statisticscollector
type StatsCollectorInfo struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	AddressType   string `json:"addressType,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	Port          string `json:"port,omitempty"`
	IpAddress     string `json:"ipAddress,omitempty"`
	ProtoBufPort  string `json:"protoBufPort,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewStatsCollectorInfo returns a new *StatsCollectorInfo
func NewStatsCollectorInfo() *StatsCollectorInfo {

	return &StatsCollectorInfo{}
}

// Identity returns the Identity of the object.
func (o *StatsCollectorInfo) Identity() bambou.Identity {

	return StatsCollectorInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StatsCollectorInfo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StatsCollectorInfo) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the StatsCollectorInfo from the server
func (o *StatsCollectorInfo) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the StatsCollectorInfo into the server
func (o *StatsCollectorInfo) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the StatsCollectorInfo from the server
func (o *StatsCollectorInfo) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the StatsCollectorInfo
func (o *StatsCollectorInfo) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the StatsCollectorInfo
func (o *StatsCollectorInfo) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the StatsCollectorInfo
func (o *StatsCollectorInfo) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the StatsCollectorInfo
func (o *StatsCollectorInfo) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
