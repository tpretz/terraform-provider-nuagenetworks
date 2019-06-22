/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKEGatewayIdentity represents the Identity of the object
var IKEGatewayIdentity = bambou.Identity{
	Name:     "ikegateway",
	Category: "ikegateways",
}

// IKEGatewaysList represents a list of IKEGateways
type IKEGatewaysList []*IKEGateway

// IKEGatewaysAncestor is the interface that an ancestor of a IKEGateway must implement.
// An Ancestor is defined as an entity that has IKEGateway as a descendant.
// An Ancestor can get a list of its child IKEGateways, but not necessarily create one.
type IKEGatewaysAncestor interface {
	IKEGateways(*bambou.FetchingInfo) (IKEGatewaysList, *bambou.Error)
}

// IKEGatewaysParent is the interface that a parent of a IKEGateway must implement.
// A Parent is defined as an entity that has IKEGateway as a child.
// A Parent is an Ancestor which can create a IKEGateway.
type IKEGatewaysParent interface {
	IKEGatewaysAncestor
	CreateIKEGateway(*IKEGateway) *bambou.Error
}

// IKEGateway represents the model of a ikegateway
type IKEGateway struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	IKEVersion             string `json:"IKEVersion,omitempty"`
	IKEv1Mode              string `json:"IKEv1Mode,omitempty"`
	IPAddress              string `json:"IPAddress,omitempty"`
	Name                   string `json:"name,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	Description            string `json:"description,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	AssociatedEnterpriseID string `json:"associatedEnterpriseID,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewIKEGateway returns a new *IKEGateway
func NewIKEGateway() *IKEGateway {

	return &IKEGateway{
		IKEVersion: "V2",
		IKEv1Mode:  "NONE",
	}
}

// Identity returns the Identity of the object.
func (o *IKEGateway) Identity() bambou.Identity {

	return IKEGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEGateway) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEGateway from the server
func (o *IKEGateway) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEGateway into the server
func (o *IKEGateway) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEGateway from the server
func (o *IKEGateway) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEGateway
func (o *IKEGateway) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEGateway
func (o *IKEGateway) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IKEGatewayConfigs retrieves the list of child IKEGatewayConfigs of the IKEGateway
func (o *IKEGateway) IKEGatewayConfigs(info *bambou.FetchingInfo) (IKEGatewayConfigsList, *bambou.Error) {

	var list IKEGatewayConfigsList
	err := bambou.CurrentSession().FetchChildren(o, IKEGatewayConfigIdentity, &list, info)
	return list, err
}

// AssignIKEGatewayConfigs assigns the list of IKEGatewayConfigs to the IKEGateway
func (o *IKEGateway) AssignIKEGatewayConfigs(children IKEGatewayConfigsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, IKEGatewayConfigIdentity)
}

// IKESubnets retrieves the list of child IKESubnets of the IKEGateway
func (o *IKEGateway) IKESubnets(info *bambou.FetchingInfo) (IKESubnetsList, *bambou.Error) {

	var list IKESubnetsList
	err := bambou.CurrentSession().FetchChildren(o, IKESubnetIdentity, &list, info)
	return list, err
}

// CreateIKESubnet creates a new child IKESubnet under the IKEGateway
func (o *IKEGateway) CreateIKESubnet(child *IKESubnet) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEGateway
func (o *IKEGateway) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEGateway
func (o *IKEGateway) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
