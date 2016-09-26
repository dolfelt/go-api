package data

import (
	"fmt"

	"github.com/gin-gonic/gin"

	validator "gopkg.in/go-playground/validator.v9"
)

// PrintErrors exports errors and translates them
func PrintErrors(err error, context *gin.Context) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
		return
	}

	for _, err := range err.(validator.ValidationErrors) {
		context.JSON(400, &gin.H{
			"error": fmt.Sprintf("Field `%s` is invalid because it must be %s", err.Field(), err.Tag()),
		})
		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
		fmt.Println(err.StructField())     // by passing alt name to ReportError like below
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Printf("Value: %s\n", err.Value())
		fmt.Printf("Param: %s\n", err.Param())
		fmt.Println(err)
	}
}
