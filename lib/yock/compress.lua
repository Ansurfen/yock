--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: undefined-field
---@diagnostic disable: param-type-mismatch
---@diagnostic disable: lowercase-global

---@param src string
---@param dst string
tarc = function(src, dst)
    local fw, err = os.Create(dst)
    yassert(err)
    local gw = gzip.NewWriter(fw)
    local tw = tar.NewWriter(gw)
    local root = filepath.Base(src)
    path.walk(src, function(fileName, info, err)
        yassert(err)
        if src == fileName then
            return true
        end
        local hdr, err = tar.FileInfoHeader(info, "")
        yassert(err)
        local relPath, err = filepath.Rel(src, fileName)
        yassert(err)
        hdr.Name = root .. "/" .. strings.TrimPrefix(strings.ReplaceAll(relPath, "\\", "/"), "/")
        yassert(tw:WriteHeader(hdr))
        if not (bit.And(info:Mode(), fs.ModeType) == 0) then
            return true
        end
        local fr, err = os.Open(fileName)
        yassert(err)
        _, err = io.Copy(tw, fr)
        yassert(err)
        fr:Close()
        return true
    end)
    tw:Close()
    gw:Close()
    fw:Close()
end

-- zipc to compress zip of source to specify path
---@param src string
---@param dst string
zipc = function(src, dst)
    if not find(path.dir(dst)) then
        mkdir(path.dir(dst))
    end
    local archive, err = os.Create(dst)
    yassert(err)
    local zw = zip.NewWriter(archive)
    dst = strings.TrimSuffix(src, string.char(path.Separator))
    path.walk(src, function(p, info, err)
        yassert(err)
        local header, err = zip.FileInfoHeader(info)
        yassert(err)
        header.Method = zip.Deflate
        header.Name, err = filepath.Rel(filepath.Dir(dst), p)
        yassert(err)
        if info:IsDir() then
            header.Name = header.Name .. string.char(path.Separator)
        end
        local headerWriter, err = zw:CreateHeader(header)
        yassert(err)
        if info:IsDir() then
            return true
        end
        local fp, err = os.Open(p)
        yassert(err)
        _, err = io.Copy(headerWriter, fp)
        yassert(err)
        fp:Close()
        return true
    end)
    zw:Close()
    archive:Close()
end

---@param src string
---@param dst string
---@return string, err
untar = function(src, dst)
    local file, err = os.Open(src)
    yassert(err)
    local gzipReader, err = gzip.NewReader(file)
    yassert(err)
    local tarReader = tar.NewReader(gzipReader)
    local long_path = ""
    while true do
        local header, err = tarReader:Next()
        if err == io.EOF then
            break
        end
        yassert(err)
        local targetPath = path.join(dst, header.Name)
        if header.Typeflag == tar.TypeDir then
            if #targetPath > #long_path then
                long_path = targetPath
            end
            if not find(targetPath) then
                mkdir(targetPath)
            end
        elseif header.Typeflag == tar.TypeReg then
            local targetPathDir = filepath.Dir(targetPath)
            if not find(targetPathDir) then
                mkdir(targetPathDir)
            end
            local fp, err = os.OpenFile(targetPath, bit.Or(os.O_CREATE, os.O_WRONLY), header:FileInfo():Mode())
            yassert(err)
            _, err = io.Copy(fp, tarReader)
            yassert(err)
            fp:Close()
        else
            print("invalid file type")
        end
    end
    file:Close()
    local rel, err = filepath.Rel(dst, long_path)
    if err ~= nil then
        return rel, err
    end
    local idx = strings.IndexAny(rel, string.char(filepath.Separator))
    return string.sub(rel, 1, idx), nil
end

---@param src string
---@param dst string
---@return string, err
unzip = function(src, dst)
    local reader, err = zip.OpenReader(src)
    yassert(err)
    local long_path = ""
    for i = 1, #reader.File, 1 do
        local file = reader.File[i]
        local filePath = path.join(dst, file.Name)
        if file:FileInfo():IsDir() then
            if #filePath > #long_path then
                long_path = filePath
            end
            if not find(filePath) then
                mkdir(filePath)
            end
            goto continue
        end
        local filePathDir = path.dir(filePath)
        if not find(filePathDir) then
            mkdir(filePathDir)
        end
        local rc, err = file:Open()
        yassert(err)
        local w, err = os.Create(filePath)
        yassert(err)
        _, err = io.Copy(w, rc)
        yassert(err)
        w:Close()
        rc:Close()
        ::continue::
    end
    local rel, err = filepath.Rel(dst, long_path)
    if err ~= nil then
        return rel, err
    end
    local idx = strings.IndexAny(rel, string.char(filepath.Separator))
    return string.sub(rel, 1, idx), nil
end

---@param src string
---@param dst string
compress = function(src, dst)
    local ext = filepath.Ext(dst)
    if ext == ".zip" then
        zipc(src, dst)
    elseif ext == ".gz" then
        tarc(src, dst)
    else
        yassert("no support the compress type")
    end
end

---@param src string
---@param dst string
---@return string, err
uncompress = function(src, dst)
    local ext = filepath.Ext(src)
    if ext == ".zip" then
        return unzip(src, dst)
    elseif ext == ".gz" then
        return untar(src, dst)
    else
        yassert("no support the uncompress type")
    end
    return "", nil
end
