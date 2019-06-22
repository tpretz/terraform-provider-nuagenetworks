/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DemarcationServiceIdentity represents the Identity of the object
var DemarcationServiceIdentity = bambou.Identity{
	Name:     "demarcationservice",
	Category: "demarcationservices",
}

// DemarcationServicesList represents a list of DemarcationServices
type DemarcationServicesList []*DemarcationService

// DemarcationServicesAncestor is the interface that an ancestor of a DemarcationService must implement.
// An Ancestor is defined as an entity that has DemarcationService as a descendant.
// An Ancestor can get a list of its child DemarcationServices, but not necessarily create one.
type DemarcationServicesAncestor interface {
	DemarcationServices(*bambou.FetchingInfo) (DemarcationServicesList, *bambou.Error)
}

// DemarcationServicesParent is the interface that a parent of a DemarcationService must implement.
// A Parent is defined as an entity that has DemarcationService as a child.
// A Parent is an Ancestor which can create a DemarcationService.
type DemarcationServicesParent interface {
	DemarcationServicesAncestor
	CreateDemarcationService(*DemarcationService) *bambou.Error
}

// DemarcationService represents the model of a demarcationservice
type DemarcationService struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	RouteDistinguisher  string `json:"routeDistinguisher,omitempty"`
	Priority            int    `json:"priority,omitempty"`
	AssociatedGatewayID string `json:"associatedGatewayID,omitempty"`
	AssociatedVLANID    string `json:"associatedVLANID,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
	Type                string `json:"type,omitempty"`
}

// NewDemarcationService returns a new *DemarcationService
func NewDemarcationService() *DemarcationService {

	return &DemarcationService{}
}

// Identity returns the Identity of the object.
func (o *DemarcationService) Identity() bambou.Identity {

	return DemarcationServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DemarcationService) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DemarcationService) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DemarcationService from the server
func (o *DemarcationService) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DemarcationService into the server
func (o *DemarcationService) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DemarcationService from the server
func (o *DemarcationService) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DemarcationService
func (o *DemarcationService) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DemarcationService
func (o *DemarcationService) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DemarcationService
func (o *DemarcationService) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DemarcationService
func (o *DemarcationService) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
