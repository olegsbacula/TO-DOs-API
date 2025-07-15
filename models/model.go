package models

type TODO struct {
	UserId    string `json:"UserId"`
	TaskID    string `json:"TaskID"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

type ListOfTodos struct{
	TODOs []TODO `json:"todos"`
}
