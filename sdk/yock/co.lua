---@diagnostic disable-next-line: lowercase-global
function co(todo)
    local sigs = {}
    local wait = function(sig)
        if sigs[sig] == nil then
            sigs[sig] = false
        end
        while not sigs[sig] do
            coroutine.yield(sig)
        end
    end
    local notify = function(sig)
        if sigs[sig] ~= nil then
            sigs[sig] = true
        end
    end
    local this = {
        wait = wait,
        notify = notify
    }
    local tasks = {}
    for key, value in pairs(todo) do
        table.insert(tasks, {
            instance = coroutine.create(function()
                value(this)
            end),
            name = key
        })
    end
    local task_cnt = 0
    local wait_cnt = 0
    while #tasks ~= 0 do
        local current_task = tasks[1]
        table.remove(tasks, 1)
        local _, sig = coroutine.resume(current_task.instance)
        task_cnt = task_cnt + 1
        if sig ~= nil then
            wait_cnt = wait_cnt + 1
        else
            wait_cnt = wait_cnt - 1
        end
        if coroutine.status(current_task.instance) == "suspended" then
            table.insert(tasks, current_task)
        end
        if task_cnt >= #tasks then
            if wait_cnt == task_cnt then
                print("deadlock!")
                return
            end
            task_cnt = 0
            wait_cnt = 0
        end
    end
end
