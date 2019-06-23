/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

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
	RouteDistinguisher  string `json:"routeDistinguisher,omitempty"`
	Priority            int    `json:"priority,omitempty"`
	AssociatedGatewayID string `json:"associatedGatewayID,omitempty"`
	AssociatedVLANID    string `json:"associatedVLANID,omitempty"`
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
