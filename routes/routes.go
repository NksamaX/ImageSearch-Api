package routes

import (
	"fmt"

	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Welcome to Images Search API",
		"routes":  []string{"/images/:name"},
	})
}

func Images(c *gin.Context) {
	url := c.Param("name")
	fmt.Println(url)
	images, _ := images(url)
	c.JSON(200, gin.H{
		"images": images,
	})
}

func Img(c *gin.Context) {
	c.JSON(200, gin.H{
		"how to use": "/images/name",
	})
}

func images(name string) ([]string, error) {
	res, err := soup.Get(fmt.Sprintf("https://ca.images.search.yahoo.com/search/images;_ylt=AwrJ7KMan9dhurIAGAy7HAx.;_ylu=Y29sbwNiZjEEcG9zAzEEdnRpZAMEc2VjA3BpdnM-?p=%v&fr2=piv-web&fr=sfp", name))
	if err != nil {
		return nil, err
	}

	var images []string

	doc := soup.HTMLParse(res)
	x := doc.Find("div", "id", "res-cont").FindAll("a", "class", "img")
	for _, y := range x {
		k := fmt.Sprintf("https://images.search.yahoo.com%v", y.Attrs()["href"])
		images = append(images, k)
	}

	return images, nil

}
