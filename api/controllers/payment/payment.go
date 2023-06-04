package payment

import (
	"fmt"

	payment "github.com/Ranco-dev/gbms/api/modules/payment"
	"github.com/gin-gonic/gin"
)

// @Summary			創建群組
// @Description	建立一個群組
// @Tags				Group
// @version			1.0
// @produce			application/json
// @param				data	body		Group	true	"body data"
// @Success			201	 "{}"
// @Failure			409		"create duplicates"
// @Failure			500		"internal server error"
// @Router			/api/v1/group [post]

func PymReq(c *gin.Context) {
	fmt.Println("PymReq")
	userId := c.Param("uid")
	amount := c.Param("amount")
	fmt.Println(userId)
	fmt.Println(amount)
	if address, code, err := payment.PaymentReq(userId, amount); err == nil {
		c.JSON(code, gin.H{
			"msg":     "OK",
			"address": address,
		})
	} else {
		c.JSON(code, gin.H{
			"msg": err.Error(),
		})
	}
}

func PymCheck(c *gin.Context) {
	fmt.Println("PymCheck")
	taskId := c.Param("tid")
	fmt.Println(taskId)
	if status, code, err := payment.PaymentCheck(taskId); err == nil {
		c.JSON(code, gin.H{
			"msg":    "OK",
			"status": status,
		})
	} else {
		c.JSON(code, gin.H{
			"msg": err.Error(),
		})
	}

}
