/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// BFDSessionIdentity represents the Identity of the object
var BFDSessionIdentity = bambou.Identity{
	Name:     "bfdsession",
	Category: "bfdsessions",
}

// BFDSessionsList represents a list of BFDSessions
type BFDSessionsList []*BFDSession

// BFDSessionsAncestor is the interface that an ancestor of a BFDSession must implement.
// An Ancestor is defined as an entity that has BFDSession as a descendant.
// An Ancestor can get a list of its child BFDSessions, but not necessarily create one.
type BFDSessionsAncestor interface {
	BFDSessions(*bambou.FetchingInfo) (BFDSessionsList, *bambou.Error)
}

// BFDSessionsParent is the interface that a parent of a BFDSession must implement.
// A Parent is defined as an entity that has BFDSession as a child.
// A Parent is an Ancestor which can create a BFDSession.
type BFDSessionsParent interface {
	BFDSessionsAncestor
	CreateBFDSession(*BFDSession) *bambou.Error
}

// BFDSession represents the model of a bfdsession
type BFDSession struct {
	ID               string `json:"ID,omitempty"`
	ParentID         string `json:"parentID,omitempty"`
	ParentType       string `json:"parentType,omitempty"`
	Owner            string `json:"owner,omitempty"`
	BFDDestinationIP string `json:"BFDDestinationIP,omitempty"`
	BFDMultiplier    int    `json:"BFDMultiplier,omitempty"`
	BFDTimer         int    `json:"BFDTimer,omitempty"`
	LastUpdatedBy    string `json:"lastUpdatedBy,omitempty"`
	EntityScope      string `json:"entityScope,omitempty"`
	MultiHopEnabled  bool   `json:"multiHopEnabled"`
	ExternalID       string `json:"externalID,omitempty"`
}

// NewBFDSession returns a new *BFDSession
func NewBFDSession() *BFDSession {

	return &BFDSession{
		BFDMultiplier:   3,
		BFDTimer:        500,
		MultiHopEnabled: false,
	}
}

// Identity returns the Identity of the object.
func (o *BFDSession) Identity() bambou.Identity {

	return BFDSessionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BFDSession) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BFDSession) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BFDSession from the server
func (o *BFDSession) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BFDSession into the server
func (o *BFDSession) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BFDSession from the server
func (o *BFDSession) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BFDSession
func (o *BFDSession) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BFDSession
func (o *BFDSession) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BFDSession
func (o *BFDSession) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BFDSession
func (o *BFDSession) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
