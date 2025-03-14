package db_struct

var ProgrammingLanguages = []ProgrammingLanguage{
	{
		ID: 1,
		Language: "C Language",
		Appeared: 1972,
		Created: []string{"Dennis Ritchie"},
		Functional: false,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{"B", "ALGOL 68"},
			Influences: []string{"C++", "Java"},
		},
	},
}
