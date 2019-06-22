/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NATMapEntryIdentity represents the Identity of the object
var NATMapEntryIdentity = bambou.Identity{
	Name:     "natmapentry",
	Category: "natmapentries",
}

// NATMapEntriesList represents a list of NATMapEntries
type NATMapEntriesList []*NATMapEntry

// NATMapEntriesAncestor is the interface that an ancestor of a NATMapEntry must implement.
// An Ancestor is defined as an entity that has NATMapEntry as a descendant.
// An Ancestor can get a list of its child NATMapEntries, but not necessarily create one.
type NATMapEntriesAncestor interface {
	NATMapEntries(*bambou.FetchingInfo) (NATMapEntriesList, *bambou.Error)
}

// NATMapEntriesParent is the interface that a parent of a NATMapEntry must implement.
// A Parent is defined as an entity that has NATMapEntry as a child.
// A Parent is an Ancestor which can create a NATMapEntry.
type NATMapEntriesParent interface {
	NATMapEntriesAncestor
	CreateNATMapEntry(*NATMapEntry) *bambou.Error
}

// NATMapEntry represents the model of a natmapentry
type NATMapEntry struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	PrivateIP              string `json:"privateIP,omitempty"`
	PrivatePort            int    `json:"privatePort,omitempty"`
	AssociatedPATNATPoolID string `json:"associatedPATNATPoolID,omitempty"`
	PublicIP               string `json:"publicIP,omitempty"`
	PublicPort             int    `json:"publicPort,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
	Type                   string `json:"type,omitempty"`
}

// NewNATMapEntry returns a new *NATMapEntry
func NewNATMapEntry() *NATMapEntry {

	return &NATMapEntry{
		Type: "ONE_TO_ONE_NAT",
	}
}

// Identity returns the Identity of the object.
func (o *NATMapEntry) Identity() bambou.Identity {

	return NATMapEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NATMapEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NATMapEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NATMapEntry from the server
func (o *NATMapEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NATMapEntry into the server
func (o *NATMapEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NATMapEntry from the server
func (o *NATMapEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NATMapEntry
func (o *NATMapEntry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NATMapEntry
func (o *NATMapEntry) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NATMapEntry
func (o *NATMapEntry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NATMapEntry
func (o *NATMapEntry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
