package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"quiz-app/controllers"
	"quiz-app/services"
	"strconv"
)

func Route(r *gin.Engine) {

	r.Use(cors.Default())

	api := r.Group("/api")
	{
		courseQuiz := api.Group("/course-quiz")
		{
			service := services.NewCourseQuizService()
			controller := controllers.NewCourseQuizController(service)

			courseQuiz.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, controller.GetAllCourse())
			})

			courseQuiz.GET("/:id", func(c *gin.Context) {
				id, _ := strconv.Atoi(c.Param("id"))
				course, err := controller.GetById(id)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, course)
			})

			courseQuiz.PUT("/:id", func(c *gin.Context) {
				id, _ := strconv.Atoi(c.Param("id"))
				err := controller.UpdateCourse(id, c)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"message": "Update successfully",
				})
			})

			courseQuiz.POST("/", func(c *gin.Context) {
				course, err := controller.AddCourse(c)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusCreated, course)
			})

			courseQuiz.DELETE("/:id", func(c *gin.Context) {
				id, _ := strconv.Atoi(c.Param("id"))
				err := controller.DeleteCourse(id)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"message": "Delete successfully",
				})
			})
		}

		dataQuiz := api.Group("/data-quiz")
		{
			service := services.NewDataQuizService()
			controller := controllers.NewDataQuizController(service)

			dataQuiz.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, controller.GetAllData())
			})

			dataQuiz.GET("/:id", func(c *gin.Context) {
				id, _ := strconv.Atoi(c.Param("id"))
				data, err := controller.GetById(id)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, data)
			})

			dataQuiz.PUT("/:id", func(c *gin.Context) {
				id, _ := strconv.Atoi(c.Param("id"))
				err := controller.UpdateData(id, c)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"message": "Update successfully",
				})
			})

			dataQuiz.POST("", func(c *gin.Context) {
				data, err := controller.AddData(c)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusCreated, data)
			})

			dataQuiz.DELETE("/:id", func(c *gin.Context) {
				id, _ := strconv.Atoi(c.Param("id"))
				err := controller.DeleteData(id)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": err.Error(),
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"message": "Delete successfully",
				})
			})
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "HelloWorld",
		})
	})

}
