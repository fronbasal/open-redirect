package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

type config struct {
	Host    string `json:"host"`
	Key     string `json:"recaptcha_site_key"`
	Secret  string `json:"recaptcha_site_secret"`
	Mongo   string `json:"mongo_url"`
	Contact string `json:"contact"`
}

type captcha struct {
	Sucess bool `json:"success"`
}

type redirect struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Pin    string `json:"pin"`
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
				"contact": config.Contact,
			})
		}
	})
	r.POST("/add", func(c *gin.Context) {
		if c.PostForm("source") != "" && c.PostForm("target") != "" && c.PostForm("g-recaptcha-response") != "" && c.PostForm("pin") != "" {
			body := strings.NewReader(`secret=` + config.Secret + `&response=` + c.PostForm("g-recaptcha-response"))
			req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", body)
			if err != nil {
				c.JSON(400, gin.H{"message": "Sorry, we could not verify your captcha. Try again!", "error": true})
				return
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				c.JSON(500, gin.H{"message": "Sorry, we could not verify your captcha. Try again!", "error": true})
				return
			}
			defer resp.Body.Close()
			s := captcha{}
			json.NewDecoder(resp.Body).Decode(&s)
			if s.Sucess {
				session, err := mgo.Dial(config.Mongo)
				if err != nil {
					c.JSON(500, gin.H{"message": "Error while establishing database link", "error": true})
					return
				}
				C := session.DB("open-redirect").C("redirections")
				result := redirect{}
				err2 := C.Find(bson.M{"source": c.PostForm("source")}).One(&result)
				if err2 == nil {
					c.JSON(400, gin.H{"message": "This domain is already in the database! (Database error)", "error": true})
					return
				}
				err1 := C.Insert(&redirect{Source: c.PostForm("source"), Target: c.PostForm("target"), Pin: c.PostForm("pin")})
				if err1 != nil {
					c.JSON(500, gin.H{"message": "Error while saving to DB", "error": true})
					return
				}
				c.HTML(200, "success.tpl", gin.H{"domain": c.PostForm("source"), "host": config.Host})
			} else {
				c.JSON(400, gin.H{"message": "Sorry, we could not verify your captcha. Try again!", "error": true})
				return
			}
		} else {
			c.JSON(400, gin.H{"message": "Invalid request!", "error": true})
			return
		}
	})

	r.POST("/delete", func(c *gin.Context) {
		if c.PostForm("source") != "" && c.PostForm("pin") != "" {
			session, err := mgo.Dial(config.Mongo)
			if err != nil {
				c.JSON(500, gin.H{"message": "Error while establishing database link", "error": true})
				return
			}
			C := session.DB("open-redirect").C("redirections")
			result := redirect{}
			err2 := C.Find(bson.M{"source": c.PostForm("source")}).One(&result)
			if err2 != nil {
				c.JSON(400, gin.H{"message": "The domain your entered does not exist! Try creating it first!", "error": true})
				return
			}

		} else {
			c.JSON(400, gin.H{"message": "Invalid request!", "error": true})
			return
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
