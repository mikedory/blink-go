package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "github.com/gin-gonic/contrib/static"
)

func main() {

    // fire this up
    router := gin.Default()
    router.Use(gin.Logger())

    // load the dang templates
    router.LoadHTMLGlob("templates/*.html")
    // router.Use(static.Serve("/static", static.LocalFile("html", false)))
    router.Static("/static", "static")

    // handle root
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "wat")
    })
    
    // check the incoming phrase
    router.GET("/blink/:phrase", func(c *gin.Context) {
        phrase := c.Param("phrase")  
        c.HTML(http.StatusOK, "main.html", gin.H{
            "title": "Main website",
            "phrase": phrase,
            "twitter_handle": "mike_dory",
            "google_analytics_id": "XXXXX-XX",
        })
    })

    // run!
    router.Run(":8081")

}

