package invade

import (
	"fmt"
	"strconv"

	"github.com/Ranco-dev/gbms/api/modules/invade"
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

func isNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func GetInvadeLogLast(c *gin.Context) {
	fmt.Println("GetInvadeLogLast")
	secondDigit := c.Param("secondDigit")

	if result, code, err := invade.GetInvadeLogLast(secondDigit); err == nil {
		c.JSON(code, gin.H{
			"msg":    "OK",
			"result": result,
		})
	} else {
		c.JSON(code, gin.H{
			"msg": err.Error(),
		})
	}
}

func InvadeLog(c *gin.Context) {
	fmt.Println("InvadeLog")
	value := c.Param("value")
	fmt.Println(value)

	if result, code, err := invade.InvadeLog(value); err == nil {
		c.JSON(code, gin.H{
			"msg": result,
		})
	} else {
		c.JSON(code, gin.H{
			"msg": err.Error(),
		})
	}
}
