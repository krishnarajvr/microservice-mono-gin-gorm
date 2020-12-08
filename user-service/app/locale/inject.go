package locale

import (
	"github.com/gin-gonic/gin"
)

func Inject(locale *Locale) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("lang", locale)
		c.Next()
	}
}
