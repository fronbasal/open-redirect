package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	mgo "gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

type config struct {
	Host   string `json:"host"`
	Key    string `json:"recaptcha_site_key"`
	Secret string `json:"recaptcha_site_secret"`
	Mongo  string `json:"mongo_url"`
}

type captcha struct {
	Sucess bool `json:"success"`
}

type redirect struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func main() {
	r := gin.Default()
	config := readConfig()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		if c.Request.Host != config.Host {
			// Load target from database
		} else {
			c.HTML(200, "index.tpl", gin.H{
				"siteKey": config.Key,
			})
		}
	})
	r.POST("/add", func(c *gin.Context) {
		if c.PostForm("source") != "" && c.PostForm("target") != "" && c.PostForm("g-recaptcha-response") != "" {
			body := strings.NewReader(`secret=` + config.Secret + `&response=` + c.PostForm("g-recaptcha-response"))
			req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", body)
			if err != nil {
				c.JSON(400, gin.H{"message": "Sorry, we could not verify your captcha. Try again!", "error": true})
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				c.JSON(500, gin.H{"message": "Sorry, we could not verify your captcha. Try again!", "error": true})
			}
			defer resp.Body.Close()
			s := captcha{}
			json.NewDecoder(resp.Body).Decode(&s)
			if s.Sucess {
				session, err := mgo.Dial(config.Mongo)
				if err != nil {
					c.JSON(500, gin.H{"message": "Error while establishing database link", "error": true})
				}
				C := session.DB("open-redirect").C("redirections")
				err1 := C.Insert(&redirect{Source: c.PostForm("source"), Target: c.PostForm("target")})
				if err1 != nil {
					c.JSON(500, gin.H{"message": "Error while saving to DB", "error": true})
				}
				c.HTML(200, "success.tpl", gin.H{"domain": c.PostForm("source"), "host": config.Host})
			} else {
				c.JSON(400, gin.H{"message": "Sorry, we could not verify your captcha. Try again!", "error": true})
			}
		} else {
			c.JSON(400, gin.H{"message": "Invalid request!", "error": true})
		}
	})
	r.Run(":5000")
}

func readConfig() config {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	c := config{}
	json.NewDecoder(file).Decode(&c)
	return c
}
