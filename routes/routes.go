package routes

import (
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
	ctx.JSON("Didn't find any requested by ID todos.")
}