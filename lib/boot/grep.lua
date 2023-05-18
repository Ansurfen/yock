local riggrep = import("riggrep@v1")
riggrep.install()

grep = function(pattern, file)
    riggrep.run()
end

grep()