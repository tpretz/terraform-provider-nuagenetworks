/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CertificateIdentity represents the Identity of the object
var CertificateIdentity = bambou.Identity{
	Name:     "certificate",
	Category: "certificates",
}

// CertificatesList represents a list of Certificates
type CertificatesList []*Certificate

// CertificatesAncestor is the interface that an ancestor of a Certificate must implement.
// An Ancestor is defined as an entity that has Certificate as a descendant.
// An Ancestor can get a list of its child Certificates, but not necessarily create one.
type CertificatesAncestor interface {
	Certificates(*bambou.FetchingInfo) (CertificatesList, *bambou.Error)
}

// CertificatesParent is the interface that a parent of a Certificate must implement.
// A Parent is defined as an entity that has Certificate as a child.
// A Parent is an Ancestor which can create a Certificate.
type CertificatesParent interface {
	CertificatesAncestor
	CreateCertificate(*Certificate) *bambou.Error
}

// Certificate represents the model of a certificate
type Certificate struct {
	ID           string `json:"ID,omitempty"`
	ParentID     string `json:"parentID,omitempty"`
	ParentType   string `json:"parentType,omitempty"`
	Owner        string `json:"owner,omitempty"`
	PemEncoded   string `json:"pemEncoded,omitempty"`
	SerialNumber int    `json:"serialNumber,omitempty"`
	EntityScope  string `json:"entityScope,omitempty"`
	IssuerDN     string `json:"issuerDN,omitempty"`
	SubjectDN    string `json:"subjectDN,omitempty"`
	PublicKey    string `json:"publicKey,omitempty"`
	ExternalID   string `json:"externalID,omitempty"`
}

// NewCertificate returns a new *Certificate
func NewCertificate() *Certificate {

	return &Certificate{}
}

// Identity returns the Identity of the object.
func (o *Certificate) Identity() bambou.Identity {

	return CertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Certificate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Certificate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Certificate from the server
func (o *Certificate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Certificate into the server
func (o *Certificate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Certificate from the server
func (o *Certificate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Certificate
func (o *Certificate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Certificate
func (o *Certificate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Certificate
func (o *Certificate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Certificate
func (o *Certificate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
