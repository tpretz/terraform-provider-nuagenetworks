/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NetworkPerformanceMeasurementIdentity represents the Identity of the object
var NetworkPerformanceMeasurementIdentity = bambou.Identity{
	Name:     "networkperformancemeasurement",
	Category: "networkperformancemeasurements",
}

// NetworkPerformanceMeasurementsList represents a list of NetworkPerformanceMeasurements
type NetworkPerformanceMeasurementsList []*NetworkPerformanceMeasurement

// NetworkPerformanceMeasurementsAncestor is the interface that an ancestor of a NetworkPerformanceMeasurement must implement.
// An Ancestor is defined as an entity that has NetworkPerformanceMeasurement as a descendant.
// An Ancestor can get a list of its child NetworkPerformanceMeasurements, but not necessarily create one.
type NetworkPerformanceMeasurementsAncestor interface {
	NetworkPerformanceMeasurements(*bambou.FetchingInfo) (NetworkPerformanceMeasurementsList, *bambou.Error)
}

// NetworkPerformanceMeasurementsParent is the interface that a parent of a NetworkPerformanceMeasurement must implement.
// A Parent is defined as an entity that has NetworkPerformanceMeasurement as a child.
// A Parent is an Ancestor which can create a NetworkPerformanceMeasurement.
type NetworkPerformanceMeasurementsParent interface {
	NetworkPerformanceMeasurementsAncestor
	CreateNetworkPerformanceMeasurement(*NetworkPerformanceMeasurement) *bambou.Error
}

// NetworkPerformanceMeasurement represents the model of a networkperformancemeasurement
type NetworkPerformanceMeasurement struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	NPMType                        string `json:"NPMType,omitempty"`
	Name                           string `json:"name,omitempty"`
	LastUpdatedBy                  string `json:"lastUpdatedBy,omitempty"`
	ReadOnly                       bool   `json:"readOnly"`
	Description                    string `json:"description,omitempty"`
	EntityScope                    string `json:"entityScope,omitempty"`
	AssociatedPerformanceMonitorID string `json:"associatedPerformanceMonitorID,omitempty"`
	ExternalID                     string `json:"externalID,omitempty"`
}

// NewNetworkPerformanceMeasurement returns a new *NetworkPerformanceMeasurement
func NewNetworkPerformanceMeasurement() *NetworkPerformanceMeasurement {

	return &NetworkPerformanceMeasurement{
		NPMType:  "NONE",
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *NetworkPerformanceMeasurement) Identity() bambou.Identity {

	return NetworkPerformanceMeasurementIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkPerformanceMeasurement) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkPerformanceMeasurement) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetworkPerformanceMeasurement from the server
func (o *NetworkPerformanceMeasurement) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetworkPerformanceMeasurement into the server
func (o *NetworkPerformanceMeasurement) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetworkPerformanceMeasurement from the server
func (o *NetworkPerformanceMeasurement) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NetworkPerformanceBindings retrieves the list of child NetworkPerformanceBindings of the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) NetworkPerformanceBindings(info *bambou.FetchingInfo) (NetworkPerformanceBindingsList, *bambou.Error) {

	var list NetworkPerformanceBindingsList
	err := bambou.CurrentSession().FetchChildren(o, NetworkPerformanceBindingIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Monitorscopes retrieves the list of child Monitorscopes of the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) Monitorscopes(info *bambou.FetchingInfo) (MonitorscopesList, *bambou.Error) {

	var list MonitorscopesList
	err := bambou.CurrentSession().FetchChildren(o, MonitorscopeIdentity, &list, info)
	return list, err
}

// CreateMonitorscope creates a new child Monitorscope under the NetworkPerformanceMeasurement
func (o *NetworkPerformanceMeasurement) CreateMonitorscope(child *Monitorscope) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
