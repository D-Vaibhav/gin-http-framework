package entity

// gte - greater than or equals to,	 lte - less than equal to
type Person struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

// binding is an attribute used to give data infomation about the property
type Video struct {
	Title       string `json:"title" binding:"min=2,max=100" validate:"is-cool"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
