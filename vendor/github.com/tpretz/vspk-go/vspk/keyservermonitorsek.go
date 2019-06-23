/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// KeyServerMonitorSEKIdentity represents the Identity of the object
var KeyServerMonitorSEKIdentity = bambou.Identity{
	Name:     "keyservermonitorsek",
	Category: "keyservermonitorseks",
}

// KeyServerMonitorSEKsList represents a list of KeyServerMonitorSEKs
type KeyServerMonitorSEKsList []*KeyServerMonitorSEK

// KeyServerMonitorSEKsAncestor is the interface that an ancestor of a KeyServerMonitorSEK must implement.
// An Ancestor is defined as an entity that has KeyServerMonitorSEK as a descendant.
// An Ancestor can get a list of its child KeyServerMonitorSEKs, but not necessarily create one.
type KeyServerMonitorSEKsAncestor interface {
	KeyServerMonitorSEKs(*bambou.FetchingInfo) (KeyServerMonitorSEKsList, *bambou.Error)
}

// KeyServerMonitorSEKsParent is the interface that a parent of a KeyServerMonitorSEK must implement.
// A Parent is defined as an entity that has KeyServerMonitorSEK as a child.
// A Parent is an Ancestor which can create a KeyServerMonitorSEK.
type KeyServerMonitorSEKsParent interface {
	KeyServerMonitorSEKsAncestor
	CreateKeyServerMonitorSEK(*KeyServerMonitorSEK) *bambou.Error
}

// KeyServerMonitorSEK represents the model of a keyservermonitorsek
type KeyServerMonitorSEK struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	SeedPayloadAuthenticationAlgorithm string `json:"seedPayloadAuthenticationAlgorithm,omitempty"`
	SeedPayloadEncryptionAlgorithm     string `json:"seedPayloadEncryptionAlgorithm,omitempty"`
	Lifetime                           int    `json:"lifetime,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	CreationTime                       int    `json:"creationTime,omitempty"`
	StartTime                          int    `json:"startTime,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
}

// NewKeyServerMonitorSEK returns a new *KeyServerMonitorSEK
func NewKeyServerMonitorSEK() *KeyServerMonitorSEK {

	return &KeyServerMonitorSEK{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitorSEK) Identity() bambou.Identity {

	return KeyServerMonitorSEKIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitorSEK) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitorSEK) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitorSEK from the server
func (o *KeyServerMonitorSEK) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitorSEK into the server
func (o *KeyServerMonitorSEK) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitorSEK from the server
func (o *KeyServerMonitorSEK) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitorSEK
func (o *KeyServerMonitorSEK) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
