/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VPortMirrorIdentity represents the Identity of the object
var VPortMirrorIdentity = bambou.Identity{
	Name:     "vportmirror",
	Category: "vportmirrors",
}

// VPortMirrorsList represents a list of VPortMirrors
type VPortMirrorsList []*VPortMirror

// VPortMirrorsAncestor is the interface that an ancestor of a VPortMirror must implement.
// An Ancestor is defined as an entity that has VPortMirror as a descendant.
// An Ancestor can get a list of its child VPortMirrors, but not necessarily create one.
type VPortMirrorsAncestor interface {
	VPortMirrors(*bambou.FetchingInfo) (VPortMirrorsList, *bambou.Error)
}

// VPortMirrorsParent is the interface that a parent of a VPortMirror must implement.
// A Parent is defined as an entity that has VPortMirror as a child.
// A Parent is an Ancestor which can create a VPortMirror.
type VPortMirrorsParent interface {
	VPortMirrorsAncestor
	CreateVPortMirror(*VPortMirror) *bambou.Error
}

// VPortMirror represents the model of a vportmirror
type VPortMirror struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	VPortName             string `json:"VPortName,omitempty"`
	LastUpdatedBy         string `json:"lastUpdatedBy,omitempty"`
	NetworkName           string `json:"networkName,omitempty"`
	MirrorDestinationID   string `json:"mirrorDestinationID,omitempty"`
	MirrorDestinationName string `json:"mirrorDestinationName,omitempty"`
	MirrorDirection       string `json:"mirrorDirection,omitempty"`
	EnterpiseName         string `json:"enterpiseName,omitempty"`
	EntityScope           string `json:"entityScope,omitempty"`
	DomainName            string `json:"domainName,omitempty"`
	VportId               string `json:"vportId,omitempty"`
	AttachedNetworkType   string `json:"attachedNetworkType,omitempty"`
	ExternalID            string `json:"externalID,omitempty"`
}

// NewVPortMirror returns a new *VPortMirror
func NewVPortMirror() *VPortMirror {

	return &VPortMirror{}
}

// Identity returns the Identity of the object.
func (o *VPortMirror) Identity() bambou.Identity {

	return VPortMirrorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VPortMirror) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VPortMirror) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VPortMirror from the server
func (o *VPortMirror) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VPortMirror into the server
func (o *VPortMirror) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VPortMirror from the server
func (o *VPortMirror) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VPortMirror
func (o *VPortMirror) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VPortMirror
func (o *VPortMirror) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VPortMirror
func (o *VPortMirror) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VPortMirror
func (o *VPortMirror) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
