require "day-15"

local COVERED = "#"

local function markcoverage(map, sensor)
    local function mark(x, y)
        if not map[x] then
            map[x] = {}
        end
        if not map[x][y] then
            map[x][y] = COVERED
        end
    end

    local d = sensor.range
    for y = sensor.position.y - d, sensor.position.y + d do
        local minx, maxx = sensor:yrange(y)
        for x = minx, maxx do
            mark(x, y)
        end
    end
end

local function limits(map)
    local minx = math.maxinteger
    local maxx = math.mininteger
    local miny = math.maxinteger
    local maxy = math.mininteger
    for x, _ in pairs(map) do
        if x < minx then
            minx = x
        end
        if x > maxx then
            maxx = x
        end
        for y, _ in pairs(map[x]) do
            if y < miny then
                miny = y
            end
            if y > maxy then
                maxy = y
            end
        end
    end
    return minx, maxx, miny, maxy
end

local function maptostr(map)
    local minx, maxx, miny, maxy = limits(map)
    local result = ""
    for y = miny, maxy do
        local line = string.format("%2d ", y)
        for x = minx, maxx do
            local value = map[x] and map[x][y]
            if not value then
                line = line .. "."
            else
                line = line .. value
            end
        end
        result = result .. line .. "\n"
    end
    return result
end

do
    local sensors = parse(example1)
    local map = makemap(sensors)
    print(maptostr(map))
    print()

    markcoverage(map, sensors[7])
    print(maptostr(map))
    print()
end
