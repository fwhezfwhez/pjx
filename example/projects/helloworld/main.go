package main

import (
	"github.com/gin-gonic/gin"
	"pjx/example/projects/helloworld/module/weatherReport/weatherReportRouter"
)

func main() {
    r := gin.Default()
    // weather report
    weatherReportRouter.HTTPRouter(r)

    r.Run(":8080")
}
