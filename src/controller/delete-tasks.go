package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	//"net/http"
)

func DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	fmt.Println(taskId)
	if taskId == "1" {
		errRest := rest_err.NewBadRequest("Invalid userId")
		c.JSON(errRest.Code, errRest)
		return
	}
	/*
		err := uc.service.DeleteUser(userId)
		if err != nil {
			logger.Error(
				"Error trying to call deleteUser service",
				err,
				zap.String("journey", "deleteUser"))
			c.JSON(err.Code, err)
			return
		}

		c.Status(http.StatusOK)
	*/
}
