--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field

yockw                                = {
    metrics = {
        counter = {},
        gauge = {},
        histogram = {},
        summary = {},
        counter_vec = {},
        gauge_vec = {},
        histogram_vec = {},
        summary_vec = {}
    }
}

local dial_yockw                     = function(router, handle, ret)
    -- if env.conf.Yockw.SelfBoot then
    local res, err = curl(handle, strf("http://localhost:%d%s", env.conf.Yockw.Port, router))
    if type(ret) == "boolean" and ret then
        return res, err
    else
        if err == nil then
            ycho.info(res)
        else
            ycho.error(string.format("%s %s", res, err))
        end
    end
    -- end
end

---@param opt option_yockw_metrics_counter
yockw.metrics.counter.new            = function(opt)
    dial_yockw("/metrics/counter/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or ""
        }),
    })
end

---@param name string
---@param f number
yockw.metrics.counter.add            = function(name, f)
    dial_yockw("/metrics/counter/add", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) }
        }),
    })
end

---@param name string
yockw.metrics.counter.inc            = function(name)
    dial_yockw("/metrics/counter/inc", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@param name string
yockw.metrics.counter.rm             = function(name)
    dial_yockw("/metrics/counter/rm", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@return string[]
yockw.metrics.counter.ls             = function()
    local res, err = dial_yockw("/metrics/counter/ls", { method = "GET" }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end

yockw.metrics.counter_vec.new        = function(opt)
    dial_yockw("/metrics/counterVec/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or "",
            labels = opt["label"] or ""
        }),
    })
end

yockw.metrics.counter_vec.add        = function(name, f, labels)
    dial_yockw("/metrics/counterVec/add", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) },
            label = { json.encode(labels) }
        }),
    })
end

yockw.metrics.counter_vec.rm         = function(name)
    dial_yockw("/metrics/counterVec/rm", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

yockw.metrics.counter_vec.ls         = function()
    local res, err = dial_yockw("/metrics/counterVec/ls", { method = "GET" }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end

---@param opt option_yockw_metrics_gauge
yockw.metrics.gauge.new              = function(opt)
    dial_yockw("/metrics/gauge/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or ""
        }),
    })
end

---@param name string
---@param f number
yockw.metrics.gauge.add              = function(name, f)
    dial_yockw("/metrics/gauge/add", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) }
        }),
    })
end

---@param name string
---@param f number
yockw.metrics.gauge.sub              = function(name, f)
    dial_yockw("/metrics/gauge/sub", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) }
        }),
    })
end

---@param name string
yockw.metrics.gauge.inc              = function(name)
    dial_yockw("/metrics/gauge/inc", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@param name string
yockw.metrics.gauge.dec              = function(name)
    dial_yockw("/metrics/gauge/dec", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@param name string
---@param f number
yockw.metrics.gauge.set              = function(name, f)
    dial_yockw("/metrics/gauge/set", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) }
        }),
    })
end

---@param name string
yockw.metrics.gauge.rm               = function(name)
    dial_yockw("/metrics/gauge/rm", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@param name string
yockw.metrics.gauge.setToCurrentTime = function(name)
    dial_yockw("/metrics/gauge/setToCurrentTime", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@param name string
yockw.metrics.gauge.rm               = function(name)
    dial_yockw("/metrics/gauge/rm", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    })
end

---@return string[]
yockw.metrics.gauge.ls               = function(name)
    local res, err = dial_yockw("/metrics/gauge/ls", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name }
        }),
    }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end


---@param opt option_yockw_metrics_gauge
yockw.metrics.gauge_vec.new              = function(opt)
    dial_yockw("/metrics/gaugeVec/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or "",
            labels = opt["label"] or ""
        }),
    })
end

---@param name string
---@param f number
yockw.metrics.gauge_vec.add              = function(name, f, label)
    dial_yockw("/metrics/gauge/add", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) },
            labels = { json.encode(label) }
        }),
    })
end

---@param name string
---@param f number
---@param label string[]|table<string,string>
yockw.metrics.gauge_vec.sub              = function(name, f, label)
    dial_yockw("/metrics/gaugeVec/sub", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) },
            labels = { json.encode(label) }
        }),
    })
end

