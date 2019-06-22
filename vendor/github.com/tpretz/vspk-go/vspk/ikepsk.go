/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKEPSKIdentity represents the Identity of the object
var IKEPSKIdentity = bambou.Identity{
	Name:     "ikepsk",
	Category: "ikepsks",
}

// IKEPSKsList represents a list of IKEPSKs
type IKEPSKsList []*IKEPSK

// IKEPSKsAncestor is the interface that an ancestor of a IKEPSK must implement.
// An Ancestor is defined as an entity that has IKEPSK as a descendant.
// An Ancestor can get a list of its child IKEPSKs, but not necessarily create one.
type IKEPSKsAncestor interface {
	IKEPSKs(*bambou.FetchingInfo) (IKEPSKsList, *bambou.Error)
}

// IKEPSKsParent is the interface that a parent of a IKEPSK must implement.
// A Parent is defined as an entity that has IKEPSK as a child.
// A Parent is an Ancestor which can create a IKEPSK.
type IKEPSKsParent interface {
	IKEPSKsAncestor
	CreateIKEPSK(*IKEPSK) *bambou.Error
}

// IKEPSK represents the model of a ikepsk
type IKEPSK struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	Name                              string `json:"name,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	Description                       string `json:"description,omitempty"`
	Signature                         string `json:"signature,omitempty"`
	SigningCertificateSerialNumber    int    `json:"signingCertificateSerialNumber,omitempty"`
	EncryptedPSK                      string `json:"encryptedPSK,omitempty"`
	EncryptingCertificateSerialNumber int    `json:"encryptingCertificateSerialNumber,omitempty"`
	UnencryptedPSK                    string `json:"unencryptedPSK,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	AssociatedEnterpriseID            string `json:"associatedEnterpriseID,omitempty"`
	AutoCreated                       bool   `json:"autoCreated"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewIKEPSK returns a new *IKEPSK
func NewIKEPSK() *IKEPSK {

	return &IKEPSK{}
}

// Identity returns the Identity of the object.
func (o *IKEPSK) Identity() bambou.Identity {

	return IKEPSKIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEPSK) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEPSK) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEPSK from the server
func (o *IKEPSK) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEPSK into the server
func (o *IKEPSK) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEPSK from the server
func (o *IKEPSK) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEPSK
func (o *IKEPSK) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEPSK
func (o *IKEPSK) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEPSK
func (o *IKEPSK) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEPSK
func (o *IKEPSK) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
