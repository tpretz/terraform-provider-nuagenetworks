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
	NATEnabled         string `json:"NATEnabled,omitempty"`
	PrivateIP          string `json:"privateIP,omitempty"`
	AssociatedDomainID string `json:"associatedDomainID,omitempty"`
	AssociatedLinkID   string `json:"associatedLinkID,omitempty"`
	PublicIP           string `json:"publicIP,omitempty"`
}

// NewOverlayPATNATEntry returns a new *OverlayPATNATEntry
func NewOverlayPATNATEntry() *OverlayPATNATEntry {

	return &OverlayPATNATEntry{}
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
