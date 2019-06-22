/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GatewaysLocationIdentity represents the Identity of the object
var GatewaysLocationIdentity = bambou.Identity{
	Name:     "gatewayslocation",
	Category: "gatewayslocations",
}

// GatewaysLocationsList represents a list of GatewaysLocations
type GatewaysLocationsList []*GatewaysLocation

// GatewaysLocationsAncestor is the interface that an ancestor of a GatewaysLocation must implement.
// An Ancestor is defined as an entity that has GatewaysLocation as a descendant.
// An Ancestor can get a list of its child GatewaysLocations, but not necessarily create one.
type GatewaysLocationsAncestor interface {
	GatewaysLocations(*bambou.FetchingInfo) (GatewaysLocationsList, *bambou.Error)
}

// GatewaysLocationsParent is the interface that a parent of a GatewaysLocation must implement.
// A Parent is defined as an entity that has GatewaysLocation as a child.
// A Parent is an Ancestor which can create a GatewaysLocation.
type GatewaysLocationsParent interface {
	GatewaysLocationsAncestor
	CreateGatewaysLocation(*GatewaysLocation) *bambou.Error
}

// GatewaysLocation represents the model of a gatewayslocation
type GatewaysLocation struct {
	ID                   string  `json:"ID,omitempty"`
	ParentID             string  `json:"parentID,omitempty"`
	ParentType           string  `json:"parentType,omitempty"`
	Owner                string  `json:"owner,omitempty"`
	LastUpdatedBy        string  `json:"lastUpdatedBy,omitempty"`
	Latitude             float64 `json:"latitude,omitempty"`
	Address              string  `json:"address,omitempty"`
	IgnoreGeocode        bool    `json:"ignoreGeocode"`
	TimeZoneID           string  `json:"timeZoneID,omitempty"`
	EntityScope          string  `json:"entityScope,omitempty"`
	Locality             string  `json:"locality,omitempty"`
	Longitude            float64 `json:"longitude,omitempty"`
	Country              string  `json:"country,omitempty"`
	AssociatedEntityName string  `json:"associatedEntityName,omitempty"`
	AssociatedEntityType string  `json:"associatedEntityType,omitempty"`
	State                string  `json:"state,omitempty"`
	ExternalID           string  `json:"externalID,omitempty"`
}

// NewGatewaysLocation returns a new *GatewaysLocation
func NewGatewaysLocation() *GatewaysLocation {

	return &GatewaysLocation{
		TimeZoneID: "UTC",
	}
}

// Identity returns the Identity of the object.
func (o *GatewaysLocation) Identity() bambou.Identity {

	return GatewaysLocationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewaysLocation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewaysLocation) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewaysLocation from the server
func (o *GatewaysLocation) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewaysLocation into the server
func (o *GatewaysLocation) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewaysLocation from the server
func (o *GatewaysLocation) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GatewaysLocation
func (o *GatewaysLocation) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GatewaysLocation
func (o *GatewaysLocation) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GatewaysLocation
func (o *GatewaysLocation) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GatewaysLocation
func (o *GatewaysLocation) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
