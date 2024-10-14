package route

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	config "weather/Config"
	connection "weather/Connection"
	controller "weather/Controller"
	googlecontroller "weather/Controller/GoogleController"
	presentview "weather/Controller/PresentView"
	provider3rdauth "weather/Controller/Provider_3rdAuth"
	middleware "weather/Middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/khaaleoo/gin-rate-limiter/core"
)

// Middlewware close Redis after excute request
// func CloseRedisConnection(rdb *redis.Client) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		//ensuring connection database close after excute request
// 		defer func() {
// 			if err := rdb.Close(); err != nil {
// 				log.Println("error closing connection ", err)
// 			}
// 		}()
// 		ctx.Next()
// 	}
// }
//Implement rate limiting

func Routes() {

	redis_addr := os.Getenv("REDIS_ADDR")
	//Get connect top redis
	ctx := context.Background()
	rdb := connection.RedisCon(redis_addr, ctx)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	config.InitSecureCookie()

	// router.Use(CloseRedisConnection(rdb))
	//Implementation rate limiter
	rateLimiterOption := core.RateLimiterOption{
		Limit: 1,
		Burst: 50,
		Len:   1 * time.Minute,
	}

	// Create an IP rate limiter instance
	rateLimiterMiddleware := core.RequireRateLimiter(core.RateLimiter{
		RateLimiterType: core.IPRateLimiter,
		Key:             "iplimiter_maximum_requests_for_ip_test",
		Option:          rateLimiterOption,
	})
	router.GET("/weather/:Location", rateLimiterMiddleware, func(c *gin.Context) {

		//Lấy tham số từ URL
		location := c.Param("Location")
		//1.Check cache
		//Ensuring consistant of keys in Redis
		cacheKey := strings.ToLower(location)
		cacheData, err := rdb.Get(ctx, cacheKey).Result()
		//2.If not accessing to the other 3 rd service (weather servive)
		if err == redis.Nil {
			//3.Fetching data into cache database
			data, fetchErr := controller.GetWeather(location)
			if fetchErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Failed to fetch weather data ",
				})
				log.Println("Error fetchng wether data :%v", err)
				return
			}
			// Chuyển đổi data thành JSON trước khi lưu vào Redis(Marshall)
			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("Failed to marshall json data :%v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Failed to marshall weather data ",
				})
				return
			}
			//4.Excuting untidy data with TTL 24 hours
			err = rdb.Set(ctx, cacheKey, jsonData, 24*time.Hour).Err()
			if err != nil {
				log.Fatalf("Failed to save into redis ")
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Failed to marshall weather data ",
				})
				return
			}
			c.JSON(http.StatusOK, data)
			// Handle any other Redis errors
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			log.Println("Redis error : %v", err)
		}

		// Return the cached weather data from Redis
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, cacheData)
	})
	//Page link to Google Authen
	router.Static("/Static", "./View/Static")
	router.GET("/google", presentview.ShowGoogleHome)
	//Google Provider
	provider3rdauth.GoogleProvide()
	//Authentication with JWT (Register and Login)

	authen := router.Group("/auth")
	{

		//Create new user
		authen.POST("/user/register", controller.RegisterUser)
		//If user already exist,then generate new token(Login)
		authen.POST("/user/login", controller.GenerateNewToken)
		//Google Oauth 2.0
		authen.GET("/:provider", googlecontroller.SignInGoogle)
		authen.GET("/:provider/callback", googlecontroller.CallBackGoogle)

		secured := authen.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/boom", controller.TestAuth)

		}
	}
	//Result after authentication
	router.GET("/success", presentview.Result)

	//Running on port 8465
	router.Run(":8465")
}
