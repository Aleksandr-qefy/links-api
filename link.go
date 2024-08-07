package links_api

type Link struct {
	Id          UUID   `json:"id"`
	UserId      UUID   `json:"userId"`
	Ref         string `json:"ref"`
	Description string `json:"description,omitempty"`
}

type Category struct {
	Id   UUID   `json:"id"`
	Name string `json:"name"`
}

type LinkCategory struct {
	LinkId     UUID `json:"linkId"`
	CategoryId UUID `json:"categoryId"`
}
