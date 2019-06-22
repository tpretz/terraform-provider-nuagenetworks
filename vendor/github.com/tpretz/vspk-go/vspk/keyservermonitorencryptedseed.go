/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// KeyServerMonitorEncryptedSeedIdentity represents the Identity of the object
var KeyServerMonitorEncryptedSeedIdentity = bambou.Identity{
	Name:     "keyservermonitorencryptedseed",
	Category: "keyservermonitorencryptedseeds",
}

// KeyServerMonitorEncryptedSeedsList represents a list of KeyServerMonitorEncryptedSeeds
type KeyServerMonitorEncryptedSeedsList []*KeyServerMonitorEncryptedSeed

// KeyServerMonitorEncryptedSeedsAncestor is the interface that an ancestor of a KeyServerMonitorEncryptedSeed must implement.
// An Ancestor is defined as an entity that has KeyServerMonitorEncryptedSeed as a descendant.
// An Ancestor can get a list of its child KeyServerMonitorEncryptedSeeds, but not necessarily create one.
type KeyServerMonitorEncryptedSeedsAncestor interface {
	KeyServerMonitorEncryptedSeeds(*bambou.FetchingInfo) (KeyServerMonitorEncryptedSeedsList, *bambou.Error)
}

// KeyServerMonitorEncryptedSeedsParent is the interface that a parent of a KeyServerMonitorEncryptedSeed must implement.
// A Parent is defined as an entity that has KeyServerMonitorEncryptedSeed as a child.
// A Parent is an Ancestor which can create a KeyServerMonitorEncryptedSeed.
type KeyServerMonitorEncryptedSeedsParent interface {
	KeyServerMonitorEncryptedSeedsAncestor
	CreateKeyServerMonitorEncryptedSeed(*KeyServerMonitorEncryptedSeed) *bambou.Error
}

// KeyServerMonitorEncryptedSeed represents the model of a keyservermonitorencryptedseed
type KeyServerMonitorEncryptedSeed struct {
	ID                                         string `json:"ID,omitempty"`
	ParentID                                   string `json:"parentID,omitempty"`
	ParentType                                 string `json:"parentType,omitempty"`
	Owner                                      string `json:"owner,omitempty"`
	SEKCreationTime                            int    `json:"SEKCreationTime,omitempty"`
	LastUpdatedBy                              string `json:"lastUpdatedBy,omitempty"`
	SeedType                                   string `json:"seedType,omitempty"`
	KeyServerCertificateSerialNumber           int    `json:"keyServerCertificateSerialNumber,omitempty"`
	EnterpriseSecuredDataID                    string `json:"enterpriseSecuredDataID,omitempty"`
	EntityScope                                string `json:"entityScope,omitempty"`
	AssociatedKeyServerMonitorSEKCreationTime  int    `json:"associatedKeyServerMonitorSEKCreationTime,omitempty"`
	AssociatedKeyServerMonitorSEKID            string `json:"associatedKeyServerMonitorSEKID,omitempty"`
	AssociatedKeyServerMonitorSeedCreationTime int    `json:"associatedKeyServerMonitorSeedCreationTime,omitempty"`
	AssociatedKeyServerMonitorSeedID           string `json:"associatedKeyServerMonitorSeedID,omitempty"`
	ExternalID                                 string `json:"externalID,omitempty"`
}

// NewKeyServerMonitorEncryptedSeed returns a new *KeyServerMonitorEncryptedSeed
func NewKeyServerMonitorEncryptedSeed() *KeyServerMonitorEncryptedSeed {

	return &KeyServerMonitorEncryptedSeed{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitorEncryptedSeed) Identity() bambou.Identity {

	return KeyServerMonitorEncryptedSeedIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitorEncryptedSeed) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitorEncryptedSeed) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitorEncryptedSeed from the server
func (o *KeyServerMonitorEncryptedSeed) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitorEncryptedSeed into the server
func (o *KeyServerMonitorEncryptedSeed) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitorEncryptedSeed from the server
func (o *KeyServerMonitorEncryptedSeed) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitorEncryptedSeed
func (o *KeyServerMonitorEncryptedSeed) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitorEncryptedSeed
func (o *KeyServerMonitorEncryptedSeed) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitorEncryptedSeed
func (o *KeyServerMonitorEncryptedSeed) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitorEncryptedSeed
func (o *KeyServerMonitorEncryptedSeed) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
