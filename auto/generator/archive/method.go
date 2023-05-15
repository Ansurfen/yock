package archive

type Method struct {
	Comment []string
	Params  []MethodArgument
	Results []MethodArgument
	Package string
	Name    string
}

type MethodArgument struct {
	Name      string
	TypeIdent string
}
