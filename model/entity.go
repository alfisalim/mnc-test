package model

type Biodata struct {
	Language 		string  	`json:"language"`
	Appeared 		int			`json:"appeared"`
	Created			[]string	`json:"created"`
	Functional 		bool		`json:"functional"`
	ObjectOriented 	bool		`json:"object-oriented"`
	Relation		Relation	`json:"relation"`
}

type Relation struct {
	InfluencedBy 	[]string	`json:"influenced-by"`
	Influences		[]string	`json:"influences"`
}

type CustomResponseHTTP struct {
	StatusCode  	int			`json:"status-code"`
	Message 		string		`json:"message"`
	Data 			interface{} `json:"data"`
}

func FillTheStruct() *Biodata {
	return &Biodata{
		Language: "C",
		Appeared: 1927,
		Created: []string{"Dennis Ritchie"},
		Functional: true,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{
				"B",
				"ALGOL 68",
				"Assembly",
				"FORTRAN",
			},
			Influences: []string{
				"C++",
				"Objective-C",
				"C#",
				"Java",
				"JavaScript",
				"PHP",
				"Go",
			},
		},
	}
}
