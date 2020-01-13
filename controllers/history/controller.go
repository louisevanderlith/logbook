package history

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/logbook/core"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, nil)
}

func Search(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, nil)
}

// /v1/history/:vehicleKey
func View(c *gin.Context) {
	vehKey := c.Param("key")

	key, err := husk.ParseKey(vehKey)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	rec, err := core.GetHistoryByVehicle(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, rec)
}

// /v1/history
func Create(c *gin.Context) {
	var obj core.History
	err := c.Bind(&obj)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	rec, err := obj.Create()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, rec)
}
