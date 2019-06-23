/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PolicyGroupCategoryIdentity represents the Identity of the object
var PolicyGroupCategoryIdentity = bambou.Identity{
	Name:     "policygroupcategory",
	Category: "policygroupcategories",
}

// PolicyGroupCategoriesList represents a list of PolicyGroupCategories
type PolicyGroupCategoriesList []*PolicyGroupCategory

// PolicyGroupCategoriesAncestor is the interface that an ancestor of a PolicyGroupCategory must implement.
// An Ancestor is defined as an entity that has PolicyGroupCategory as a descendant.
// An Ancestor can get a list of its child PolicyGroupCategories, but not necessarily create one.
type PolicyGroupCategoriesAncestor interface {
	PolicyGroupCategories(*bambou.FetchingInfo) (PolicyGroupCategoriesList, *bambou.Error)
}

// PolicyGroupCategoriesParent is the interface that a parent of a PolicyGroupCategory must implement.
// A Parent is defined as an entity that has PolicyGroupCategory as a child.
// A Parent is an Ancestor which can create a PolicyGroupCategory.
type PolicyGroupCategoriesParent interface {
	PolicyGroupCategoriesAncestor
	CreatePolicyGroupCategory(*PolicyGroupCategory) *bambou.Error
}

// PolicyGroupCategory represents the model of a policygroupcategory
type PolicyGroupCategory struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	DefaultTag    bool   `json:"defaultTag"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPolicyGroupCategory returns a new *PolicyGroupCategory
func NewPolicyGroupCategory() *PolicyGroupCategory {

	return &PolicyGroupCategory{
		DefaultTag: false,
	}
}

// Identity returns the Identity of the object.
func (o *PolicyGroupCategory) Identity() bambou.Identity {

	return PolicyGroupCategoryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyGroupCategory) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyGroupCategory) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyGroupCategory from the server
func (o *PolicyGroupCategory) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyGroupCategory into the server
func (o *PolicyGroupCategory) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyGroupCategory from the server
func (o *PolicyGroupCategory) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PolicyGroupCategory
func (o *PolicyGroupCategory) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyGroupCategory
func (o *PolicyGroupCategory) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyGroupCategory
func (o *PolicyGroupCategory) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyGroupCategory
func (o *PolicyGroupCategory) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PolicyGroups retrieves the list of child PolicyGroups of the PolicyGroupCategory
func (o *PolicyGroupCategory) PolicyGroups(info *bambou.FetchingInfo) (PolicyGroupsList, *bambou.Error) {

	var list PolicyGroupsList
	err := bambou.CurrentSession().FetchChildren(o, PolicyGroupIdentity, &list, info)
	return list, err
}
