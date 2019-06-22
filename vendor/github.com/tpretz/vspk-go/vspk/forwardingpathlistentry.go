/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ForwardingPathListEntryIdentity represents the Identity of the object
var ForwardingPathListEntryIdentity = bambou.Identity{
	Name:     "forwardingpathlistentry",
	Category: "forwardingpathlistentries",
}

// ForwardingPathListEntriesList represents a list of ForwardingPathListEntries
type ForwardingPathListEntriesList []*ForwardingPathListEntry

// ForwardingPathListEntriesAncestor is the interface that an ancestor of a ForwardingPathListEntry must implement.
// An Ancestor is defined as an entity that has ForwardingPathListEntry as a descendant.
// An Ancestor can get a list of its child ForwardingPathListEntries, but not necessarily create one.
type ForwardingPathListEntriesAncestor interface {
	ForwardingPathListEntries(*bambou.FetchingInfo) (ForwardingPathListEntriesList, *bambou.Error)
}

// ForwardingPathListEntriesParent is the interface that a parent of a ForwardingPathListEntry must implement.
// A Parent is defined as an entity that has ForwardingPathListEntry as a child.
// A Parent is an Ancestor which can create a ForwardingPathListEntry.
type ForwardingPathListEntriesParent interface {
	ForwardingPathListEntriesAncestor
	CreateForwardingPathListEntry(*ForwardingPathListEntry) *bambou.Error
}

// ForwardingPathListEntry represents the model of a forwardingpathlistentry
type ForwardingPathListEntry struct {
	ID               string `json:"ID,omitempty"`
	ParentID         string `json:"parentID,omitempty"`
	ParentType       string `json:"parentType,omitempty"`
	Owner            string `json:"owner,omitempty"`
	FCOverride       string `json:"FCOverride,omitempty"`
	LastUpdatedBy    string `json:"lastUpdatedBy,omitempty"`
	EntityScope      string `json:"entityScope,omitempty"`
	ForwardingAction string `json:"forwardingAction,omitempty"`
	UplinkPreference string `json:"uplinkPreference,omitempty"`
	Priority         int    `json:"priority,omitempty"`
	ExternalID       string `json:"externalID,omitempty"`
}

// NewForwardingPathListEntry returns a new *ForwardingPathListEntry
func NewForwardingPathListEntry() *ForwardingPathListEntry {

	return &ForwardingPathListEntry{}
}

// Identity returns the Identity of the object.
func (o *ForwardingPathListEntry) Identity() bambou.Identity {

	return ForwardingPathListEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ForwardingPathListEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ForwardingPathListEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ForwardingPathListEntry from the server
func (o *ForwardingPathListEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ForwardingPathListEntry into the server
func (o *ForwardingPathListEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ForwardingPathListEntry from the server
func (o *ForwardingPathListEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ForwardingPathListEntry
func (o *ForwardingPathListEntry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ForwardingPathListEntry
func (o *ForwardingPathListEntry) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ForwardingPathListEntry
func (o *ForwardingPathListEntry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ForwardingPathListEntry
func (o *ForwardingPathListEntry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
