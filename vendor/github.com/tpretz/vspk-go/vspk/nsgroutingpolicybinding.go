/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// NSGRoutingPolicyBindingIdentity represents the Identity of the object
var NSGRoutingPolicyBindingIdentity = bambou.Identity{
	Name:     "nsgroutingpolicybinding",
	Category: "nsgroutingpolicybindings",
}

// NSGRoutingPolicyBindingsList represents a list of NSGRoutingPolicyBindings
type NSGRoutingPolicyBindingsList []*NSGRoutingPolicyBinding

// NSGRoutingPolicyBindingsAncestor is the interface that an ancestor of a NSGRoutingPolicyBinding must implement.
// An Ancestor is defined as an entity that has NSGRoutingPolicyBinding as a descendant.
// An Ancestor can get a list of its child NSGRoutingPolicyBindings, but not necessarily create one.
type NSGRoutingPolicyBindingsAncestor interface {
	NSGRoutingPolicyBindings(*bambou.FetchingInfo) (NSGRoutingPolicyBindingsList, *bambou.Error)
}

// NSGRoutingPolicyBindingsParent is the interface that a parent of a NSGRoutingPolicyBinding must implement.
// A Parent is defined as an entity that has NSGRoutingPolicyBinding as a child.
// A Parent is an Ancestor which can create a NSGRoutingPolicyBinding.
type NSGRoutingPolicyBindingsParent interface {
	NSGRoutingPolicyBindingsAncestor
	CreateNSGRoutingPolicyBinding(*NSGRoutingPolicyBinding) *bambou.Error
}

// NSGRoutingPolicyBinding represents the model of a nsgroutingpolicybinding
type NSGRoutingPolicyBinding struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	Name                            string `json:"name,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	Description                     string `json:"description,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	AssociatedExportRoutingPolicyID string `json:"associatedExportRoutingPolicyID,omitempty"`
	AssociatedImportRoutingPolicyID string `json:"associatedImportRoutingPolicyID,omitempty"`
	AssociatedPolicyObjectGroupID   string `json:"associatedPolicyObjectGroupID,omitempty"`
	ExportToOverlay                 string `json:"exportToOverlay,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
}

// NewNSGRoutingPolicyBinding returns a new *NSGRoutingPolicyBinding
func NewNSGRoutingPolicyBinding() *NSGRoutingPolicyBinding {

	return &NSGRoutingPolicyBinding{
		ExportToOverlay: "INHERITED",
	}
}

// Identity returns the Identity of the object.
func (o *NSGRoutingPolicyBinding) Identity() bambou.Identity {

	return NSGRoutingPolicyBindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGRoutingPolicyBinding) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGRoutingPolicyBinding) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGRoutingPolicyBinding from the server
func (o *NSGRoutingPolicyBinding) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGRoutingPolicyBinding into the server
func (o *NSGRoutingPolicyBinding) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGRoutingPolicyBinding from the server
func (o *NSGRoutingPolicyBinding) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NSGRoutingPolicyBinding
func (o *NSGRoutingPolicyBinding) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSGRoutingPolicyBinding
func (o *NSGRoutingPolicyBinding) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSGRoutingPolicyBinding
func (o *NSGRoutingPolicyBinding) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSGRoutingPolicyBinding
func (o *NSGRoutingPolicyBinding) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
