-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

co({
    task1 = function(this)
        for i = 1, 5 do
            this.wait("x")
            print("I am task 1, executing step " .. i)
            coroutine.yield()
        end
    end,
    task2 = function(this)
        for i = 1, 10 do
            this.wait("x")
            print("I am task 2, executing step " .. i)
            coroutine.yield()
        end
    end,
    task3 = function(this)
        for i = 1, 10 do
            print("I am task 3, executing step " .. i)
            if i == 5 then
                this.wait("y")
            end
            coroutine.yield()
        end
        this.notify("x")
    end,
    task4 = function(this)
        for i = 1, 10 do
            print("I am task 4, executing step " .. i)
            if i == 9 then
                this.notify("y")
            end
            coroutine.yield()
        end
    end
})
