/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// IPReservationIdentity represents the Identity of the object
var IPReservationIdentity = bambou.Identity{
	Name:     "ipreservation",
	Category: "ipreservations",
}

// IPReservationsList represents a list of IPReservations
type IPReservationsList []*IPReservation

// IPReservationsAncestor is the interface that an ancestor of a IPReservation must implement.
// An Ancestor is defined as an entity that has IPReservation as a descendant.
// An Ancestor can get a list of its child IPReservations, but not necessarily create one.
type IPReservationsAncestor interface {
	IPReservations(*bambou.FetchingInfo) (IPReservationsList, *bambou.Error)
}

// IPReservationsParent is the interface that a parent of a IPReservation must implement.
// A Parent is defined as an entity that has IPReservation as a child.
// A Parent is an Ancestor which can create a IPReservation.
type IPReservationsParent interface {
	IPReservationsAncestor
	CreateIPReservation(*IPReservation) *bambou.Error
}

// IPReservation represents the model of a ipreservation
type IPReservation struct {
	ID                       string `json:"ID,omitempty"`
	ParentID                 string `json:"parentID,omitempty"`
	ParentType               string `json:"parentType,omitempty"`
	Owner                    string `json:"owner,omitempty"`
	MAC                      string `json:"MAC,omitempty"`
	IPAddress                string `json:"IPAddress,omitempty"`
	LastUpdatedBy            string `json:"lastUpdatedBy,omitempty"`
	EntityScope              string `json:"entityScope,omitempty"`
	ExternalID               string `json:"externalID,omitempty"`
	DynamicAllocationEnabled bool   `json:"dynamicAllocationEnabled"`
}

// NewIPReservation returns a new *IPReservation
func NewIPReservation() *IPReservation {

	return &IPReservation{}
}

// Identity returns the Identity of the object.
func (o *IPReservation) Identity() bambou.Identity {

	return IPReservationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IPReservation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IPReservation) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IPReservation from the server
func (o *IPReservation) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IPReservation into the server
func (o *IPReservation) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IPReservation from the server
func (o *IPReservation) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IPReservation
func (o *IPReservation) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IPReservation
func (o *IPReservation) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IPReservation
func (o *IPReservation) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IPReservation
func (o *IPReservation) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the IPReservation
func (o *IPReservation) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
