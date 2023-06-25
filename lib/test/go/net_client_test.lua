local cli = http.Client()

local repo = "https://github.com"

cli.Timeout = time.second * 10

local res, err = cli:Get(repo)
yassert(err)
local text, err = ioutil.ReadAll(res.Body)
yassert(err)
print(text)
