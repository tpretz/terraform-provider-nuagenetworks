/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// OverlayMirrorDestinationIdentity represents the Identity of the object
var OverlayMirrorDestinationIdentity = bambou.Identity{
	Name:     "overlaymirrordestination",
	Category: "overlaymirrordestinations",
}

// OverlayMirrorDestinationsList represents a list of OverlayMirrorDestinations
type OverlayMirrorDestinationsList []*OverlayMirrorDestination

// OverlayMirrorDestinationsAncestor is the interface that an ancestor of a OverlayMirrorDestination must implement.
// An Ancestor is defined as an entity that has OverlayMirrorDestination as a descendant.
// An Ancestor can get a list of its child OverlayMirrorDestinations, but not necessarily create one.
type OverlayMirrorDestinationsAncestor interface {
	OverlayMirrorDestinations(*bambou.FetchingInfo) (OverlayMirrorDestinationsList, *bambou.Error)
}

// OverlayMirrorDestinationsParent is the interface that a parent of a OverlayMirrorDestination must implement.
// A Parent is defined as an entity that has OverlayMirrorDestination as a child.
// A Parent is an Ancestor which can create a OverlayMirrorDestination.
type OverlayMirrorDestinationsParent interface {
	OverlayMirrorDestinationsAncestor
	CreateOverlayMirrorDestination(*OverlayMirrorDestination) *bambou.Error
}

// OverlayMirrorDestination represents the model of a overlaymirrordestination
type OverlayMirrorDestination struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	ESI               string `json:"ESI,omitempty"`
	Name              string `json:"name,omitempty"`
	LastUpdatedBy     string `json:"lastUpdatedBy,omitempty"`
	RedundancyEnabled bool   `json:"redundancyEnabled"`
	TemplateID        string `json:"templateID,omitempty"`
	Description       string `json:"description,omitempty"`
	DestinationType   string `json:"destinationType,omitempty"`
	VirtualNetworkID  string `json:"virtualNetworkID,omitempty"`
	EndPointType      string `json:"endPointType,omitempty"`
	EntityScope       string `json:"entityScope,omitempty"`
	TriggerType       string `json:"triggerType,omitempty"`
	ExternalID        string `json:"externalID,omitempty"`
}

// NewOverlayMirrorDestination returns a new *OverlayMirrorDestination
func NewOverlayMirrorDestination() *OverlayMirrorDestination {

	return &OverlayMirrorDestination{
		EndPointType: "NONE",
	}
}

// Identity returns the Identity of the object.
func (o *OverlayMirrorDestination) Identity() bambou.Identity {

	return OverlayMirrorDestinationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OverlayMirrorDestination) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OverlayMirrorDestination) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OverlayMirrorDestination from the server
func (o *OverlayMirrorDestination) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OverlayMirrorDestination into the server
func (o *OverlayMirrorDestination) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OverlayMirrorDestination from the server
func (o *OverlayMirrorDestination) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the OverlayMirrorDestination
func (o *OverlayMirrorDestination) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the OverlayMirrorDestination
func (o *OverlayMirrorDestination) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the OverlayMirrorDestination
func (o *OverlayMirrorDestination) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the OverlayMirrorDestination
func (o *OverlayMirrorDestination) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the OverlayMirrorDestination
func (o *OverlayMirrorDestination) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}

// AssignVPorts assigns the list of VPorts to the OverlayMirrorDestination
func (o *OverlayMirrorDestination) AssignVPorts(children VPortsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, VPortIdentity)
}
