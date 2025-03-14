package modules

import (
	"fmt"
	"os"

	"github.com/pius706975/golang-test/package/db_struct"
)

func SaveDataToFile() {
	filePath := "package/db_struct/data.go"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create data.go:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("package db_struct\n\n")
	if err != nil {
		fmt.Println("Failed to write data.go:", err)
		return
	}

	_, err = file.WriteString("var ProgrammingLanguages = []ProgrammingLanguage{\n")
	if err != nil {
		fmt.Println("Failed to write data.go:", err)
		return
	}

	for _, lang := range db_struct.ProgrammingLanguages {
		_, err := file.WriteString(fmt.Sprintf("\t{\n"+
			"\t\tID: %d,\n"+
			"\t\tLanguage: \"%s\",\n"+
			"\t\tAppeared: %d,\n"+
			"\t\tCreated: %#v,\n"+
			"\t\tFunctional: %t,\n"+
			"\t\tObjectOriented: %t,\n"+
			"\t\tRelation: Relation{\n"+
			"\t\t\tInfluencedBy: %#v,\n"+
			"\t\t\tInfluences: %#v,\n"+
			"\t\t},\n"+
			"\t},\n",
			lang.ID,
			lang.Language,
			lang.Appeared,
			lang.Created,
			lang.Functional,
			lang.ObjectOriented,
			lang.Relation.InfluencedBy,
			lang.Relation.Influences,
		))
		if err != nil {
			fmt.Println("Failed to write data.go:", err)
			return
		}
	}

	_, err = file.WriteString("}\n")
	if err != nil {
		fmt.Println("Failed to write data.go:", err)
	}
}
