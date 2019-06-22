/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CustomPropertyIdentity represents the Identity of the object
var CustomPropertyIdentity = bambou.Identity{
	Name:     "customproperty",
	Category: "customproperties",
}

// CustomPropertiesList represents a list of CustomProperties
type CustomPropertiesList []*CustomProperty

// CustomPropertiesAncestor is the interface that an ancestor of a CustomProperty must implement.
// An Ancestor is defined as an entity that has CustomProperty as a descendant.
// An Ancestor can get a list of its child CustomProperties, but not necessarily create one.
type CustomPropertiesAncestor interface {
	CustomProperties(*bambou.FetchingInfo) (CustomPropertiesList, *bambou.Error)
}

// CustomPropertiesParent is the interface that a parent of a CustomProperty must implement.
// A Parent is defined as an entity that has CustomProperty as a child.
// A Parent is an Ancestor which can create a CustomProperty.
type CustomPropertiesParent interface {
	CustomPropertiesAncestor
	CreateCustomProperty(*CustomProperty) *bambou.Error
}

// CustomProperty represents the model of a customproperty
type CustomProperty struct {
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	AttributeName  string `json:"attributeName,omitempty"`
	AttributeValue string `json:"attributeValue,omitempty"`
}

// NewCustomProperty returns a new *CustomProperty
func NewCustomProperty() *CustomProperty {

	return &CustomProperty{}
}

// Identity returns the Identity of the object.
func (o *CustomProperty) Identity() bambou.Identity {

	return CustomPropertyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CustomProperty) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CustomProperty) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CustomProperty from the server
func (o *CustomProperty) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CustomProperty into the server
func (o *CustomProperty) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CustomProperty from the server
func (o *CustomProperty) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
