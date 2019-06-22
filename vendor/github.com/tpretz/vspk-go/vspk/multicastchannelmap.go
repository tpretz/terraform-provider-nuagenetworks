/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MultiCastChannelMapIdentity represents the Identity of the object
var MultiCastChannelMapIdentity = bambou.Identity{
	Name:     "multicastchannelmap",
	Category: "multicastchannelmaps",
}

// MultiCastChannelMapsList represents a list of MultiCastChannelMaps
type MultiCastChannelMapsList []*MultiCastChannelMap

// MultiCastChannelMapsAncestor is the interface that an ancestor of a MultiCastChannelMap must implement.
// An Ancestor is defined as an entity that has MultiCastChannelMap as a descendant.
// An Ancestor can get a list of its child MultiCastChannelMaps, but not necessarily create one.
type MultiCastChannelMapsAncestor interface {
	MultiCastChannelMaps(*bambou.FetchingInfo) (MultiCastChannelMapsList, *bambou.Error)
}

// MultiCastChannelMapsParent is the interface that a parent of a MultiCastChannelMap must implement.
// A Parent is defined as an entity that has MultiCastChannelMap as a child.
// A Parent is an Ancestor which can create a MultiCastChannelMap.
type MultiCastChannelMapsParent interface {
	MultiCastChannelMapsAncestor
	CreateMultiCastChannelMap(*MultiCastChannelMap) *bambou.Error
}

// MultiCastChannelMap represents the model of a multicastchannelmap
type MultiCastChannelMap struct {
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

// NewMultiCastChannelMap returns a new *MultiCastChannelMap
func NewMultiCastChannelMap() *MultiCastChannelMap {

	return &MultiCastChannelMap{}
}

// Identity returns the Identity of the object.
func (o *MultiCastChannelMap) Identity() bambou.Identity {

	return MultiCastChannelMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MultiCastChannelMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MultiCastChannelMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MultiCastChannelMap from the server
func (o *MultiCastChannelMap) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MultiCastChannelMap into the server
func (o *MultiCastChannelMap) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MultiCastChannelMap from the server
func (o *MultiCastChannelMap) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MultiCastChannelMap
func (o *MultiCastChannelMap) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MultiCastChannelMap
func (o *MultiCastChannelMap) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MultiCastChannelMap
func (o *MultiCastChannelMap) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MultiCastChannelMap
func (o *MultiCastChannelMap) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// MultiCastRanges retrieves the list of child MultiCastRanges of the MultiCastChannelMap
func (o *MultiCastChannelMap) MultiCastRanges(info *bambou.FetchingInfo) (MultiCastRangesList, *bambou.Error) {

	var list MultiCastRangesList
	err := bambou.CurrentSession().FetchChildren(o, MultiCastRangeIdentity, &list, info)
	return list, err
}

// CreateMultiCastRange creates a new child MultiCastRange under the MultiCastChannelMap
func (o *MultiCastChannelMap) CreateMultiCastRange(child *MultiCastRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the MultiCastChannelMap
func (o *MultiCastChannelMap) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
