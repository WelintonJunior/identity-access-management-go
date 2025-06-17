package main

import "github.com/WelintonJunior/identity-access-management-go/cmd"

// @title WattsUp API
// @version 1.0
// @description API de autenticação e gestão de usuários
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
