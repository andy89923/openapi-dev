/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.510 V16.8.0; 5G System; Network Function Repository Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.510/
 *
 * API version: 1.1.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import "github.com/golang-jwt/jwt/v5"

// The claims data structure for the access token
type AccessTokenClaims struct {
	Iss                string      `json:"iss" yaml:"iss" bson:"iss,omitempty"`
	Sub                string      `json:"sub" yaml:"sub" bson:"sub,omitempty"`
	Aud                interface{} `json:"aud" yaml:"aud" bson:"aud,omitempty"`
	Scope              string      `json:"scope" yaml:"scope" bson:"scope,omitempty"`
	Exp                int32       `json:"exp" yaml:"exp" bson:"exp,omitempty"`
	ConsumerPlmnId     *PlmnId     `json:"consumerPlmnId,omitempty" yaml:"consumerPlmnId" bson:"consumerPlmnId,omitempty"`
	ProducerPlmnId     *PlmnId     `json:"producerPlmnId,omitempty" yaml:"producerPlmnId" bson:"producerPlmnId,omitempty"`
	ProducerSnssaiList []Snssai    `json:"producerSnssaiList,omitempty" yaml:"producerSnssaiList" bson:"producerSnssaiList,omitempty"`
	ProducerNsiList    []string    `json:"producerNsiList,omitempty" yaml:"producerNsiList" bson:"producerNsiList,omitempty"`
	ProducerNfSetId    string      `json:"producerNfSetId,omitempty" yaml:"producerNfSetId" bson:"producerNfSetId,omitempty"`
	jwt.RegisteredClaims
}
