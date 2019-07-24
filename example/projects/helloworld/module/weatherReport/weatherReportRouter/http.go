package weatherReportRouter

import (
	"github.com/gin-gonic/gin"
	"pjx/example/projects/helloworld/module/weatherReport/weatherReportService"
)

func HTTPRouter(r gin.IRouter) {
	r.GET("/weather/", weatherReportService.ShowWeather)
}
