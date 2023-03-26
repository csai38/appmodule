package appmodule

type DirectResponse struct {
	ReqType string         `json:"type"`
	TId     int            `json:"tid"`
	Action  string         `json:"action"`
	Method  string         `json:"method"`
	Result  map[string]any `json:"result"`
}

type DataResponse struct {
	Success      bool   `json:"success"`
	Data         any    `json:"data"`
	MetaData     any    `json:"metaData"`
	ErrorMessage string `json:"errorMessage"`
}

type TreeResponse struct {
	Success      bool   `json:"success"`
	Children     any    `json:"children"`
	MetaData     any    `json:"metaData"`
	ErrorMessage string `json:"errorMessage"`
}
