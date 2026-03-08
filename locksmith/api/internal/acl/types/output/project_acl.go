package output

type ProjectAcl struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Modules []Module `json:"modules"`
}

type Module struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Actions []Action `json:"actions"`
}

type Action struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
