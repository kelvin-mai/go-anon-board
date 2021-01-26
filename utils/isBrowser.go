package utils

import (
	"github.com/avct/uasurfer"
	"github.com/gin-gonic/gin"
)

func IsBrowser(c *gin.Context) bool {
	s := c.GetHeader("User-Agent")
	ua := uasurfer.Parse(s)
	return ua.Browser.Name.String() != "BrowserUnknown"
}
