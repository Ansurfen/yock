---@diagnostic disable: undefined-global
http.HandleFunc("/", function(w, req)
    fmt.Fprintf(w, "Hello World!\n")
end)
http.ListenAndServe(":8080", nil)
