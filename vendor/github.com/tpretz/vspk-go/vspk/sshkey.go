/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SSHKeyIdentity represents the Identity of the object
var SSHKeyIdentity = bambou.Identity{
	Name:     "sshkey",
	Category: "sshkeys",
}

// SSHKeysList represents a list of SSHKeys
type SSHKeysList []*SSHKey

// SSHKeysAncestor is the interface that an ancestor of a SSHKey must implement.
// An Ancestor is defined as an entity that has SSHKey as a descendant.
// An Ancestor can get a list of its child SSHKeys, but not necessarily create one.
type SSHKeysAncestor interface {
	SSHKeys(*bambou.FetchingInfo) (SSHKeysList, *bambou.Error)
}

// SSHKeysParent is the interface that a parent of a SSHKey must implement.
// A Parent is defined as an entity that has SSHKey as a child.
// A Parent is an Ancestor which can create a SSHKey.
type SSHKeysParent interface {
	SSHKeysAncestor
	CreateSSHKey(*SSHKey) *bambou.Error
}

// SSHKey represents the model of a sshkey
type SSHKey struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	KeyType       string `json:"keyType,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	PublicKey     string `json:"publicKey,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewSSHKey returns a new *SSHKey
func NewSSHKey() *SSHKey {

	return &SSHKey{
		KeyType: "RSA",
	}
}

// Identity returns the Identity of the object.
func (o *SSHKey) Identity() bambou.Identity {

	return SSHKeyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SSHKey) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SSHKey) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SSHKey from the server
func (o *SSHKey) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SSHKey into the server
func (o *SSHKey) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SSHKey from the server
func (o *SSHKey) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the SSHKey
func (o *SSHKey) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SSHKey
func (o *SSHKey) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SSHKey
func (o *SSHKey) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SSHKey
func (o *SSHKey) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
