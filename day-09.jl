
# https://adventofcode.com/2020/day/9

# XMAS starts by transmitting a preamble of 25 numbers. After that, each number
# you receive should be the sum of any two of the 25 immediately previous numbers.
# The two numbers will have different values, and there might be more than one
# such pair.

function is_xmas_seq(numbers; preamble_size = 25)
  number = numbers[end]
  preamble = sort(numbers[(end-preamble_size):(end-1)], rev = true)
  @assert length(preamble) == preamble_size

  for (i, x) in enumerate(preamble)
    if x > number
      continue
    end
    if (number - x) in preamble[(i+1):end]
      return true
    end
  end
  return false
end

@assert is_xmas_seq([1:25; 26])
@assert is_xmas_seq([1:25; 6])
@assert is_xmas_seq([1:25; 49])
@assert !is_xmas_seq([1:25; 100])
@assert !is_xmas_seq([1:25; 50])
@assert !is_xmas_seq([1:25; 2])

function read_array(string::AbstractString)::Vector{Int}
  return map(x -> parse(Int, x), filter(x -> x != "", split(string, '\n')))
end

example = read_array("
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
")

function find_first_invalid(numbers; preamble_size = 25)
  for pos = (preamble_size+1):length(numbers)
    if !is_xmas_seq(numbers[1:pos], preamble_size = preamble_size)
      return numbers[pos], pos
    end
  end
end

@assert find_first_invalid([1:25; 100]) == (100, 26)
@assert find_first_invalid([1:25; 27; 100]) == (100, 27)
@assert find_first_invalid(example, preamble_size = 5) == (127, 15)

function part1(numbers)
  number, _ = find_first_invalid(numbers)
  return number
end

function sum_until(numbers, target)
  tot = numbers[1]
  for i = 2:length(numbers)
    tot += numbers[i]
    if tot == target
      return numbers[1:i]
    end
  end
  return nothing
end

@assert sum_until(1:10, 10) == [1:4;]
@assert sum_until(1:10, 100) === nothing

function find_weak_sequence(numbers; preamble_size = 25)
  number, pos = find_first_invalid(numbers, preamble_size = preamble_size)
  for (i, x) in enumerate(numbers)
    if i == pos
      break
    end
    weak_sequence = sum_until(numbers[i:(pos-1)], number)
    if weak_sequence !== nothing
      return weak_sequence
    end
  end
end

@assert find_weak_sequence(example, preamble_size = 5) == [15, 25, 47, 40]

function part2(numbers; preamble_size = 25)
  weak_sequence = find_weak_sequence(numbers, preamble_size = preamble_size)
  return minimum(weak_sequence) + maximum(weak_sequence)
end

@assert part2(example, preamble_size = 5) == 62

test = read_array(read("data/day-09.txt", String))
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 25918798
@assert part2(test) == 3340942
