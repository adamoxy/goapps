package main

import (
	"go-rest/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUserHandler(c *gin.Context, queries *db.Queries) {
	var req db.CreateUserParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := queries.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GenerateOTPHandler(c *gin.Context, queries *db.Queries) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := queries.GetUserByPhoneNumber(c, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	otp := GenerateRandomOTP() // Implement this function

	// Create a pgtype.Text variable
	var myText pgtype.Text
	myText.String = otp

	expirationTime := time.Now().Add(time.Minute)
	var pgExpirationTime pgtype.Timestamp
	pgExpirationTime.Time = expirationTime

	err = queries.GenerateOTP(c, db.GenerateOTPParams{
		PhoneNumber:       req.PhoneNumber,
		Otp:               myText,
		OtpExpirationTime: pgExpirationTime,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP generated"})
}

func VerifyOTPHandler(c *gin.Context, queries *db.Queries) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
		OTP         string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userOTP, err := queries.GetUserOTP(c, req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if userOTP.Otp.String != req.OTP || time.Now().After(userOTP.OtpExpirationTime.Time) {
		message := "Invalid OTP"
		if time.Now().After(userOTP.OtpExpirationTime.Time) {
			message = "OTP expired"
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
