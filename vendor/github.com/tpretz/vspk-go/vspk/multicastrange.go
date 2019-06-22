/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MultiCastRangeIdentity represents the Identity of the object
var MultiCastRangeIdentity = bambou.Identity{
	Name:     "multicastrange",
	Category: "multicastranges",
}

// MultiCastRangesList represents a list of MultiCastRanges
type MultiCastRangesList []*MultiCastRange

// MultiCastRangesAncestor is the interface that an ancestor of a MultiCastRange must implement.
// An Ancestor is defined as an entity that has MultiCastRange as a descendant.
// An Ancestor can get a list of its child MultiCastRanges, but not necessarily create one.
type MultiCastRangesAncestor interface {
	MultiCastRanges(*bambou.FetchingInfo) (MultiCastRangesList, *bambou.Error)
}

// MultiCastRangesParent is the interface that a parent of a MultiCastRange must implement.
// A Parent is defined as an entity that has MultiCastRange as a child.
// A Parent is an Ancestor which can create a MultiCastRange.
type MultiCastRangesParent interface {
	MultiCastRangesAncestor
	CreateMultiCastRange(*MultiCastRange) *bambou.Error
}

// MultiCastRange represents the model of a multicastrange
type MultiCastRange struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	MaxAddress    string `json:"maxAddress,omitempty"`
	MinAddress    string `json:"minAddress,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewMultiCastRange returns a new *MultiCastRange
func NewMultiCastRange() *MultiCastRange {

	return &MultiCastRange{}
}

// Identity returns the Identity of the object.
func (o *MultiCastRange) Identity() bambou.Identity {

	return MultiCastRangeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MultiCastRange) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MultiCastRange) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MultiCastRange from the server
func (o *MultiCastRange) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MultiCastRange into the server
func (o *MultiCastRange) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MultiCastRange from the server
func (o *MultiCastRange) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MultiCastRange
func (o *MultiCastRange) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MultiCastRange
func (o *MultiCastRange) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MultiCastRange
func (o *MultiCastRange) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MultiCastRange
func (o *MultiCastRange) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the MultiCastRange
func (o *MultiCastRange) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
