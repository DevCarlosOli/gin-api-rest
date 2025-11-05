package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevCarlosOli/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/carlos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", response.Code, http.StatusOK)
	}
}
