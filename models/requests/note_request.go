package requests

type CreateNoteRequest struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}

type UpdateNoteRequest struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}
