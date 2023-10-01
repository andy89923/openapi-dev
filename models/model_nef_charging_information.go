/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type NefChargingInformation struct {
	ExternalIndividualIdentifier string            `json:"externalIndividualIdentifier,omitempty" yaml:"externalIndividualIdentifier" bson:"externalIndividualIdentifier" mapstructure:"ExternalIndividualIdentifier"`
	ExternalGroupIdentifier      string            `json:"externalGroupIdentifier,omitempty" yaml:"externalGroupIdentifier" bson:"externalGroupIdentifier" mapstructure:"ExternalGroupIdentifier"`
	GroupIdentifier              string            `json:"groupIdentifier,omitempty" yaml:"groupIdentifier" bson:"groupIdentifier" mapstructure:"GroupIdentifier"`
	APIDirection                 ApiDirection      `json:"aPIDirection,omitempty" yaml:"aPIDirection" bson:"aPIDirection" mapstructure:"APIDirection"`
	APITargetNetworkFunction     *NfIdentification `json:"aPITargetNetworkFunction,omitempty" yaml:"aPITargetNetworkFunction" bson:"aPITargetNetworkFunction" mapstructure:"APITargetNetworkFunction"`
	APIResultCode                int32             `json:"aPIResultCode,omitempty" yaml:"aPIResultCode" bson:"aPIResultCode" mapstructure:"APIResultCode"`
	APIName                      string            `json:"aPIName" yaml:"aPIName" bson:"aPIName" mapstructure:"APIName"`
	APIReference                 string            `json:"aPIReference,omitempty" yaml:"aPIReference" bson:"aPIReference" mapstructure:"APIReference"`
	APIContent                   string            `json:"aPIContent,omitempty" yaml:"aPIContent" bson:"aPIContent" mapstructure:"APIContent"`
}