local m = tea.NewModel()
local quit = false
local choices = { "apple", "banana" }
local cursor = 1
local res = {}
KeyEvent = {
    Enter = "enter",
    Up = "up",
    Down = "down",
    Space = " "
}

m.InitCallback = function()
    return nil
end

local KeyMsg = reflect.TypeOf(tea.KeyMsg):String()

m.UpdateCallback = function(msg)
    local rt = reflect.TypeOf(msg):String()
    if rt == KeyMsg then
        local methodString = reflect.ValueOf(msg):MethodByName("String")
        local key = methodString:Call({})[1]:String()
        if key == KeyEvent.Enter then
            quit = true
            return m, tea.Quit
        elseif key == KeyEvent.Up then
            if cursor > 1 then
                cursor = cursor - 1
            end
        elseif key == KeyEvent.Down then
            if cursor < #choices then
                cursor = cursor + 1
            end
        elseif key == KeyEvent.Space then
            local selected = choices[cursor]
            if res[selected] == nil then
                res[selected] = true
            else
                res[selected] = nil
            end
        end
    end
    return m, nil
end
local theme_danger = tea.NewStyle():Foreground("#d251a6")
local theme_info = tea.NewStyle():Foreground("#8866e9")
m.ViewCallback = function()
    if quit then
        return "exit"
    end
    local s = "select\n\n"
    for i, choice in ipairs(choices) do
        local c = " "
        if cursor == i then
            c = ">"
        end
        local checked = " "
        if res[choice] ~= nil then
            checked = "x"
        end
        s = s .. string.format("%s [%s] %s\n", c, theme_danger:Render(checked), choice)
    end
    return s
end

local p = tea.NewProgram(m)
local _, err = p:Run()
yassert(err)
