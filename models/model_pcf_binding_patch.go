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

// Identifies an Individual PCF binding used in an HTTP Patch method.
type PcfBindingPatch struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166 with the OpenAPI defined 'nullable: true' property.
	Ipv4Addr   string `json:"ipv4Addr,omitempty" yaml:"ipv4Addr" bson:"ipv4Addr,omitempty"`
	IpDomain   string `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty" yaml:"ipv6Prefix" bson:"ipv6Prefix,omitempty"`
	// The additional IPv6 Address Prefixes of the served UE.
	AddIpv6Prefixes []string `json:"addIpv6Prefixes,omitempty" yaml:"addIpv6Prefixes" bson:"addIpv6Prefixes,omitempty"`
	// \"String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042 with the OpenAPI 'nullable: true' property.\"
	MacAddr48 string `json:"macAddr48,omitempty" yaml:"macAddr48" bson:"macAddr48,omitempty"`
	// The additional MAC Addresses of the served UE.
	AddMacAddrs []string `json:"addMacAddrs,omitempty" yaml:"addMacAddrs" bson:"addMacAddrs,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PcfId string `json:"pcfId,omitempty" yaml:"pcfId" bson:"pcfId,omitempty"`
	// Fully Qualified Domain Name
	PcfFqdn string `json:"pcfFqdn,omitempty" yaml:"pcfFqdn" bson:"pcfFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_PolicyAuthorization service.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty" yaml:"pcfIpEndPoints" bson:"pcfIpEndPoints,omitempty"`
	// Fully Qualified Domain Name
	PcfDiamHost string `json:"pcfDiamHost,omitempty" yaml:"pcfDiamHost" bson:"pcfDiamHost,omitempty"`
	// Fully Qualified Domain Name
	PcfDiamRealm string `json:"pcfDiamRealm,omitempty" yaml:"pcfDiamRealm" bson:"pcfDiamRealm,omitempty"`
}