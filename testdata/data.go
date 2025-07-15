package testdata
import(
	"todosAPI/models"
)
var ( 
TODOS= models.ListOfTodos{
	TODOs: []models.TODO{
		{
		UserId: "1",
		TaskID:"1",
		Title: "Clean a room",
		Completed: false,
},

{
		UserId: "2",
		TaskID:"2",
		Title: "Clean a house",
		Completed: true,
},

{
		UserId: "3",
		TaskID:"3",
		Title: "Earn some money",
		Completed: false,
},

{
		UserId: "4",
		TaskID:"4",
		Title: "Buy a car",
		Completed: true,
},

{
		UserId: "5",
		TaskID:"5",
		Title: "Eat a breakfast",
		Completed: true,
},
	},
}
	
)