---@param name string
yockw.metrics.gauge_vec.inc              = function(name, label)
    dial_yockw("/metrics/gaugeVec/inc", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            labels = { json.encode(label) }
        }),
    })
end

---@param name string
yockw.metrics.gauge_vec.dec              = function(name, label)
    dial_yockw("/metrics/gaugeVec/dec", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            labels = { json.encode(label) }
        }),
    })
end

---@param name string
---@param f number
yockw.metrics.gauge_vec.set              = function(name, f, label)
    dial_yockw("/metrics/gaugeVec/set", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) },
            labels = { json.encode(label) }
        }),
    })
end

---@param name string
yockw.metrics.gauge_vec.rm               = function(name)
    dial_yockw("/metrics/gaugeVec/rm", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
        }),
    })
end

---@param name string
---@param label string[]|table<string, string>
yockw.metrics.gauge_vec.setToCurrentTime = function(name, label)
    dial_yockw("/metrics/gaugeVec/setToCurrentTime", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            labels = { json.encode(label) }
        }),
    })
end

---@return string[]
yockw.metrics.gauge_vec.ls               = function()
    local res, err = dial_yockw("/metrics/gaugeVec/ls", { method = "GET" }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end

yockw.metrics.histogram.new              = function(opt)
    dial_yockw("/metrics/histogram/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or "",
            buckets = opt["buckets"] or {}
        })
    })
end

yockw.metrics.histogram.observe          = function(name, f)
    dial_yockw("/metrics/histogram/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) }
        }),
    })
end

yockw.metrics.histogram.rm               = function(name)
    dial_yockw("/metrics/histogram/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
        }),
    })
end

yockw.metrics.histogram.ls               = function()
    local res, err = dial_yockw("/metrics/histogram/ls", {
        method = "GET"
    }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end

yockw.metrics.histogram_vec.new          = function(opt)
    dial_yockw("/metrics/histogramVec/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or "",
            labels = opt["label"] or "",
            buckets = opt["buckets"] or {}
        })
    })
end

yockw.metrics.histogram_vec.observe      = function(name, f, label)
    dial_yockw("/metrics/histogramVec/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) },
            labels = { json.encode(label) }
        }),
    })
end

yockw.metrics.histogram_vec.rm           = function(name)
    dial_yockw("/metrics/histogramVec/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
        }),
    })
end

yockw.metrics.histogram_vec.ls           = function()
    local res, err = dial_yockw("/metrics/histogram/ls", {
        method = "GET"
    }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end

yockw.metrics.summary.new                = function(opt)
    dial_yockw("/metrics/summary/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or "",
            objectives = opt["objectives"] or {}
        })
    })
end

yockw.metrics.summary.observe            = function(name, f)
    dial_yockw("/metrics/summary/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) }
        }),
    })
end

yockw.metrics.summary.rm                 = function(name)
    dial_yockw("/metrics/summary/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
        }),
    })
end

yockw.metrics.summary.ls                 = function()
    local res, err = dial_yockw("/metrics/summary/ls", {
        method = "GET"
    }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end

yockw.metrics.summary_vec.new            = function(opt)
    dial_yockw("/metrics/summaryVec/new", {
        method = "POST",
        data = json.encode({
            name = opt["name"] or "",
            help = opt["help"] or "",
            namespace = opt["namespace"] or "",
            subsystem = opt["subsystem"] or "",
            labels = opt["label"] or "",
            objectives = opt["objectives"] or {}
        })
    })
end

yockw.metrics.summary.observe            = function(name, f, label)
    dial_yockw("/metrics/summary/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
            f = { tostring(f) },
            labels = { json.encode(label) }
        }),
    })
end

yockw.metrics.summary_vec.rm             = function(name)
    dial_yockw("/metrics/summaryVec/observe", {
        method = "POST",
        header = {
            ["Content-Type"] = "application/x-www-form-urlencoded"
        },
        data = formdata.encode({
            name = { name },
        }),
    })
end

yockw.metrics.summary_vec.ls             = function()
    local res, err = dial_yockw("/metrics/summaryVec/ls", {
        method = "GET"
    }, true)
    if err ~= nil then
        return {}
    end
    return json.decode(res)
end
