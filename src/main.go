package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	db "user-service/src/config"
	users "user-service/src/models"
	"user-service/src/repository"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.GET("/users", getUserAll)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
	r.Run(":8000")
}

func createUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestData UserRequest
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userRepository.CreateUser(&users.User{Name: requestData.Name, Email: requestData.Email})

	c.JSON(http.StatusOK, gin.H{"message": "Request recebido com sucesso"})
}

func getUser(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userResponse, err := userRepository.GetUser(uint(id))

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

func getUserAll(c *gin.Context) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userResponse, err := userRepository.GetUserAll()

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

func updateUser(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestData UserRequest
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userRepository.UpdateUser(&users.User{ID: uint(id), Name: requestData.Name, Email: requestData.Email})

	c.JSON(http.StatusOK, gin.H{"message": "Atualizado com sucesso"})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	rowAffect, err := userRepository.DeleteUser(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if rowAffect == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum registro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deletado com sucesso"})
}
