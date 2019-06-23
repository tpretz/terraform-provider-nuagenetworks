/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PGExpressionTemplateIdentity represents the Identity of the object
var PGExpressionTemplateIdentity = bambou.Identity{
	Name:     "pgexpressiontemplate",
	Category: "pgexpressiontemplates",
}

// PGExpressionTemplatesList represents a list of PGExpressionTemplates
type PGExpressionTemplatesList []*PGExpressionTemplate

// PGExpressionTemplatesAncestor is the interface that an ancestor of a PGExpressionTemplate must implement.
// An Ancestor is defined as an entity that has PGExpressionTemplate as a descendant.
// An Ancestor can get a list of its child PGExpressionTemplates, but not necessarily create one.
type PGExpressionTemplatesAncestor interface {
	PGExpressionTemplates(*bambou.FetchingInfo) (PGExpressionTemplatesList, *bambou.Error)
}

// PGExpressionTemplatesParent is the interface that a parent of a PGExpressionTemplate must implement.
// A Parent is defined as an entity that has PGExpressionTemplate as a child.
// A Parent is an Ancestor which can create a PGExpressionTemplate.
type PGExpressionTemplatesParent interface {
	PGExpressionTemplatesAncestor
	CreatePGExpressionTemplate(*PGExpressionTemplate) *bambou.Error
}

// PGExpressionTemplate represents the model of a pgexpressiontemplate
type PGExpressionTemplate struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	Expression    string `json:"expression,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPGExpressionTemplate returns a new *PGExpressionTemplate
func NewPGExpressionTemplate() *PGExpressionTemplate {

	return &PGExpressionTemplate{}
}

// Identity returns the Identity of the object.
func (o *PGExpressionTemplate) Identity() bambou.Identity {

	return PGExpressionTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PGExpressionTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PGExpressionTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PGExpressionTemplate from the server
func (o *PGExpressionTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PGExpressionTemplate into the server
func (o *PGExpressionTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PGExpressionTemplate from the server
func (o *PGExpressionTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
