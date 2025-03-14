package modules

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pius706975/golang-test/package/db_struct"
)

func CreateLanguageService(newLanguage db_struct.ProgrammingLanguage) (gin.H, int) {
	if strings.TrimSpace(newLanguage.Language) == "" || newLanguage.Appeared == 0 {
		return gin.H{"status": 400, "message": "Language name and appeared year are required"}, 400
	}

	lastID := 0
	if len(db_struct.ProgrammingLanguages) > 0 {
		lastID = db_struct.ProgrammingLanguages[len(db_struct.ProgrammingLanguages)-1].ID
	}

	newLanguage.ID = lastID + 1

	CreateLanguageRepo(newLanguage)

	return gin.H{"status": 201, "message": "Language created successfully", "data": newLanguage}, 201
}

func UpdateLanguageService(idStr string, updatedData map[string]interface{}) (gin.H, int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return gin.H{"status": 400, "message": "Invalid ID format"}, 400
	}

	updatedLanguage, updated := UpdateLanguageRepo(id, updatedData)
	if !updated {
		return gin.H{"status": 404, "message": "Language not found"}, 404
	}

	return gin.H{"status": 200, "message": "Language updated successfully", "data": updatedLanguage}, 200
}

func DeleteLanguageService(idStr string) (gin.H, int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return gin.H{"status": 400, "message": "Invalid ID format"}, 400
	}

	deleted := DeleteLanguageRepo(id)
	if !deleted {
		return gin.H{"status": 404, "message": "Language not found"}, 404
	}

	return gin.H{"status": 200, "message": "Language deleted successfully"}, 200
}

func GetLanguagesService() (gin.H, int) {
	languages := GetLanguagesRepo()
	return gin.H{"status": 200, "data": languages}, 200
}

func GetLanguageByIDService(idStr string) (gin.H, int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return gin.H{"status": 400, "message": "Invalid ID format"}, 400
	}

	language, found := GetLanguageByIDRepo(id)
	if !found {
		return gin.H{"status": 404, "message": "Language not found"}, 404
	}

	return gin.H{"status": 200, "data": language}, 200
}

func GetPalindromeService(text string) (gin.H, int) {
	cleanedText := strings.ToLower(strings.ReplaceAll(text, " ", ""))

	isPalindrome := true
	length := len(cleanedText)

	for i := 0; i < length/2; i++ {
		if cleanedText[i] != cleanedText[length-1-i] {
			isPalindrome = false
			break
		}
	}

	if !isPalindrome {
		return gin.H{"status": 400, "message": text + " is not a palindrome"}, 400
	}

	return gin.H{"status": 200, "message": text + " is a palindrome"}, 200
}
