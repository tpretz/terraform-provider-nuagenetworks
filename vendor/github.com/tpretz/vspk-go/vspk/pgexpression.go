/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PGExpressionIdentity represents the Identity of the object
var PGExpressionIdentity = bambou.Identity{
	Name:     "pgexpression",
	Category: "pgexpressions",
}

// PGExpressionsList represents a list of PGExpressions
type PGExpressionsList []*PGExpression

// PGExpressionsAncestor is the interface that an ancestor of a PGExpression must implement.
// An Ancestor is defined as an entity that has PGExpression as a descendant.
// An Ancestor can get a list of its child PGExpressions, but not necessarily create one.
type PGExpressionsAncestor interface {
	PGExpressions(*bambou.FetchingInfo) (PGExpressionsList, *bambou.Error)
}

// PGExpressionsParent is the interface that a parent of a PGExpression must implement.
// A Parent is defined as an entity that has PGExpression as a child.
// A Parent is an Ancestor which can create a PGExpression.
type PGExpressionsParent interface {
	PGExpressionsAncestor
	CreatePGExpression(*PGExpression) *bambou.Error
}

// PGExpression represents the model of a pgexpression
type PGExpression struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	TemplateID    string `json:"templateID,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	Expression    string `json:"expression,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPGExpression returns a new *PGExpression
func NewPGExpression() *PGExpression {

	return &PGExpression{}
}

// Identity returns the Identity of the object.
func (o *PGExpression) Identity() bambou.Identity {

	return PGExpressionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PGExpression) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PGExpression) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PGExpression from the server
func (o *PGExpression) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PGExpression into the server
func (o *PGExpression) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PGExpression from the server
func (o *PGExpression) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
