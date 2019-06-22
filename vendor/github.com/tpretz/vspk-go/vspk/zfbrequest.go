/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ZFBRequestIdentity represents the Identity of the object
var ZFBRequestIdentity = bambou.Identity{
	Name:     "zfbrequest",
	Category: "zfbrequests",
}

// ZFBRequestsList represents a list of ZFBRequests
type ZFBRequestsList []*ZFBRequest

// ZFBRequestsAncestor is the interface that an ancestor of a ZFBRequest must implement.
// An Ancestor is defined as an entity that has ZFBRequest as a descendant.
// An Ancestor can get a list of its child ZFBRequests, but not necessarily create one.
type ZFBRequestsAncestor interface {
	ZFBRequests(*bambou.FetchingInfo) (ZFBRequestsList, *bambou.Error)
}

// ZFBRequestsParent is the interface that a parent of a ZFBRequest must implement.
// A Parent is defined as an entity that has ZFBRequest as a child.
// A Parent is an Ancestor which can create a ZFBRequest.
type ZFBRequestsParent interface {
	ZFBRequestsAncestor
	CreateZFBRequest(*ZFBRequest) *bambou.Error
}

// ZFBRequest represents the model of a zfbrequest
type ZFBRequest struct {
	ID                       string  `json:"ID,omitempty"`
	ParentID                 string  `json:"parentID,omitempty"`
	ParentType               string  `json:"parentType,omitempty"`
	Owner                    string  `json:"owner,omitempty"`
	MACAddress               string  `json:"MACAddress,omitempty"`
	ZFBApprovalStatus        string  `json:"ZFBApprovalStatus,omitempty"`
	ZFBBootstrapEnabled      bool    `json:"ZFBBootstrapEnabled"`
	ZFBInfo                  string  `json:"ZFBInfo,omitempty"`
	ZFBRequestRetryTimer     int     `json:"ZFBRequestRetryTimer,omitempty"`
	SKU                      string  `json:"SKU,omitempty"`
	IPAddress                string  `json:"IPAddress,omitempty"`
	CPUType                  string  `json:"CPUType,omitempty"`
	NSGVersion               string  `json:"NSGVersion,omitempty"`
	UUID                     string  `json:"UUID,omitempty"`
	Family                   string  `json:"family,omitempty"`
	LastConnectedTime        float64 `json:"lastConnectedTime,omitempty"`
	LastUpdatedBy            string  `json:"lastUpdatedBy,omitempty"`
	SerialNumber             string  `json:"serialNumber,omitempty"`
	EntityScope              string  `json:"entityScope,omitempty"`
	Hostname                 string  `json:"hostname,omitempty"`
	AssociatedEnterpriseID   string  `json:"associatedEnterpriseID,omitempty"`
	AssociatedEnterpriseName string  `json:"associatedEnterpriseName,omitempty"`
	AssociatedEntityType     string  `json:"associatedEntityType,omitempty"`
	AssociatedGatewayID      string  `json:"associatedGatewayID,omitempty"`
	AssociatedGatewayName    string  `json:"associatedGatewayName,omitempty"`
	AssociatedNSGatewayID    string  `json:"associatedNSGatewayID,omitempty"`
	AssociatedNSGatewayName  string  `json:"associatedNSGatewayName,omitempty"`
	StatusString             string  `json:"statusString,omitempty"`
	ExternalID               string  `json:"externalID,omitempty"`
}

// NewZFBRequest returns a new *ZFBRequest
func NewZFBRequest() *ZFBRequest {

	return &ZFBRequest{}
}

// Identity returns the Identity of the object.
func (o *ZFBRequest) Identity() bambou.Identity {

	return ZFBRequestIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ZFBRequest) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ZFBRequest) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ZFBRequest from the server
func (o *ZFBRequest) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ZFBRequest into the server
func (o *ZFBRequest) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ZFBRequest from the server
func (o *ZFBRequest) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ZFBRequest
func (o *ZFBRequest) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ZFBRequest
func (o *ZFBRequest) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ZFBRequest
func (o *ZFBRequest) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ZFBRequest
func (o *ZFBRequest) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the ZFBRequest
func (o *ZFBRequest) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
