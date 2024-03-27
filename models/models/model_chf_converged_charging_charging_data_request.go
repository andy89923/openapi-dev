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

import (
	"time"
)

type ChfConvergedChargingChargingDataRequest struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	SubscriberIdentifier string `json:"subscriberIdentifier,omitempty" yaml:"subscriberIdentifier" bson:"subscriberIdentifier,omitempty"`
	TenantIdentifier     string `json:"tenantIdentifier,omitempty" yaml:"tenantIdentifier" bson:"tenantIdentifier,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	ChargingId               int32                                 `json:"chargingId,omitempty" yaml:"chargingId" bson:"chargingId,omitempty"`
	MnSConsumerIdentifier    string                                `json:"mnSConsumerIdentifier,omitempty" yaml:"mnSConsumerIdentifier" bson:"mnSConsumerIdentifier,omitempty"`
	NfConsumerIdentification *ChfConvergedChargingNfIdentification `json:"nfConsumerIdentification" yaml:"nfConsumerIdentification" bson:"nfConsumerIdentification,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	InvocationTimeStamp *time.Time `json:"invocationTimeStamp" yaml:"invocationTimeStamp" bson:"invocationTimeStamp,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	InvocationSequenceNumber int32            `json:"invocationSequenceNumber" yaml:"invocationSequenceNumber" bson:"invocationSequenceNumber,omitempty"`
	RetransmissionIndicator  bool             `json:"retransmissionIndicator,omitempty" yaml:"retransmissionIndicator" bson:"retransmissionIndicator,omitempty"`
	OneTimeEvent             bool             `json:"oneTimeEvent,omitempty" yaml:"oneTimeEvent" bson:"oneTimeEvent,omitempty"`
	OneTimeEventType         OneTimeEventType `json:"oneTimeEventType,omitempty" yaml:"oneTimeEventType" bson:"oneTimeEventType,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri,omitempty" yaml:"notifyUri" bson:"notifyUri,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures        string                                  `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	ServiceSpecificationInfo string                                  `json:"serviceSpecificationInfo,omitempty" yaml:"serviceSpecificationInfo" bson:"serviceSpecificationInfo,omitempty"`
	MultipleUnitUsage        []ChfConvergedChargingMultipleUnitUsage `json:"multipleUnitUsage,omitempty" yaml:"multipleUnitUsage" bson:"multipleUnitUsage,omitempty"`
	Triggers                 []ChfConvergedChargingTrigger           `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	Easid                    string                                  `json:"easid,omitempty" yaml:"easid" bson:"easid,omitempty"`
	Ednid                    string                                  `json:"ednid,omitempty" yaml:"ednid" bson:"ednid,omitempty"`
	EASProviderIdentifier    string                                  `json:"eASProviderIdentifier,omitempty" yaml:"eASProviderIdentifier" bson:"eASProviderIdentifier,omitempty"`
	// String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).
	AMFId                                         string                                             `json:"aMFId,omitempty" yaml:"aMFId" bson:"aMFId,omitempty"`
	PDUSessionChargingInformation                 *ChfConvergedChargingPduSessionChargingInformation `json:"pDUSessionChargingInformation,omitempty" yaml:"pDUSessionChargingInformation" bson:"pDUSessionChargingInformation,omitempty"`
	RoamingQBCInformation                         *ChfConvergedChargingRoamingQbcInformation         `json:"roamingQBCInformation,omitempty" yaml:"roamingQBCInformation" bson:"roamingQBCInformation,omitempty"`
	SMSChargingInformation                        *SmsChargingInformation                            `json:"sMSChargingInformation,omitempty" yaml:"sMSChargingInformation" bson:"sMSChargingInformation,omitempty"`
	NEFChargingInformation                        *NefChargingInformation                            `json:"nEFChargingInformation,omitempty" yaml:"nEFChargingInformation" bson:"nEFChargingInformation,omitempty"`
	RegistrationChargingInformation               *RegistrationChargingInformation                   `json:"registrationChargingInformation,omitempty" yaml:"registrationChargingInformation" bson:"registrationChargingInformation,omitempty"`
	N2ConnectionChargingInformation               *N2ConnectionChargingInformation                   `json:"n2ConnectionChargingInformation,omitempty" yaml:"n2ConnectionChargingInformation" bson:"n2ConnectionChargingInformation,omitempty"`
	LocationReportingChargingInformation          *LocationReportingChargingInformation              `json:"locationReportingChargingInformation,omitempty" yaml:"locationReportingChargingInformation" bson:"locationReportingChargingInformation,omitempty"`
	NSPAChargingInformation                       *NspaChargingInformation                           `json:"nSPAChargingInformation,omitempty" yaml:"nSPAChargingInformation" bson:"nSPAChargingInformation,omitempty"`
	NSMChargingInformation                        *NsmChargingInformation                            `json:"nSMChargingInformation,omitempty" yaml:"nSMChargingInformation" bson:"nSMChargingInformation,omitempty"`
	MMTelChargingInformation                      *MmTelChargingInformation                          `json:"mMTelChargingInformation,omitempty" yaml:"mMTelChargingInformation" bson:"mMTelChargingInformation,omitempty"`
	IMSChargingInformation                        *ImsChargingInformation                            `json:"iMSChargingInformation,omitempty" yaml:"iMSChargingInformation" bson:"iMSChargingInformation,omitempty"`
	EdgeInfrastructureUsageChargingInformation    *EdgeInfrastructureUsageChargingInformation        `json:"edgeInfrastructureUsageChargingInformation,omitempty" yaml:"edgeInfrastructureUsageChargingInformation'" bson:"edgeInfrastructureUsageChargingInformation',omitempty"`
	EASDeploymentChargingInformation              *EasDeploymentChargingInformation                  `json:"eASDeploymentChargingInformation,omitempty" yaml:"eASDeploymentChargingInformation" bson:"eASDeploymentChargingInformation,omitempty"`
	DirectEdgeEnablingServiceChargingInformation  *NefChargingInformation                            `json:"directEdgeEnablingServiceChargingInformation,omitempty" yaml:"directEdgeEnablingServiceChargingInformation" bson:"directEdgeEnablingServiceChargingInformation,omitempty"`
	ExposedEdgeEnablingServiceChargingInformation *NefChargingInformation                            `json:"exposedEdgeEnablingServiceChargingInformation,omitempty" yaml:"exposedEdgeEnablingServiceChargingInformation" bson:"exposedEdgeEnablingServiceChargingInformation,omitempty"`
	ProSeChargingInformation                      *ProseChargingInformation                          `json:"proSeChargingInformation,omitempty" yaml:"proSeChargingInformation" bson:"proSeChargingInformation,omitempty"`
}