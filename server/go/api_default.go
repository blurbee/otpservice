/*
 * One-time Password Service
 *
 * One time password service built in Goglang that uses redis for keystore and can get user information from mongo or postgres.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"github.com/gin-gonic/gin"
)

type OTPServerAPI struct {
}

// Post /otp
// create a new otp session
func (api *OTPServerAPI) CreateOTPSession(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /otp/:id/sendotp
// send OTP to the chosen destination
func (api *OTPServerAPI) SendOTP(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /otp/:id/validate
// given the code, validate the OTP received by the user
func (api *OTPServerAPI) VailidateOTP(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
