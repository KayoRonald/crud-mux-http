package response

type CreateResponseTasks struct {
	Name         string `json:"name"`
	Descripition string `json:"descripition"`
	Done         bool 	`json:"done"`
}
// type CreateResponseTasks struct {
// 	Name         string `json:"name" validate:"required"`
// 	Descripition string `json:"descripition" validate:"required"`
// 	Done         bool 	`json:"done" validate:"required"`
// }