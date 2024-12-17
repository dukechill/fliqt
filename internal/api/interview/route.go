package interview

import "github.com/gin-gonic/gin"

const (
	InterviewAPIPath = "/interview"
)

func Route(r *gin.RouterGroup) {

	g := r.Group(InterviewAPIPath)

	//// CORS settings and add OPTIONS route to prevent 404 error
	//g.Use(middleware.CorsFactory(middleware.Admin))
	//g.Use(middleware.LanguageAccept())
	//middleware.CorsRoute(g)
	//
	//auth.Route(g)
	//
	//// All APIs need auth except auth
	//g.Use(middleware.AdminAuthRequired())

	interviewMiddleware := NewInterviewMiddleware(interviewRepo, logger)
	r.GET("/interviews", interviewMiddleware.ListInterviews)
	r.GET("/interviews/:id", interviewMiddleware.GetInterviews)
	r.POST("/interviews", interviewMiddleware.CreateInterviews)
	r.PUT("/interviews/:id", interviewMiddleware.UpdateInterviews)
	r.DELETE("/interviews/:id", interviewMiddleware.DeleteInterviews)

}
