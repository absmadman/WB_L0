package entities

type Page struct {
	Content     string
	CurrPageNum int
	TotalPages  int
	NextPageNum int
	PrevPageNum int
	LastPageNum int
}

func NewPage(content string, curr, total, next, prev, last int) *Page {
	return &Page{
		content,
		curr,
		total,
		next,
		prev,
		last,
	}
}
