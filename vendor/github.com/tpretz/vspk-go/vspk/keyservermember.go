/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// KeyServerMemberIdentity represents the Identity of the object
var KeyServerMemberIdentity = bambou.Identity{
	Name:     "keyservermember",
	Category: "keyservermembers",
}

// KeyServerMembersList represents a list of KeyServerMembers
type KeyServerMembersList []*KeyServerMember

// KeyServerMembersAncestor is the interface that an ancestor of a KeyServerMember must implement.
// An Ancestor is defined as an entity that has KeyServerMember as a descendant.
// An Ancestor can get a list of its child KeyServerMembers, but not necessarily create one.
type KeyServerMembersAncestor interface {
	KeyServerMembers(*bambou.FetchingInfo) (KeyServerMembersList, *bambou.Error)
}

// KeyServerMembersParent is the interface that a parent of a KeyServerMember must implement.
// A Parent is defined as an entity that has KeyServerMember as a child.
// A Parent is an Ancestor which can create a KeyServerMember.
type KeyServerMembersParent interface {
	KeyServerMembersAncestor
	CreateKeyServerMember(*KeyServerMember) *bambou.Error
}

// KeyServerMember represents the model of a keyservermember
type KeyServerMember struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	PemEncoded              string `json:"pemEncoded,omitempty"`
	CertificateSerialNumber int    `json:"certificateSerialNumber,omitempty"`
	EntityScope             string `json:"entityScope,omitempty"`
	Fqdn                    string `json:"fqdn,omitempty"`
	IssuerDN                string `json:"issuerDN,omitempty"`
	SubjectDN               string `json:"subjectDN,omitempty"`
	PublicKey               string `json:"publicKey,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
}

// NewKeyServerMember returns a new *KeyServerMember
func NewKeyServerMember() *KeyServerMember {

	return &KeyServerMember{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMember) Identity() bambou.Identity {

	return KeyServerMemberIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMember) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMember) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMember from the server
func (o *KeyServerMember) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMember into the server
func (o *KeyServerMember) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMember from the server
func (o *KeyServerMember) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMember
func (o *KeyServerMember) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMember
func (o *KeyServerMember) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMember
func (o *KeyServerMember) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMember
func (o *KeyServerMember) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
