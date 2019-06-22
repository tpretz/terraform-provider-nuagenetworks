/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// OverlayPATNATEntryIdentity represents the Identity of the object
var OverlayPATNATEntryIdentity = bambou.Identity{
	Name:     "overlaypatnatentry",
	Category: "overlaypatnatentries",
}

// OverlayPATNATEntriesList represents a list of OverlayPATNATEntries
type OverlayPATNATEntriesList []*OverlayPATNATEntry

// OverlayPATNATEntriesAncestor is the interface that an ancestor of a OverlayPATNATEntry must implement.
// An Ancestor is defined as an entity that has OverlayPATNATEntry as a descendant.
// An Ancestor can get a list of its child OverlayPATNATEntries, but not necessarily create one.
type OverlayPATNATEntriesAncestor interface {
	OverlayPATNATEntries(*bambou.FetchingInfo) (OverlayPATNATEntriesList, *bambou.Error)
}

// OverlayPATNATEntriesParent is the interface that a parent of a OverlayPATNATEntry must implement.
// A Parent is defined as an entity that has OverlayPATNATEntry as a child.
// A Parent is an Ancestor which can create a OverlayPATNATEntry.
type OverlayPATNATEntriesParent interface {
	OverlayPATNATEntriesAncestor
	CreateOverlayPATNATEntry(*OverlayPATNATEntry) *bambou.Error
}

// OverlayPATNATEntry represents the model of a overlaypatnatentry
type OverlayPATNATEntry struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	NATEnabled         bool   `json:"NATEnabled"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	PrivateIP          string `json:"privateIP,omitempty"`
	AssociatedDomainID string `json:"associatedDomainID,omitempty"`
	AssociatedLinkID   string `json:"associatedLinkID,omitempty"`
	PublicIP           string `json:"publicIP,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewOverlayPATNATEntry returns a new *OverlayPATNATEntry
func NewOverlayPATNATEntry() *OverlayPATNATEntry {

	return &OverlayPATNATEntry{
		NATEnabled: true,
	}
}

// Identity returns the Identity of the object.
func (o *OverlayPATNATEntry) Identity() bambou.Identity {

	return OverlayPATNATEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OverlayPATNATEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OverlayPATNATEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OverlayPATNATEntry from the server
func (o *OverlayPATNATEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OverlayPATNATEntry into the server
func (o *OverlayPATNATEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OverlayPATNATEntry from the server
func (o *OverlayPATNATEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the OverlayPATNATEntry
func (o *OverlayPATNATEntry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the OverlayPATNATEntry
func (o *OverlayPATNATEntry) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the OverlayPATNATEntry
func (o *OverlayPATNATEntry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the OverlayPATNATEntry
func (o *OverlayPATNATEntry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
