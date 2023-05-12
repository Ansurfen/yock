package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/yuin/gopher-lua/parse"
)

func TestRestoreScript(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
    local function foo()
        print("Hello, world!")
		return true, 10, false
    end

    foo()
    `))
	chunk, err := parse.Parse(reader, "<string>")
	fmt.Println(parse.Dump(chunk))
	if err != nil {
		panic(err)
	}
	fmt.Println(BuildLuaScript(chunk, nil))
}

func TestRestoreScript2(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
	job("c", function(cenv)
    print("c")
    table.dump(cenv)
    cenv.c = 6
    table.dump(cenv)
    a = 10
    return true
end)
    `))
	chunk, err := parse.Parse(reader, "<string>")
	fmt.Println(parse.Dump(chunk))
	if err != nil {
		panic(err)
	}
	fmt.Println(BuildLuaScript(chunk, nil))
}

func TestRestoreScript3(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
	set_driver("unzip", "yock")

unzip({

}, "./test/abc.zip")

set_driver("unzip", "bandizip")

unzip({
    out = "D:/al/yock/yock/cli/test/out",
	10
}, "./test/test.zip")

    `))
	chunk, err := parse.Parse(reader, "<string>")
	fmt.Println(parse.Dump(chunk))
	if err != nil {
		panic(err)
	}
	fmt.Println(BuildLuaScript(chunk, nil))
}

func TestRestoreScript4(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`
	co({
		task3 = function(this)
			for i = 1, 10 do
				print("I am task 3, executing step " .. i)
				if i == 5 and i == 3 or i == 2 then
					this.wait("y")
				end
				if i == 3 and i == 7 then
					print(8)
				elseif i == 8 then
					print(6)
				elseif i == 8 or (i == 9) then
					print(7)
				else
				end
				coroutine.yield()
			end
			this.notify("x")
		end
	})
	
    `))
	chunk, err := parse.Parse(reader, "<string>")
	fmt.Println(parse.Dump(chunk))
	if err != nil {
		panic(err)
	}
	fmt.Println(BuildLuaScript(chunk, nil))
}

func TestRestoreScript5(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(`print("checkpoint one")

	job("build1", function(cenv)
		print("build1")
		return true
	end)
	
	print("checkpoint two")
	
	job("build2", function(cenv)
		print("build2")
		return true
	end)
	
	job("deploy", function(cenv)
		print("deploy")
		return true
	end)
	
	job("clean", function(cenv)
		print("clean")
		return true
	end)
	
	jobs("all", "build1", "build2", "clean", "deploy")
	jobs("pony", "clean", "deploy")`))
	chunk, err := parse.Parse(reader, "<string>")
	fmt.Println(parse.Dump(chunk))
	if err != nil {
		panic(err)
	}
	fmt.Println(BuildLuaScript(chunk, nil))
	Decomposition(DecompositionOpt{
		Modes: []string{"all", "pony", "build2"},
	}, chunk)
}

func TestBuildBootScript(t *testing.T) {
	buildBootScript("temp.lua", "decomposition.tpl", []string{"host1", "host2"})
}
