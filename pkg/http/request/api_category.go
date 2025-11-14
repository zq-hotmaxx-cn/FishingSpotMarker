package request

type Post_Api_Category_Request struct {
	Name        string `json:"name" validate:"required, min=1, max=128"`
	Description string `json:"description" validate:"min=0, max=1024"`
}
