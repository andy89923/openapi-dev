/*
 * Nnef_Authentication
 *
 * NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.256 V17.3.0; 5G System;Uncrewed Aerial Systems Network Function (UAS-NF); Aerial Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.256/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// UAV related notification
type AuthNotification struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi           string                           `json:"gpsi" yaml:"gpsi" bson:"gpsi,omitempty"`
	ServiceLevelId string                           `json:"serviceLevelId" yaml:"serviceLevelId" bson:"serviceLevelId,omitempty"`
	NotifyCorrId   string                           `json:"notifyCorrId" yaml:"notifyCorrId" bson:"notifyCorrId,omitempty"`
	AuthMsg        *RefToBinaryData                 `json:"authMsg,omitempty" yaml:"authMsg" bson:"authMsg,omitempty"`
	AuthContainer  []NefAuthenticationAuthContainer `json:"authContainer,omitempty" yaml:"authContainer" bson:"authContainer,omitempty"`
	NotifType      NotifType                        `json:"notifType" yaml:"notifType" bson:"notifType,omitempty"`
}