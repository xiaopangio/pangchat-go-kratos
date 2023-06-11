// Package pkg  @Author xiaobaiio 2023/3/24 16:32:00
package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Ok(c *gin.Context, body any) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "body": body})
}
func OkFail(c *gin.Context, body any, err error) {
	if err != nil {
		Fail(c, err)
	} else {
		Ok(c, body)
	}
}
func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
}
func FailMessage(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": message})
}
func SuccessCodeMessage(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
func SuccessMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": message})
}
func FailCode(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": err.Error()})
}
func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "无权限访问"})
}
func Validator(c *gin.Context, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		errs := gin.H{}
		for _, e := range errors {
			errs[e.StructField()] = e.Translate(Trans)
		}
		c.JSON(http.StatusUnprocessableEntity, errs)
	} else {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
}
func ValidatorError(c *gin.Context, tag, message string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{tag: message})
}
