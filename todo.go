package todo

type TodoList struct {
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id int
	UserIt int
	ListId int
}

type TodoItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

type ListItem struct {
	Id int
	ListId int
	ItemId int
}