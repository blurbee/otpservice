/*
 * One-time Password Service
 *
 * One time password service built in Goglang that uses redis for keystore and can get user information from mongo or postgres.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type ChannelValue struct {

	Id string `json:"id"`

	FullValue string `json:"fullValue"`
}