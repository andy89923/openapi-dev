/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.521 V17.7.0; 5G System; Binding Support Management Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.521/
 *
 * API version: 1.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the requested modification to an Individual PCF for an MBS Session binding.
type PcfMbsBindingPatch struct {
	// Fully Qualified Domain Name
	PcfFqdn        string       `json:"pcfFqdn,omitempty" yaml:"pcfFqdn" bson:"pcfFqdn,omitempty"`
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty" yaml:"pcfIpEndPoints" bson:"pcfIpEndPoints,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PcfId string `json:"pcfId,omitempty" yaml:"pcfId" bson:"pcfId,omitempty"`
}