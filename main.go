package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tarefa struct {
	ID    string `json:id`
	Item  string `json:title`
	Feito bool   `json:completed`
}

var tarefas = []tarefa{
	{ID: "1", Item: "Arrumar o quarto", Feito: false},
	{ID: "2", Item: "Ler um livro", Feito: true},
	{ID: "3", Item: "Estudar para a prova", Feito: false},
}

func getTarefas(contexto *gin.Context) {
	contexto.IndentedJSON(http.StatusOK, tarefas)
}

func main() {
	rota := gin.Default()
	rota.GET("/tarefas", getTarefas)
	rota.Run("localhost:9090")
}
