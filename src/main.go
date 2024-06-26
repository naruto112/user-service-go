package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"user-service/src/adapter/dto"
	"user-service/src/adapter/request"
	user_entity "user-service/src/core/domain/entity"
	"user-service/src/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health-check", Healthcheck)
	r.POST("/users/login", loginUser)
	r.Use(request.ValidateTokenMiddleware())
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.GET("/users", getUserAll)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)

	if err != nil {
		intPort = 8000
	}

	r.Run(":" + strconv.Itoa(intPort))
}

func Healthcheck(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Healthcheck ok!</h1>"))
}

func createUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestData request.UserRequest
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.NewUserDTORequest(&requestData)
	userService := services.NewUserServices(userDTO)
	userService.CreateUser(userDTO)

	c.JSON(http.StatusOK, gin.H{"message": "Request criado com sucesso"})
}

func loginUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestData request.UserRequest
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.NewUserDTOLoginRequest(&requestData)
	userService := services.NewUserServices(userDTO)
	token, err := userService.LoginUser(userDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)

}

func getUser(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.NewUserServices(nil)
	user, err := userService.GetUser(uint(id))

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.NewUserDTOResponse(user)

	c.JSON(http.StatusOK, userDTO)
}

func getUserAll(c *gin.Context) {

	userService := services.NewUserServices(nil)
	users, err := userService.GetUserAll()

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.NewUserDTOResponseList(users)

	c.JSON(http.StatusOK, userDTO)
}

func updateUser(c *gin.Context) {

	body, _ := io.ReadAll(c.Request.Body)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestData request.UserRequest
	if err := json.Unmarshal(body, &requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.NewUserServices(nil)
	updatedUser, err := userService.UpdateUser(&user_entity.User{
		ID:       uint(id),
		Name:     requestData.Name,
		Email:    requestData.Email,
		Password: requestData.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updatedUser == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nenhum registro encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Atualizado com sucesso"})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.NewUserServices(nil)
	rowAffect, err := userService.DeleteUser(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if rowAffect == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum registro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deletado com sucesso"})
}
