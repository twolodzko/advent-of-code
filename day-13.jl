# https://adventofcode.com/2020/day/13

example = "
939
7,13,x,x,59,x,31,19
"

function read_input(string)
  earliest_timestamp, bus_numbers = filter(x -> x != "", split(string, '\n'))
  earliest_timestamp = parse(Int, earliest_timestamp)
  bus_numbers = map(x -> tryparse(Int, x), split(bus_numbers, ','))
  return earliest_timestamp, bus_numbers
end

@assert read_input(example) == (939, [7, 13, nothing, nothing, 59, nothing, 31, 19])

function closest_arrival(bus_number, timestamp)
  divisor, reminder = divrem(timestamp, bus_number)
  if reminder == 0
    return 0
  else
    return bus_number * (divisor + 1) - timestamp
  end
end

function part1(input)
  earliest_timestamp, bus_numbers = read_input(input)
  bus_numbers = filter(x -> !isnothing(x), bus_numbers)
  closest_arrivals = map(x -> closest_arrival(x, earliest_timestamp), bus_numbers)
  i = argmin(closest_arrivals)
  return closest_arrivals[i] * bus_numbers[i]
end

@assert part1(example) == 295

function allequal(numbers)
  for i = 2:length(numbers)
    if numbers[i-1] != numbers[i]
      return false
    end
  end
  return true
end

@assert !allequal([1, 2, 3, 4, 5])
@assert allequal([1, 1, 1, 1, 1])

function all_aligned(timestamps, bus_numbers)
  return all(rem.(timestamps, bus_numbers) .== 0)
end

# TODO: brute-force would not work!
function part2(input)
  _, bus_numbers = read_input(input)

  nonmissing = map(x -> !isnothing(x), bus_numbers)
  indexes = collect(0:length(bus_numbers)-1)

  indexes = indexes[nonmissing]
  indexes = indexes .- minimum(indexes)
  bus_numbers = bus_numbers[nonmissing]

  multiplier, pos = findmax(bus_numbers)
  corrections = indexes .- indexes[pos]

  timestamps = Int[]
  i = 1

  while true
    timestamp = i * multiplier
    timestamps = timestamp .+ corrections

    if all_aligned(timestamps, bus_numbers)
      break
    end

    i += 1
  end

  return timestamps[1]
end

@assert part2(example) == 1068781
@assert part2("0\n17,x,13,19") == 3417
@assert part2("0\n67,7,59,61") == 754018
@assert part2("0\n67,x,7,59,61") == 779210
@assert part2("0\n67,7,x,59,61") == 1261476
@assert part2("0\n1789,37,47,1889") == 1202161486

test = read("data/day-13.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 4207
# @assert part2(test) == 29401
