print(Windows())
local out, err = cmd("java --version")
yassert(err)
local reg = regexp.MustCompile("java (\\d+)")
print(reg:FindStringSubmatch(out)[2])
