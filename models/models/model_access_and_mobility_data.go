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

import (
	"time"
)

// Represents Access and Mobility data for a UE.
type AccessAndMobilityData struct {
	Location *UserLocation `json:"location,omitempty" yaml:"location" bson:"location,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LocationTs *time.Time `json:"locationTs,omitempty" yaml:"locationTs" bson:"locationTs,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone" bson:"timeZone,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeZoneTs *time.Time `json:"timeZoneTs,omitempty" yaml:"timeZoneTs" bson:"timeZoneTs,omitempty"`
	AccessType AccessType `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	RegStates  []RmInfo   `json:"regStates,omitempty" yaml:"regStates" bson:"regStates,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RegStatesTs *time.Time `json:"regStatesTs,omitempty" yaml:"regStatesTs" bson:"regStatesTs,omitempty"`
	ConnStates  []CmInfo   `json:"connStates,omitempty" yaml:"connStates" bson:"connStates,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ConnStatesTs       *time.Time     `json:"connStatesTs,omitempty" yaml:"connStatesTs" bson:"connStatesTs,omitempty"`
	ReachabilityStatus UeReachability `json:"reachabilityStatus,omitempty" yaml:"reachabilityStatus" bson:"reachabilityStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ReachabilityStatusTs *time.Time `json:"reachabilityStatusTs,omitempty" yaml:"reachabilityStatusTs" bson:"reachabilityStatusTs,omitempty"`
	SmsOverNasStatus     SmsSupport `json:"smsOverNasStatus,omitempty" yaml:"smsOverNasStatus" bson:"smsOverNasStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	SmsOverNasStatusTs *time.Time `json:"smsOverNasStatusTs,omitempty" yaml:"smsOverNasStatusTs" bson:"smsOverNasStatusTs,omitempty"`
	// True  The serving PLMN of the UE is different from the HPLMN of the UE; False The serving PLMN of the UE is the HPLMN of the UE.
	RoamingStatus bool `json:"roamingStatus,omitempty" yaml:"roamingStatus" bson:"roamingStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RoamingStatusTs *time.Time `json:"roamingStatusTs,omitempty" yaml:"roamingStatusTs" bson:"roamingStatusTs,omitempty"`
	CurrentPlmn     *PlmnId    `json:"currentPlmn,omitempty" yaml:"currentPlmn" bson:"currentPlmn,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	CurrentPlmnTs *time.Time `json:"currentPlmnTs,omitempty" yaml:"currentPlmnTs" bson:"currentPlmnTs,omitempty"`
	RatType       []RatType  `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RatTypesTs *time.Time `json:"ratTypesTs,omitempty" yaml:"ratTypesTs" bson:"ratTypesTs,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string   `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat,omitempty"`
	ResetIds []string `json:"resetIds,omitempty" yaml:"resetIds" bson:"resetIds,omitempty"`
}