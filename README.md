
# Weather_API v1.0
This is my project following[ https://roadmap.sh/projects/weather-api-wrapper-service ](https://roadmap.sh/projects/weather-api-wrapper-service)
- A picture below will visualize  how the project was made.
  <p align="center" width="100%" >
    <img width="100%" src="https://assets.roadmap.sh/guest/weather-api-f8i1q.png" > 
</p>

# Update
- 10/08:Implement rate limiting
- 11/08: Add new JWT token function

# Presequites: 
- Install Go V1.16 or superior
- Install Redis
- Install Postman or RESTful API 

# How to use it ?
1  Clone my project on your machine 
```http
  git clone https://github.com/ducgiang2003/weather-API-v1.0.git
```
2 Then change your current working directory  
```http
 cd WeatherAPI
```
3  After that, run the Docker file to pull Redis .If you already have Redis on your local machine , you don't need do this step   
- Build and compose up Docker
```http
docker compose -f "docker-compose.yml" up -d --build
```
- Make sure your Radis is running; you can configure it in redis.go file in Connection folder
4 Install Go dependencies:
```http
go mod download
```
5 Run the application:
```http
go run main.go
```
# Usage
Send a GET request to the following endpoint:
```http
http://localhost:8465/weather/your_location 
Example :
http://localhost:8465/weather/HoChiMinh
```
# License
This project is licensed under the MIT License.

