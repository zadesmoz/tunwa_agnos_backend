package route

import (
	"net/http"
	"tunwa/db"
	"tunwa/serializer"
	"tunwa/utils"

	"github.com/gin-gonic/gin"
)

func PasswordStrengthHandler(c *gin.Context) {
	req := new(serializer.PasswordRequest)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Error{
			Error:   "Bad Request",
			Details: err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, serializer.Error{
			Error:   "Bad Request",
			Details: err.Error()})
		return
	}

	numSteps := utils.CalculatePasswordStrength(req.InitPassword)
	res := serializer.PasswordResponse{NumOfSteps: numSteps}

	// Save Log request and response
	db.LogRequestResponse(c.Request, res)

	c.JSON(http.StatusOK, res)
}
