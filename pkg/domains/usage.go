package domains

type UsageResponse struct {
	Page            string  `json:"page,omitempty"`
	PageSize        string  `json:"page_size,omitempty"`
	NumPages        string  `json:"num_pages,omitempty"`
	Start           string  `json:"start,omitempty"`
	Total           string  `json:"total,omitempty"`
	End             string  `json:"end,omitempty"`
	Uri             string  `json:"uri,omitempty"`
	FirstPageUri    string  `json:"first_page_uri,omitempty"`
	LastPageUri     string  `json:"last_page_uri,omitempty"`
	NextPageUri     string  `json:"next_page_uri,omitempty"`
	PreviousPageUri string  `json:"previous_page_uri,omitempty"`
	Usages          []Usage `json:"usages,omitempty"`
}

type Usage struct {
	Sid         string `json:"sid,omitempty"`
	Product     string `json:"product,omitempty"`
	ProdutctId  string `json:"product_id,omitempty"`
	Month       string `json:"month,omitempty"`
	Year        string `json:"year,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
	AverageCost string `json:"average_cost,omitempty"`
	TotalCost   string `json:"total_cost,omitempty"`
	Uri         string `json:"uri,omitempty"`
}
