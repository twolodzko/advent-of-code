# https://adventofcode.com/2020/day/12

#  * Action `N` means to move north by the given value.
#  * Action `S` means to move south by the given value.
#  * Action `E` means to move east by the given value.
#  * Action `W` means to move west by the given value.
#  * Action `L` means to turn left the given number of degrees.
#  * Action `R` means to turn right the given number of degrees.
#  * Action `F` means to move forward by the given value in the
#    direction the ship is currently facing.

example = "
F10
N3
F7
R90
F11
"

const directions = (:North, :East, :South, :West)
const char_to_direction = Dict("N" => :North, "E" => :East, "S" => :South, "W" => :West)

Point = Tuple{Integer,Integer}

mutable struct Ship
  facing::Symbol
  position::Point
  waypoint::Point
end

facing(ship::Ship) = ship.facing
position(ship::Ship) = ship.position
waypoint(ship::Ship) = ship.waypoint

function rotate(facing::Symbol, direction::Symbol, deg::Integer = 90)
  steps, reminder = divrem(deg, 90)
  reminder == 0 || error("ups, it can move diagonally!")
  direction = direction == :Right ? +1 : -1
  direction_index = findfirst(directions .== facing)
  direction_index = mod1(direction_index + steps * direction, 4)
  return directions[direction_index]
end

@assert rotate(:East, :Left) == :North
@assert rotate(:North, :Left) == :West
@assert rotate(:North, :Right) == :East
@assert rotate(:East, :Right) == :South
@assert rotate(:North, :Left, 360) == :North

function part1(string::AbstractString; verbose = false)
  ship = Ship(:East, (0, 0), (0, 0))

  for row in split(string, '\n')
    row = strip(row)
    if row == ""
      continue
    end
    m = match(r"([NSEWLRF])(\d+)", row)
    if isnothing(m)
      error("invalid input")
    end
    action, value = m.captures
    value = parse(Int, value)

    if action in ("L", "R")
      direction = rotate(facing(ship), action == "L" ? :Left : :Right, value)
      ship.facing = direction
      value = 0
    elseif action == "F"
      direction = facing(ship)
    else
      direction = char_to_direction[action]
    end

    x, y = position(ship)
    if direction == :North
      x += value
    elseif direction == :South
      x -= value
    elseif direction == :East
      y += value
    else
      y -= value
    end
    ship.position = (x, y)

  end
  return position(ship) .|> abs |> sum
end

@assert part1(example) == 25

function rotate(point::Point, deg::Integer)::Point
  θ = deg2rad(deg)
  x, y = point
  x_new = cos(θ) * x - sin(θ) * y
  y_new = sin(θ) * x + cos(θ) * y
  return Int(round(x_new)), Int(round(y_new))
end

@assert rotate((10, 4), 90) == (-4, 10)

function part2(string::AbstractString; verbose = false)
  ship = Ship(:East, (0, 0), (1, 10))

  for row in split(string, '\n')
    row = strip(row)
    if row == ""
      continue
    end
    m = match(r"([NSEWLRF])(\d+)", row)
    if isnothing(m)
      error("invalid input")
    end
    action, value = m.captures
    value = parse(Int, value)

    if action in ("N", "S", "W", "E")
      x, y = waypoint(ship)
      if action == "N"
        x += value
      elseif action == "S"
        x -= value
      elseif action == "E"
        y += value
      else
        y -= value
      end
      ship.waypoint = (x, y)
    elseif action in ("L", "R")
      if action == "L"
        value = 360 - value
      end
      ship.waypoint = rotate(waypoint(ship), value)
    else
      wx, wy = waypoint(ship)
      px, py = position(ship)
      ship.position = (px + wx * value, py + wy * value)
    end

    verbose && println("$(row):\t Ship at $(position(ship)), with waypoint $(waypoint(ship))")

  end
  return position(ship) .|> abs |> sum
end

@assert part2(example) == 286

test = read("data/day-12.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 415
@assert part2(test) == 29401
