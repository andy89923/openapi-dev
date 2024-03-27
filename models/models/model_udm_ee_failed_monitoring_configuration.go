/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.10.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the event type and failed cause of the failed Monitoring Configuration in the EE subscription
type UdmEeFailedMonitoringConfiguration struct {
	EventType   UdmEeEventType   `json:"eventType" yaml:"eventType" bson:"eventType,omitempty"`
	FailedCause UdmEeFailedCause `json:"failedCause" yaml:"failedCause" bson:"failedCause,omitempty"`
}