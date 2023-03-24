package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.ID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, Book := range BookDatas {
		if id == Book.ID {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].ID = id
			break
		}

	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id : not found", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been succsessfully updated", id),
	})
}
func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var BookData Book

	for i, Book := range BookDatas {
		if id == Book.ID {
			condition = true
			BookData = BookDatas[i]
			break

		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Book": BookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if id == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", id),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been succsessfully deleted", id),
	})
}
