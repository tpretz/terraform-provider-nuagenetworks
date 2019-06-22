/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FirewallRuleIdentity represents the Identity of the object
var FirewallRuleIdentity = bambou.Identity{
	Name:     "firewallrule",
	Category: "firewallrules",
}

// FirewallRulesList represents a list of FirewallRules
type FirewallRulesList []*FirewallRule

// FirewallRulesAncestor is the interface that an ancestor of a FirewallRule must implement.
// An Ancestor is defined as an entity that has FirewallRule as a descendant.
// An Ancestor can get a list of its child FirewallRules, but not necessarily create one.
type FirewallRulesAncestor interface {
	FirewallRules(*bambou.FetchingInfo) (FirewallRulesList, *bambou.Error)
}

// FirewallRulesParent is the interface that a parent of a FirewallRule must implement.
// A Parent is defined as an entity that has FirewallRule as a child.
// A Parent is an Ancestor which can create a FirewallRule.
type FirewallRulesParent interface {
	FirewallRulesAncestor
	CreateFirewallRule(*FirewallRule) *bambou.Error
}

// FirewallRule represents the model of a firewallrule
type FirewallRule struct {
	ID                       string `json:"ID,omitempty"`
	ParentID                 string `json:"parentID,omitempty"`
	ParentType               string `json:"parentType,omitempty"`
	Owner                    string `json:"owner,omitempty"`
	ACLTemplateName          string `json:"ACLTemplateName,omitempty"`
	ICMPCode                 string `json:"ICMPCode,omitempty"`
	ICMPType                 string `json:"ICMPType,omitempty"`
	IPv6AddressOverride      string `json:"IPv6AddressOverride,omitempty"`
	DSCP                     string `json:"DSCP,omitempty"`
	LastUpdatedBy            string `json:"lastUpdatedBy,omitempty"`
	Action                   string `json:"action,omitempty"`
	AddressOverride          string `json:"addressOverride,omitempty"`
	WebFilterID              string `json:"webFilterID,omitempty"`
	WebFilterType            string `json:"webFilterType,omitempty"`
	Description              string `json:"description,omitempty"`
	DestinationPort          string `json:"destinationPort,omitempty"`
	NetworkID                string `json:"networkID,omitempty"`
	NetworkType              string `json:"networkType,omitempty"`
	MirrorDestinationID      string `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled       bool   `json:"flowLoggingEnabled"`
	EnterpriseName           string `json:"enterpriseName,omitempty"`
	EntityScope              string `json:"entityScope,omitempty"`
	LocationID               string `json:"locationID,omitempty"`
	LocationType             string `json:"locationType,omitempty"`
	DomainName               string `json:"domainName,omitempty"`
	SourcePort               string `json:"sourcePort,omitempty"`
	Priority                 int    `json:"priority,omitempty"`
	Protocol                 string `json:"protocol,omitempty"`
	AssociatedLiveTemplateID string `json:"associatedLiveTemplateID,omitempty"`
	AssociatedTrafficType    string `json:"associatedTrafficType,omitempty"`
	AssociatedTrafficTypeID  string `json:"associatedTrafficTypeID,omitempty"`
	AssociatedfirewallACLID  string `json:"associatedfirewallACLID,omitempty"`
	Stateful                 bool   `json:"stateful"`
	StatsID                  string `json:"statsID,omitempty"`
	StatsLoggingEnabled      bool   `json:"statsLoggingEnabled"`
	EtherType                string `json:"etherType,omitempty"`
	ExternalID               string `json:"externalID,omitempty"`
}

// NewFirewallRule returns a new *FirewallRule
func NewFirewallRule() *FirewallRule {

	return &FirewallRule{
		Stateful: false,
	}
}

// Identity returns the Identity of the object.
func (o *FirewallRule) Identity() bambou.Identity {

	return FirewallRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FirewallRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FirewallRule) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FirewallRule from the server
func (o *FirewallRule) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FirewallRule into the server
func (o *FirewallRule) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FirewallRule from the server
func (o *FirewallRule) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the FirewallRule
func (o *FirewallRule) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the FirewallRule
func (o *FirewallRule) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the FirewallRule
func (o *FirewallRule) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the FirewallRule
func (o *FirewallRule) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
