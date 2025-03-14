package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/golang-test/package/db_struct"
)

// CreateLanguage godoc
// @Summary Create new language
// @Description Add a new programming language
// @Tags Languages
// @Accept json
// @Produce json
// @Param request body db_struct.ProgrammingLanguage true "Language Data"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /language [post]
func CreateLanguageController(ctx *gin.Context) {
	var newLanguage db_struct.ProgrammingLanguage

	if err := ctx.ShouldBindJSON(&newLanguage); err != nil {
		ctx.JSON(400, gin.H{"status": 400, "message": "Invalid request payload"})
		return
	}

	response, status := CreateLanguageService(newLanguage)

	ctx.JSON(status, response)
}

// UpdateLanguage godoc
// @Summary Update language by ID
// @Description Update specific fields of a programming language by its ID
// @Tags Languages
// @Accept json
// @Produce json
// @Param id path int true "Language ID"
// @Param request body map[string]interface{} true "Fields to update"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /language/{id} [patch]
func UpdateLanguageController(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedData map[string]interface{}

	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(400, gin.H{"status": 400, "message": "Invalid request payload"})
		return
	}

	response, status := UpdateLanguageService(id, updatedData)

	ctx.JSON(status, response)
}


// DeleteLanguage godoc
// @Summary Delete language by ID
// @Description Delete a programming language by its ID
// @Tags Languages
// @Accept json
// @Produce json
// @Param id path int true "Language ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /language/{id} [delete]
func DeleteLanguageController(ctx *gin.Context) {
	id := ctx.Param("id")

	response, status := DeleteLanguageService(id)

	ctx.JSON(status, response)
}

// GetLanguages godoc
// @Summary Get all languages
// @Description Fetch all languages
// @Tags Languages
// @Accept json
// @Produce json
// @Success 200
// @Failure 500
// @Router /languages [get]
func GetLanguagesController(ctx *gin.Context) {
	languages, status := GetLanguagesService()

	ctx.JSON(status, languages)
}

// GetLanguageByID godoc
// @Summary Get language by ID
// @Description Fetch a programming language by its ID
// @Tags Languages
// @Accept json
// @Produce json
// @Param id path int true "Language ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /language/{id} [get]
func GetLanguageByIDController(ctx *gin.Context) {
	id := ctx.Param("id")

	response, status := GetLanguageByIDService(id)

	ctx.JSON(status, response)
}

// GetPalindrome godoc
// @Summary Get palindrome
// @Description Check if a text is palindrome
// @Tags Palindrome
// @Accept json
// @Produce json
// @Param text query string true "Text to check if palindrome"
// @Success 200
// @Failure 400
// @Router /palindrome [get]
func GetPalindromeController(ctx *gin.Context) {
	text := ctx.Query("text")

	if text == "" {
		ctx.JSON(400, gin.H{"status": 400, "message": "Text parameter is required"})
		return
	}

	palindrome, status := GetPalindromeService(text)

	ctx.JSON(status, palindrome)
}
