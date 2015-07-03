package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {

    // fire this up
    router := gin.Default()

    // load the dang templates
    router.LoadHTMLGlob("templates/*.html")
    router.Static("/static", "static")
    //routerouter.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

    // handle root
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "wat")
    })

    
    // check the incoming phrase
    router.GET("/:phrase", func(c *gin.Context) {
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

