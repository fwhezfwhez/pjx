package weatherReportService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func ShowWeather(c *gin.Context) {
	c.String(200, fmt.Sprintf("%s %s", time.Now().Format("2006-01-02"), "rain"))
}
