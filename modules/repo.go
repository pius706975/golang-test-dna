package modules

import "slices"

import "github.com/pius706975/golang-test/package/db_struct"

func CreateLanguageRepo(newLanguage db_struct.ProgrammingLanguage) {
	db_struct.ProgrammingLanguages = append(db_struct.ProgrammingLanguages, newLanguage)

	SaveDataToFile()
}

func UpdateLanguageRepo(id int, updatedData map[string]interface{}) (db_struct.ProgrammingLanguage, bool) {
	for i, lang := range db_struct.ProgrammingLanguages {
		if lang.ID == id {
			if language, ok := updatedData["language"].(string); ok {
				db_struct.ProgrammingLanguages[i].Language = language
			}
			if appeared, ok := updatedData["appeared"].(int); ok {
				db_struct.ProgrammingLanguages[i].Appeared = appeared
			}
			if created, ok := updatedData["created"].([]interface{}); ok {
				var createdStr []string
				for _, v := range created {
					if str, ok := v.(string); ok {
						createdStr = append(createdStr, str)
					}
				}
				db_struct.ProgrammingLanguages[i].Created = createdStr
			}
			if functional, ok := updatedData["functional"].(bool); ok {
				db_struct.ProgrammingLanguages[i].Functional = functional
			}
			if objectOriented, ok := updatedData["object-oriented"].(bool); ok {
				db_struct.ProgrammingLanguages[i].ObjectOriented = objectOriented
			}
			if relation, ok := updatedData["relation"].(map[string]interface{}); ok {
				if influencedBy, exists := relation["influenced-by"].([]interface{}); exists {
					var influencedByStr []string
					for _, v := range influencedBy {
						if str, ok := v.(string); ok {
							influencedByStr = append(influencedByStr, str)
						}
					}
					db_struct.ProgrammingLanguages[i].Relation.InfluencedBy = influencedByStr
				}
				if influences, exists := relation["influences"].([]interface{}); exists {
					var influencesStr []string
					for _, v := range influences {
						if str, ok := v.(string); ok {
							influencesStr = append(influencesStr, str)
						}
					}
					db_struct.ProgrammingLanguages[i].Relation.Influences = influencesStr
				}
			}

			SaveDataToFile()

			return db_struct.ProgrammingLanguages[i], true
		}
	}
	return db_struct.ProgrammingLanguage{}, false
}

func DeleteLanguageRepo(id int) bool {
	for i, language := range db_struct.ProgrammingLanguages {
		if language.ID == id {
			db_struct.ProgrammingLanguages = slices.Delete(db_struct.ProgrammingLanguages, i, i+1)

			SaveDataToFile()
			return true
		}
	}
	return false
}

func GetLanguagesRepo() []db_struct.ProgrammingLanguage {
	return db_struct.ProgrammingLanguages
}

func GetLanguageByIDRepo(id int) (db_struct.ProgrammingLanguage, bool) {
	for _, language := range db_struct.ProgrammingLanguages {
		if language.ID == id {
			return language, true
		}
	}
	return db_struct.ProgrammingLanguage{}, false
}
