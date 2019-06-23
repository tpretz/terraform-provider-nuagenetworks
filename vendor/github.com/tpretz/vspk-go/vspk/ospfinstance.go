/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// OSPFInstanceIdentity represents the Identity of the object
var OSPFInstanceIdentity = bambou.Identity{
	Name:     "ospfinstance",
	Category: "ospfinstances",
}

// OSPFInstancesList represents a list of OSPFInstances
type OSPFInstancesList []*OSPFInstance

// OSPFInstancesAncestor is the interface that an ancestor of a OSPFInstance must implement.
// An Ancestor is defined as an entity that has OSPFInstance as a descendant.
// An Ancestor can get a list of its child OSPFInstances, but not necessarily create one.
type OSPFInstancesAncestor interface {
	OSPFInstances(*bambou.FetchingInfo) (OSPFInstancesList, *bambou.Error)
}

// OSPFInstancesParent is the interface that a parent of a OSPFInstance must implement.
// A Parent is defined as an entity that has OSPFInstance as a child.
// A Parent is an Ancestor which can create a OSPFInstance.
type OSPFInstancesParent interface {
	OSPFInstancesAncestor
	CreateOSPFInstance(*OSPFInstance) *bambou.Error
}

// OSPFInstance represents the model of a ospfinstance
type OSPFInstance struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	Name                            string `json:"name,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	Description                     string `json:"description,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	Preference                      int    `json:"preference,omitempty"`
	AssociatedExportRoutingPolicyID string `json:"associatedExportRoutingPolicyID,omitempty"`
	AssociatedImportRoutingPolicyID string `json:"associatedImportRoutingPolicyID,omitempty"`
	SuperBackboneEnabled            bool   `json:"superBackboneEnabled"`
	ExportLimit                     int    `json:"exportLimit,omitempty"`
	ExportToOverlay                 bool   `json:"exportToOverlay"`
	ExternalID                      string `json:"externalID,omitempty"`
	ExternalPreference              int    `json:"externalPreference,omitempty"`
}

// NewOSPFInstance returns a new *OSPFInstance
func NewOSPFInstance() *OSPFInstance {

	return &OSPFInstance{
		Preference:           10,
		SuperBackboneEnabled: false,
		ExportToOverlay:      false,
		ExternalPreference:   150,
	}
}

// Identity returns the Identity of the object.
func (o *OSPFInstance) Identity() bambou.Identity {

	return OSPFInstanceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OSPFInstance) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OSPFInstance) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OSPFInstance from the server
func (o *OSPFInstance) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OSPFInstance into the server
func (o *OSPFInstance) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OSPFInstance from the server
func (o *OSPFInstance) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the OSPFInstance
func (o *OSPFInstance) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the OSPFInstance
func (o *OSPFInstance) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the OSPFInstance
func (o *OSPFInstance) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the OSPFInstance
func (o *OSPFInstance) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// OSPFAreas retrieves the list of child OSPFAreas of the OSPFInstance
func (o *OSPFInstance) OSPFAreas(info *bambou.FetchingInfo) (OSPFAreasList, *bambou.Error) {

	var list OSPFAreasList
	err := bambou.CurrentSession().FetchChildren(o, OSPFAreaIdentity, &list, info)
	return list, err
}

// CreateOSPFArea creates a new child OSPFArea under the OSPFInstance
func (o *OSPFInstance) CreateOSPFArea(child *OSPFArea) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
