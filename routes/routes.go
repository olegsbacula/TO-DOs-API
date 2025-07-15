package routes

import (
	"encoding/json"
	"todosAPI/models"
	"todosAPI/testdata"

	"azugo.io/azugo"
	"github.com/valyala/fasthttp"
)

func GetInfobyID(ctx *azugo.Context){
	var found bool
	var response models.TODO
	taskID := ctx.Params.String("taskID")
	found = false
	for _, u := range testdata.TODOS.TODOs{
		if u.TaskID == taskID{
			response = models.TODO{
				UserId: u.UserId,
				TaskID:u.TaskID,
				Title: u.Title,
				Completed: u.Completed,
			}
			found = true
		}
	}
	if (found){
		ctx.StatusCode(fasthttp.StatusOK)
		ctx.ContentType("application/json")
		ctx.JSON(response)
		return
	}
	ctx.StatusCode(fasthttp.StatusBadRequest)
	ctx.ContentType("text/plain")
	ctx.Context().SetBodyString("Didn't find any requested by ID todos.")
}


func GetAllTodos(ctx *azugo.Context) {
  ctx.StatusCode(fasthttp.StatusOK)
  ctx.ContentType("application/json")
  ctx.JSON(testdata.TODOS.TODOs)
}

func AddATodo(ctx *azugo.Context){
	var response models.TODO
	 err := json.Unmarshal(ctx.Body.Bytes(),&response) 
	 if err != nil{
		ctx.StatusCode(fasthttp.StatusBadRequest)
		ctx.ContentType("text/plain")
		ctx.Context().SetBodyString("Error while parsing your JSON")
		return
	 }
	 for _, u := range testdata.TODOS.TODOs{
		if u.TaskID == response.TaskID{
				ctx.StatusCode(fasthttp.StatusBadRequest)
				ctx.ContentType("text/plain")
				ctx.Context().SetBodyString("Do not enter tasks with the same id's")
				return
			}
		}
		
	testdata.TODOS.TODOs=append(testdata.TODOS.TODOs,response)
	ctx.StatusCode(fasthttp.StatusOK)
	ctx.ContentType("application/json")
	ctx.JSON(response)
}