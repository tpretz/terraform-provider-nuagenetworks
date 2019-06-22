/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CommandIdentity represents the Identity of the object
var CommandIdentity = bambou.Identity{
	Name:     "command",
	Category: "commands",
}

// CommandsList represents a list of Commands
type CommandsList []*Command

// CommandsAncestor is the interface that an ancestor of a Command must implement.
// An Ancestor is defined as an entity that has Command as a descendant.
// An Ancestor can get a list of its child Commands, but not necessarily create one.
type CommandsAncestor interface {
	Commands(*bambou.FetchingInfo) (CommandsList, *bambou.Error)
}

// CommandsParent is the interface that a parent of a Command must implement.
// A Parent is defined as an entity that has Command as a child.
// A Parent is an Ancestor which can create a Command.
type CommandsParent interface {
	CommandsAncestor
	CreateCommand(*Command) *bambou.Error
}

// Command represents the model of a command
type Command struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	Detail              string `json:"detail,omitempty"`
	DetailedStatus      string `json:"detailedStatus,omitempty"`
	DetailedStatusCode  int    `json:"detailedStatusCode,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	Command             string `json:"command,omitempty"`
	CommandInformation  string `json:"commandInformation,omitempty"`
	Progress            string `json:"progress,omitempty"`
	AssocEntityType     string `json:"assocEntityType,omitempty"`
	AssociatedParam     string `json:"associatedParam,omitempty"`
	AssociatedParamType string `json:"associatedParamType,omitempty"`
	Status              string `json:"status,omitempty"`
	FullCommand         string `json:"fullCommand,omitempty"`
	Summary             string `json:"summary,omitempty"`
	Override            string `json:"override,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

// NewCommand returns a new *Command
func NewCommand() *Command {

	return &Command{
		DetailedStatusCode: 0,
		Command:            "UNKNOWN",
		Status:             "UNKNOWN",
		Override:           "UNSPECIFIED",
	}
}

// Identity returns the Identity of the object.
func (o *Command) Identity() bambou.Identity {

	return CommandIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Command) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Command) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Command from the server
func (o *Command) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Command into the server
func (o *Command) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Command from the server
func (o *Command) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
