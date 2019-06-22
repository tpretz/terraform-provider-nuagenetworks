/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// StatisticsPolicyIdentity represents the Identity of the object
var StatisticsPolicyIdentity = bambou.Identity{
	Name:     "statisticspolicy",
	Category: "statisticspolicies",
}

// StatisticsPoliciesList represents a list of StatisticsPolicies
type StatisticsPoliciesList []*StatisticsPolicy

// StatisticsPoliciesAncestor is the interface that an ancestor of a StatisticsPolicy must implement.
// An Ancestor is defined as an entity that has StatisticsPolicy as a descendant.
// An Ancestor can get a list of its child StatisticsPolicies, but not necessarily create one.
type StatisticsPoliciesAncestor interface {
	StatisticsPolicies(*bambou.FetchingInfo) (StatisticsPoliciesList, *bambou.Error)
}

// StatisticsPoliciesParent is the interface that a parent of a StatisticsPolicy must implement.
// A Parent is defined as an entity that has StatisticsPolicy as a child.
// A Parent is an Ancestor which can create a StatisticsPolicy.
type StatisticsPoliciesParent interface {
	StatisticsPoliciesAncestor
	CreateStatisticsPolicy(*StatisticsPolicy) *bambou.Error
}

// StatisticsPolicy represents the model of a statisticspolicy
type StatisticsPolicy struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	Name                    string `json:"name,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	DataCollectionFrequency int    `json:"dataCollectionFrequency,omitempty"`
	Description             string `json:"description,omitempty"`
	EntityScope             string `json:"entityScope,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
}

// NewStatisticsPolicy returns a new *StatisticsPolicy
func NewStatisticsPolicy() *StatisticsPolicy {

	return &StatisticsPolicy{}
}

// Identity returns the Identity of the object.
func (o *StatisticsPolicy) Identity() bambou.Identity {

	return StatisticsPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StatisticsPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StatisticsPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the StatisticsPolicy from the server
func (o *StatisticsPolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the StatisticsPolicy into the server
func (o *StatisticsPolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the StatisticsPolicy from the server
func (o *StatisticsPolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the StatisticsPolicy
func (o *StatisticsPolicy) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the StatisticsPolicy
func (o *StatisticsPolicy) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the StatisticsPolicy
func (o *StatisticsPolicy) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the StatisticsPolicy
func (o *StatisticsPolicy) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
