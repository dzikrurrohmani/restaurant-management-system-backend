package api

import (
	"livecode-wmb-2/delivery/api/response"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type BaseApi struct{}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) ParseRequestFormData(c *gin.Context, requestModel interface{}, postFormKey ...string) error {
	mapRes := make(map[string]interface{})
	for _, v := range postFormKey {
		mapRes[v] = c.PostForm(v)
	}
	err := mapstructure.Decode(mapRes, &requestModel)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) Success(c *gin.Context, data interface{}) {
	response.NewSuccessJsonResponse(c, data).Send()
}

func (b *BaseApi) Failed(c *gin.Context, err error) {
	response.NewErrorJsonResponse(c, err).Send()
}
