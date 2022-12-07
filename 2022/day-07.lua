require "common"

Dir = { type = "dir" }

function Dir:new(parent, content)
    local dir = { parent = parent, content = content or {} }
    setmetatable(dir, self)
    self.__index = self
    return dir
end

function Dir:size()
    local total = 0
    if self.content then
        for _, obj in pairs(self.content) do
            total = total + obj:size()
        end
    end
    return total
end

function Dir:add(name, obj)
    if not self.content[name] then
        self.content[name] = obj
    end
end

File = { type = "file" }

function File:new(size)
    local file = { _size = size }
    setmetatable(file, self)
    self.__index = self
    return file
end

function File:size()
    return self._size
end

assert(File:new(6):size() == 6)
assert(Dir:new():size() == 0)
assert(Dir:new(nil, { a = File:new(5), b = File:new(3) }):size() == 8)
assert(Dir:new(nil, { a = File:new(5), b = Dir:new(nil, { c = File:new(2) }), d = Dir:new() }):size() == 7)

local example1 = [[
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
]]

local function parse(input)
    local root = Dir:new()
    local current = root

    for line in lines(input) do
        local cmd, pos = string.match(line, "$ (%g+)()")
        if line == "" or cmd == "ls" then
            -- skip
        elseif cmd == "cd" then
            local arg = string.match(line, "(%g+).*", pos)
            if arg == ".." then
                current = current.parent
            elseif arg == "/" then
                current = root
            else
                current = current.content[arg]
            end
        else
            local desc, name = string.match(line, "(%g+) (%g+)")
            if desc == "dir" then
                current:add(name, Dir:new(current))
            else
                current:add(name, File:new(tonumber(desc)))
            end
        end
    end
    return root
end

assert(parse(example1):size() == 48381165)

local function traverse1(tree, maxsize)
    local total = 0
    if tree.content then
        for _, obj in pairs(tree.content) do
            if obj.type == "dir" then
                local size = obj:size()
                if size <= maxsize then
                    total = total + size
                end
                total = total + traverse1(obj, maxsize)
            end
        end
    end
    return total
end

function problem1(input)
    local root = parse(input)
    return traverse1(root, 100000)
end

assert(problem1(example1) == 95437)
print(problem1(readfile("data/day-07.txt")))

local function neededspace(dir)
    return 30000000 - (70000000 - dir:size())
end

assert(neededspace(parse(example1)) == 8381165)

local function smallest(...)
    local current
    for _, v in pairs({ ... }) do
        if not current or v and v < current then
            current = v
        end
    end
    return current
end

local function traverse2(tree, minsize)
    local candidate
    if tree.content then
        for _, obj in pairs(tree.content) do
            local size = obj:size()
            if obj.type == "dir" and size >= minsize then
                candidate = smallest(
                    candidate, size, traverse2(obj, minsize))
            end
        end
    end
    return candidate
end

function problem2(input)
    local root = parse(input)
    local tofree = neededspace(root)
    return traverse2(root, tofree)
end

assert(problem2(example1) == 24933642)
print(problem2(readfile("data/day-07.txt")))
