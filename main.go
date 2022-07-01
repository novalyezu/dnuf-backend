package main

import (
	"dnuf/auth"
	"dnuf/campaign"
	"dnuf/handler"
	"dnuf/helper"
	"dnuf/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_dnuf?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to database")

	authService := auth.NewJwtService()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.Login)
	api.POST("/users/check-email", userHandler.CheckEmail)
	api.POST("/users/avatar", verifyToken(authService, userService), userHandler.UpdateAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:slug", campaignHandler.GetCampaign)
	router.Run()
}

func verifyToken(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if !strings.Contains(authorization, "Bearer") {
			c.JSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, false, "Unauthorized", ""))
			return
		}

		var tokenString string
		tokenSplit := strings.Split(authorization, " ")
		if len(tokenSplit) == 2 {
			tokenString = tokenSplit[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, false, "Unauthorized", ""))
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, false, "Unauthorized", ""))
			return
		}

		rsUser, err := userService.GetUserById(int(claim["user_id"].(float64)))
		if err != nil {
			c.JSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, false, "Unauthorized", ""))
			return
		}

		c.Set("currentUser", rsUser)
	}
}
