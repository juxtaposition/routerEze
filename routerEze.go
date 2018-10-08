package routerEze

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Person struct{
	Name string
	Age int
}

type MyChan struct {
	Channel chan string
} 

func (e *Person) SayHello(greet string) string {
	return greet + e.Name + " with: " + strconv.Itoa(e.Age) + " years old"
}

func (c *MyChan) SendData(data string) bool {
	c.Channel <- data
	return true
}

func JokeHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"developing",
  })
}

func LikeJoke(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message":"developing",
  })
}
