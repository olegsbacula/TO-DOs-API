package routes

import (
	"encoding/json"
	"todosAPI/models"
	"todosAPI/testdata"

	"azugo.io/azugo"
	"github.com/valyala/fasthttp"
)

func GetInfobyID(ctx *azugo.Context) {
	var found bool
	var response models.TODO
	taskID := ctx.Params.String("taskID")
	found = false
	for _, u := range testdata.TODOS.TODOs {
		if u.TaskID == taskID {
			response = models.TODO{
				UserId:    u.UserId,
				TaskID:    u.TaskID,
				Title:     u.Title,
				Completed: u.Completed,
			}
			found = true
		}
	}
	if found {
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

func AddATodo(ctx *azugo.Context) {
	var response models.TODO
	err := json.Unmarshal(ctx.Body.Bytes(), &response)
	if err != nil {
		ctx.StatusCode(fasthttp.StatusBadRequest)
		ctx.ContentType("text/plain")
		ctx.Context().SetBodyString("Error while parsing your JSON")
		return
	}
	for _, u := range testdata.TODOS.TODOs {
		if u.TaskID == response.TaskID {
			ctx.StatusCode(fasthttp.StatusBadRequest)
			ctx.ContentType("text/plain")
			ctx.Context().SetBodyString("Do not enter tasks with the same id's")
			return
		}
	}

	testdata.TODOS.TODOs = append(testdata.TODOS.TODOs, response)
	ctx.StatusCode(fasthttp.StatusOK)
	ctx.JSON(response)
}

func DeleteToDo(ctx *azugo.Context) {
	taskID := ctx.Params.String("taskID")
	var updated []models.TODO
	var found bool

	for _, todo := range testdata.TODOS.TODOs {
		if todo.TaskID != taskID {
			updated = append(updated, todo)
		} else {
			found = true
		}
	}

	if !found {
		ctx.StatusCode(fasthttp.StatusNotFound)
		ctx.ContentType("text/plain")
		ctx.Context().SetBodyString("Task not found")
		return
	}

	testdata.TODOS.TODOs = updated

	ctx.StatusCode(fasthttp.StatusOK)
	ctx.ContentType("text/plain")
	ctx.Context().SetBodyString("Task deleted from memory")
}


func EditAToDo(ctx *azugo.Context) {
	taskID := ctx.Params.String("taskID")
	var response models.TODO
	var updated []models.TODO
	var found bool

	err := json.Unmarshal(ctx.Body.Bytes(), &response)
	if err != nil {
		ctx.StatusCode(fasthttp.StatusBadRequest)
		ctx.ContentType("text/plain")
		ctx.Context().SetBodyString("Error while parsing your JSON")
		return
	}

	for _, todo := range testdata.TODOS.TODOs {
		if todo.TaskID == taskID {
			updated = append(updated, response)
			found = true
		} else {
			updated = append(updated, todo)
		}
	}

	if !found {
		ctx.StatusCode(fasthttp.StatusNotFound)
		ctx.ContentType("text/plain")
		ctx.Context().SetBodyString("Task not found")
		return
	}

	testdata.TODOS.TODOs = updated

	ctx.StatusCode(fasthttp.StatusOK)
	ctx.ContentType("text/plain")
	ctx.Context().SetBodyString("Task updated in memory")
}
