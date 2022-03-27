package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"manager"
)

type SampleRequest struct {
	Ids[] int `json:"ids"`
}


func HandleRequest(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var t SampleRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	manager.Process()
}