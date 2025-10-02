package context

type Context struct {
	Description      string
	LocalDescription string
	Language         string
	Local            Local
}

type Local struct {
	StartLine int
	EndLine   int
	StartChar int
	EndChar   int

	Package   string
	Module    string
	Workspace string
}
