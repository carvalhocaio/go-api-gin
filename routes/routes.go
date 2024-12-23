package routes

import (
	"github.com/carvalhocaio/go-api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) }

func HandleRequests() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
