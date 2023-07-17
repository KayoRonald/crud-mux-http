package response

type CreateTasks struct {
	Name         string `json:"name"`
	Descripition string `json:"descripition"`
	Done         bool   `json:"done"`
}