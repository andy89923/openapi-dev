/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NefChargingInformation struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	ExternalIndividualIdentifier string   `json:"externalIndividualIdentifier,omitempty" yaml:"externalIndividualIdentifier" bson:"externalIndividualIdentifier,omitempty"`
	ExternalIndividualIdList     []string `json:"externalIndividualIdList,omitempty" yaml:"externalIndividualIdList" bson:"externalIndividualIdList,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.
	ExternalGroupIdentifier string `json:"externalGroupIdentifier,omitempty" yaml:"externalGroupIdentifier" bson:"externalGroupIdentifier,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.
	GroupIdentifier          string                                `json:"groupIdentifier,omitempty" yaml:"groupIdentifier" bson:"groupIdentifier,omitempty"`
	APIDirection             ApiDirection                          `json:"aPIDirection,omitempty" yaml:"aPIDirection" bson:"aPIDirection,omitempty"`
	APITargetNetworkFunction *ChfConvergedChargingNfIdentification `json:"aPITargetNetworkFunction,omitempty" yaml:"aPITargetNetworkFunction" bson:"aPITargetNetworkFunction,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	APIResultCode int32  `json:"aPIResultCode,omitempty" yaml:"aPIResultCode" bson:"aPIResultCode,omitempty"`
	APIName       string `json:"aPIName" yaml:"aPIName" bson:"aPIName,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	APIReference string `json:"aPIReference,omitempty" yaml:"aPIReference" bson:"aPIReference,omitempty"`
	APIContent   string `json:"aPIContent,omitempty" yaml:"aPIContent" bson:"aPIContent,omitempty"`
}