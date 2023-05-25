package util

var (
	WorkSpace  string
	PluginPath string
	DriverPath string
	// executable file path
	YockPath string
)

// Pathf to format path
//
// @/abc => {WorkSpace}/abc (WorkSpace = UserHome + .yock)
//
// ~/abc => {YockPath}/abc (YockPath = executable file path)
func Pathf(path string) string {
	if len(path) > 0 {
		if path[0] == '@' {
			path = WorkSpace + path[1:]
		} else if path[0] == '~' {
			path = YockPath + path[1:]
		}
	}
	return path
}
