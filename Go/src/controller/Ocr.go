package controller

import (
	"bytes"
	"github.com/ledongthuc/pdf"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)


func Ocr(c *gin.Context) {

	file, _ := c.FormFile("file")
	buffer, _ := file.Open()
	defer buffer.Close()
	content, _ := readPdf(buffer));

	c.JSON(http.StatusOK, gin.H{
		"data": "file received",
	})
}


func readPdf(file *os.File ) (string, error) {
	pdf.NewReader(file.Read())
	totalPage := file.NumPage()

	var textBuilder bytes.Buffer
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		textBuilder.WriteString(p.GetPlainText("\n"))
	}
	return textBuilder.String(), nil
}