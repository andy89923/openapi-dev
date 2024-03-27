/*
 * Nnef_PFDmanagement Service API
 *
 * Packet Flow Description Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.551 v17.9.0, 5G System; Packet Flow Description Management Service
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.551/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Represents the PFDs for an application identifier.
type PfdDataForApp struct {
	// String providing an application identifier.
	ApplicationId string       `json:"applicationId" yaml:"applicationId" bson:"applicationId,omitempty"`
	Pfds          []PfdContent `json:"pfds,omitempty" yaml:"pfds" bson:"pfds,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	CachingTime *time.Time `json:"cachingTime,omitempty" yaml:"cachingTime" bson:"cachingTime,omitempty"`
	// indicating a time in seconds.
	CachingTimer int32 `json:"cachingTimer,omitempty" yaml:"cachingTimer" bson:"cachingTimer,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	PfdTimestamp *time.Time `json:"pfdTimestamp,omitempty" yaml:"pfdTimestamp" bson:"pfdTimestamp,omitempty"`
	// Indication of partial update of PFDs for an existing application identifier if it is included and set to true. Otherwise set to false indicates not supporting partial update of PFDs for an existing application identifier. The default value false applies if the attribute is not present.
	PartialFlag bool `json:"partialFlag,omitempty" yaml:"partialFlag" bson:"partialFlag,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
}