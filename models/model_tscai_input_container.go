/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V17.9.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Indicates TSC Traffic pattern.
type TscaiInputContainer struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Periodicity int32 `json:"periodicity,omitempty" yaml:"periodicity" bson:"periodicity,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	BurstArrivalTime *time.Time `json:"burstArrivalTime,omitempty" yaml:"burstArrivalTime" bson:"burstArrivalTime,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInNumMsg int32 `json:"surTimeInNumMsg,omitempty" yaml:"surTimeInNumMsg" bson:"surTimeInNumMsg,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInTime int32 `json:"surTimeInTime,omitempty" yaml:"surTimeInTime" bson:"surTimeInTime,omitempty"`
}