package responces

type AutocompleteResponce struct {
	Total int64 `json:"total"`
	Docs  []Doc `json:"docs"`
}

type Doc struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  int64   `json:"_score"`
	Source Source  `json:"_source"`
	Sort   []int64 `json:"sort"`
}

type Source struct {
	FNS FNS `json:"ФНС"`
}

type FNS struct {
	Head     []Head `json:"Руководители"`
	INN      string `json:"ИНН"`
	Activity int64  `json:"Активность"`
	URSHORT  string `json:"НаимЮЛСокр"`
	OGRN     string `json:"ОГРН"`
	URLONG   string `json:"НаимЮЛПолн"`
}

type Head struct {
	INNFL string `json:"ИННФЛ"`
	NED   int64  `json:"Нед"`
	FIO   string `json:"ФИО"`
}
