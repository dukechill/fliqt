package interview

import "github.com/gin-gonic/gin"

const (
	InterviewAPIPath = "/interview"
)

func Route(r *gin.RouterGroup) {
	g := r.Group(InterviewAPIPath)

	g.GET("/interviews", ListInterviews)
	g.GET("/interviews/:id", GetInterviews)
	g.POST("/interviews", CreateInterviews)
	g.PUT("/interviews/:id", UpdateInterviews)
	g.DELETE("/interviews/:id", DeleteInterviews)
}
