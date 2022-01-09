package word

import (
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"words/param"

	"github.com/gin-gonic/gin"
	validate "github.com/tonyhb/govalidate"
)

type Words struct {
	Word      string `json:"word"`
	Frequency int    `json:"frequency"`
}

type File struct {
	Text  string `json:"text" validate:"NotEmpty"`
	Limit int    `json:"limit"`
}

func CountWords(c *gin.Context) {
	p := param.Eject(c)
	defer func() {
		if err := recover(); err != nil {
			p.LogrusEntry.Warn("retry again, something went wrong")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "retry again,",
			})
			return
		}
	}()

	var file File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate field text, If empty return the field empty status
	if err := validate.Run(file); err != nil {
		p.LogrusEntry.Warn("empty text field input")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// only alphanumerics are allowed
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	// set the spacing
	data := reg.ReplaceAllString(file.Text, " ")

	words := strings.Fields(data)

	list := make(map[string]int)

	for _, word := range words {
		_, matched := list[word]
		if matched {
			list[word] += 1
		} else {
			list[word] = 1
		}
	}

	ws := []Words{}
	for word, frequency := range list {
		ws = append(ws, Words{word, frequency})
	}

	// sort the words according to count, In decending order
	sort.Slice(ws, func(i, j int) bool {
		return ws[i].Frequency > ws[j].Frequency
	})

	// if input limit is greater then the length of words set the length as limit
	// if input limit is set as zero or less than zero, set the length as limit
	if len(ws) < file.Limit || file.Limit <= 0 {
		file.Limit = len(ws)
	}
	c.JSON(http.StatusOK, gin.H{"data": ws[:file.Limit]})
}
