package appmodule

type DataResponse struct {
	Success      bool   `json:"success,omitempty"`
	Data         any    `json:"data,omitempty"`
	MetaData     any    `json:"metaData,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type TreeResponse struct {
	Success      bool   `json:"success,omitempty"`
	Children     any    `json:"children,omitempty"`
	MetaData     any    `json:"metaData,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
