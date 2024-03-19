package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

// Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return fmt.Errorf("parse request query failed: %v", err.Error())

	}
	return nil
}

// Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return fmt.Errorf("parse request form failed: %v", err.Error())
	}
	return nil
}

// Parse body form data to struct
func ParseFormPost(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.FormPost); err != nil {
		return fmt.Errorf("parse request post failed: %v", err.Error())
	}
	return nil
}

// Parse body form data to struct
func ParseFormMultipart(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.FormMultipart); err != nil {
		return fmt.Errorf("parse request form failed: %v", err.Error())
	}
	return nil
}

// Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return fmt.Errorf("parse request json failed: %v", err.Error())
	}
	return nil
}

func ParseQueryID(c *gin.Context, key string) uint64 {
	val := c.Query(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func ParseParamID(c *gin.Context, key string) uint64 {
	val := c.Param(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// Param returns the value of the URL param
func ParseParamIDs(c *gin.Context, key string) []string {
	return c.QueryArray(key)
}
