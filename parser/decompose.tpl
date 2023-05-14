job_option({
    flags = {
        {{- $item_len := (len .) -}}
        {{- range $i, $item := . }}
        {{.Name}} = {
            ip = "localhost"
        }{{if ne (add $i 1) $item_len}},{{end}}

        {{- end }}
    }
})
{{ range . -}}
    job("{{.Name}}", function(cenv)
        print("{{.Name}}\n")
        parse_flags(cenv, {
            ip = flag_type.string_type,
        })
        table.dump(cenv)
        optional({
            case(Windows(), function()
                optional({
                    case(is_localhost(cenv.flags["ip"]), function()
                        print("localhost")
                    end),
                    case(not is_localhost(cenv.flags["ip"]), function()
                        print("ssh")
                    end)
                })
            end),
            case(Linux(), function()
                optional({
                    case(is_localhost(cenv.flags["ip"]), function()
                        print("localhost")
                    end),
                    case(not is_localhost(cenv.flags["ip"]), function()
                        print("ssh")
                    end)
                })
            end)
        })
        return true
    end)
{{ end }}