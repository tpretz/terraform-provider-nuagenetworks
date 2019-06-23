/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// JobIdentity represents the Identity of the object
var JobIdentity = bambou.Identity{
	Name:     "job",
	Category: "jobs",
}

// JobsList represents a list of Jobs
type JobsList []*Job

// JobsAncestor is the interface that an ancestor of a Job must implement.
// An Ancestor is defined as an entity that has Job as a descendant.
// An Ancestor can get a list of its child Jobs, but not necessarily create one.
type JobsAncestor interface {
	Jobs(*bambou.FetchingInfo) (JobsList, *bambou.Error)
}

// JobsParent is the interface that a parent of a Job must implement.
// A Parent is defined as an entity that has Job as a child.
// A Parent is an Ancestor which can create a Job.
type JobsParent interface {
	JobsAncestor
	CreateJob(*Job) *bambou.Error
}

// Job represents the model of a job
type Job struct {
	ID              string      `json:"ID,omitempty"`
	ParentID        string      `json:"parentID,omitempty"`
	ParentType      string      `json:"parentType,omitempty"`
	Owner           string      `json:"owner,omitempty"`
	Parameters      interface{} `json:"parameters,omitempty"`
	LastUpdatedBy   string      `json:"lastUpdatedBy,omitempty"`
	Result          interface{} `json:"result,omitempty"`
	EntityScope     string      `json:"entityScope,omitempty"`
	Command         string      `json:"command,omitempty"`
	Progress        float64     `json:"progress,omitempty"`
	AssocEntityType string      `json:"assocEntityType,omitempty"`
	Status          string      `json:"status,omitempty"`
	ExternalID      string      `json:"externalID,omitempty"`
}

// NewJob returns a new *Job
func NewJob() *Job {

	return &Job{}
}

// Identity returns the Identity of the object.
func (o *Job) Identity() bambou.Identity {

	return JobIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Job) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Job) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Job from the server
func (o *Job) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Job into the server
func (o *Job) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Job from the server
func (o *Job) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Job
func (o *Job) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Job
func (o *Job) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Job
func (o *Job) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Job
func (o *Job) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
