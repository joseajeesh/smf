/*
 * Nnef_PFDmanagement Sevice API
 *
 * Packet Flow Description Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PfdContent struct {
	// Identifies a PDF of an application identifier.
	PfdId string `json:"pfdId,omitempty" yaml:"pfdId" bson:"pfdId" mapstructure:"PfdId"`
	// Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic.
	FlowDescriptions []string `json:"flowDescriptions,omitempty" yaml:"flowDescriptions" bson:"flowDescriptions" mapstructure:"FlowDescriptions"`
	// Indicates a URL or a regular expression which is used to match the significant parts of the URL.
	Urls []string `json:"urls,omitempty" yaml:"urls" bson:"urls" mapstructure:"Urls"`
	// Indicates an FQDN or a regular expression as a domain name matching criteria.
	DomainNames []string `json:"domainNames,omitempty" yaml:"domainNames" bson:"domainNames" mapstructure:"DomainNames"`
}