package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const googleMetadataURL = "http://metadata.google.internal/computeMetadata/v1/instance/zone"

type Model struct {
	SourcePodName  string
	SourceNodeZone string
	DestPodName    string
	DestNodeZone   string
}

func main() {
	log.Println("Starting Response Service")
	router := SetRoute()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Cannot start response service: ", err)
	}
}

func SetRoute() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	router.GET("/", Home)
	router.GET("/ping", HealthCheck)
	router.GET("/data", Render)

	return router
}

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello from Response Service!")
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Render(ctx *gin.Context) {
	nodeInfo := Model{}
	sourceName := ctx.Query("name")
	zoneName := ctx.Query("zone")

	nodeInfo.SourcePodName = sourceName
	nodeInfo.SourceNodeZone = zoneName

	nodeInfo.DestPodName = os.Getenv("PODNAME")
	nodeZone := strings.Split(getNodeZone(), "/")
	nodeInfo.DestNodeZone = nodeZone[len(nodeZone)-1]

	log.Printf("%+v\n", nodeInfo)
	ctx.JSON(http.StatusOK, nodeInfo)
}

func getNodeZone() string {
	req, err := http.NewRequest("GET", googleMetadataURL, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	req.Header.Add("Metadata-Flavour", "Google")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	zone_temp := strings.Split(string(resBody), "/")
	zone := zone_temp[len(zone_temp)-1]
	return zone
}
