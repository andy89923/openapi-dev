/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the route selector parameters (PDU session types, SSC modes and ATSSS information) per DNN
type DnnRouteSelectionDescriptor struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn          string           `json:"dnn" yaml:"dnn" bson:"dnn,omitempty"`
	SscModes     []SscMode        `json:"sscModes,omitempty" yaml:"sscModes" bson:"sscModes,omitempty"`
	PduSessTypes []PduSessionType `json:"pduSessTypes,omitempty" yaml:"pduSessTypes" bson:"pduSessTypes,omitempty"`
	// Indicates whether MA PDU session establishment is allowed for this DNN. When set to value true MA PDU session establishment is allowed for this DNN.
	AtsssInfo bool `json:"atsssInfo,omitempty" yaml:"atsssInfo" bson:"atsssInfo,omitempty"`
}