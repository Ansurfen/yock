job("multi", function(cenv)
    go(function()
        local idx = 0
        while idx ~= 5 do
            print("task 1")
            time.sleep(1 * time.second)
            idx = idx + 1
        end
        notify("x")
        print("task1 fine")
    end)
    go(function()
        print("task 2")
        wait("x")
        print("task2 fine")
    end)

    go(function()
        print("task 3")
        wait("x")
        print("task3 fine")
    end)

    go(function()
        time.sleep(20 * time.second)
        notify("a")
        print("push a")
    end)

    waits("x", "a", "b")
    return true
end)
