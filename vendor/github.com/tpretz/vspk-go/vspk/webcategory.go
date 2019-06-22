/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// WebCategoryIdentity represents the Identity of the object
var WebCategoryIdentity = bambou.Identity{
	Name:     "webcategory",
	Category: "webcategories",
}

// WebCategoriesList represents a list of WebCategories
type WebCategoriesList []*WebCategory

// WebCategoriesAncestor is the interface that an ancestor of a WebCategory must implement.
// An Ancestor is defined as an entity that has WebCategory as a descendant.
// An Ancestor can get a list of its child WebCategories, but not necessarily create one.
type WebCategoriesAncestor interface {
	WebCategories(*bambou.FetchingInfo) (WebCategoriesList, *bambou.Error)
}

// WebCategoriesParent is the interface that a parent of a WebCategory must implement.
// A Parent is defined as an entity that has WebCategory as a child.
// A Parent is an Ancestor which can create a WebCategory.
type WebCategoriesParent interface {
	WebCategoriesAncestor
	CreateWebCategory(*WebCategory) *bambou.Error
}

// WebCategory represents the model of a webcategory
type WebCategory struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	Name            string `json:"name,omitempty"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	DefaultCategory bool   `json:"defaultCategory"`
	Description     string `json:"description,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
	Type            string `json:"type,omitempty"`
}

// NewWebCategory returns a new *WebCategory
func NewWebCategory() *WebCategory {

	return &WebCategory{
		DefaultCategory: false,
		Type:            "WEB_DOMAIN_NAME",
	}
}

// Identity returns the Identity of the object.
func (o *WebCategory) Identity() bambou.Identity {

	return WebCategoryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *WebCategory) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *WebCategory) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the WebCategory from the server
func (o *WebCategory) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the WebCategory into the server
func (o *WebCategory) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the WebCategory from the server
func (o *WebCategory) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// WebDomainNames retrieves the list of child WebDomainNames of the WebCategory
func (o *WebCategory) WebDomainNames(info *bambou.FetchingInfo) (WebDomainNamesList, *bambou.Error) {

	var list WebDomainNamesList
	err := bambou.CurrentSession().FetchChildren(o, WebDomainNameIdentity, &list, info)
	return list, err
}

// AssignWebDomainNames assigns the list of WebDomainNames to the WebCategory
func (o *WebCategory) AssignWebDomainNames(children WebDomainNamesList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, WebDomainNameIdentity)
}

// Metadatas retrieves the list of child Metadatas of the WebCategory
func (o *WebCategory) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the WebCategory
func (o *WebCategory) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the WebCategory
func (o *WebCategory) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the WebCategory
func (o *WebCategory) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
