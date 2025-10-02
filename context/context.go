package context

type Context struct {
	Description      string
	LocalDescription string
	Language         string
	Local            Local
}

type Local struct {
	StartLine int `json:"start-line"`
	EndLine   int `json:"end-line"`
	StartChar int `json:"start-char"`
	EndChar   int `json:"end-char"`

	File      string `json:"file"`
	Package   string `json:"package"`
	Module    string `json:"module"`
	Workspace string `json:"workspace"`
}
