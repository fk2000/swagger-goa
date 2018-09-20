package main

import (
	"time"

	"github.com/fk2000/swagger-goa/app"
	"github.com/goadesign/goa"
)

// TasksController implements the Tasks resource.
type TasksController struct {
	*goa.Controller
}

// NewTasksController creates a Tasks controller.
func NewTasksController(service *goa.Service) *TasksController {
	return &TasksController{Controller: service.NewController("TasksController")}
}

// Create runs the create action.
func (c *TasksController) Create(ctx *app.CreateTasksContext) error {
	// TasksController_Create: start_implement

	ctx.ResponseData.Header().Set("Location", app.TasksHref(999))

	return ctx.Created()

	// TasksController_Create: end_implement
}

// Delete runs the delete action.
func (c *TasksController) Delete(ctx *app.DeleteTasksContext) error {
	// TasksController_Delete: start_implement

	return ctx.NoContent()

	// TasksController_Delete: end_implement
}

// List runs the list action.
func (c *TasksController) List(ctx *app.ListTasksContext) error {
	// TasksController_List: start_implement

	res := app.XLearningGoaCollection{
		{
			ID:        1,
			Title:     "task1",
			Done:      false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "task2",
			Done:      false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Title:     "task3",
			Done:      true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return ctx.OK(res)

	// TasksController_List: end_implement
}

// Show runs the show action.
func (c *TasksController) Show(ctx *app.ShowTasksContext) error {
	// TasksController_Show: start_implement

	if ctx.TaskID == 0 {
		return ctx.NotFound()
	}

	res := &app.XLearningGoa{
		ID:        ctx.TaskID,
		Title:     "example task title",
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return ctx.OK(res)

	// TasksController_Show: end_implement
}

// Update runs the update action.
func (c *TasksController) Update(ctx *app.UpdateTasksContext) error {
	// TasksController_Update: start_implement

	return ctx.NoContent()

	// TasksController_Update: end_implement
}
