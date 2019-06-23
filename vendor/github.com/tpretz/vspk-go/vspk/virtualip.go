/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VirtualIPIdentity represents the Identity of the object
var VirtualIPIdentity = bambou.Identity{
	Name:     "virtualip",
	Category: "virtualips",
}

// VirtualIPsList represents a list of VirtualIPs
type VirtualIPsList []*VirtualIP

// VirtualIPsAncestor is the interface that an ancestor of a VirtualIP must implement.
// An Ancestor is defined as an entity that has VirtualIP as a descendant.
// An Ancestor can get a list of its child VirtualIPs, but not necessarily create one.
type VirtualIPsAncestor interface {
	VirtualIPs(*bambou.FetchingInfo) (VirtualIPsList, *bambou.Error)
}

// VirtualIPsParent is the interface that a parent of a VirtualIP must implement.
// A Parent is defined as an entity that has VirtualIP as a child.
// A Parent is an Ancestor which can create a VirtualIP.
type VirtualIPsParent interface {
	VirtualIPsAncestor
	CreateVirtualIP(*VirtualIP) *bambou.Error
}

// VirtualIP represents the model of a virtualip
type VirtualIP struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	MAC                    string `json:"MAC,omitempty"`
	IPType                 string `json:"IPType,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	VirtualIP              string `json:"virtualIP,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	AssociatedFloatingIPID string `json:"associatedFloatingIPID,omitempty"`
	SubnetID               string `json:"subnetID,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewVirtualIP returns a new *VirtualIP
func NewVirtualIP() *VirtualIP {

	return &VirtualIP{
		IPType: "IPV4",
	}
}

// Identity returns the Identity of the object.
func (o *VirtualIP) Identity() bambou.Identity {

	return VirtualIPIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VirtualIP) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VirtualIP) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VirtualIP from the server
func (o *VirtualIP) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VirtualIP into the server
func (o *VirtualIP) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VirtualIP from the server
func (o *VirtualIP) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VirtualIP
func (o *VirtualIP) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VirtualIP
func (o *VirtualIP) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VirtualIP
func (o *VirtualIP) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VirtualIP
func (o *VirtualIP) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the VirtualIP
func (o *VirtualIP) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
