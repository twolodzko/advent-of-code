# https://adventofcode.com/2020/day/14

import Base: ==

example = "
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
"

const MEM_SIZE = 36

struct Program
  mask::Vector{Union{Nothing,Bool}}
  mem::Vector{Tuple{Integer,Integer}}
end

(==)(x::Program, y::Program) = (x.mask == y.mask) && (x.mem == y.mem)

function parse_input(string)
  programs = Program[]
  mask, mem = [], []
  for (i, row) in enumerate(split(string, '\n', keepempty = false))
    row = strip(row)
    key, value = split(row, '=')
    key, value = strip(key), strip(value)

    if key == "mask"
      if i > 1
        push!(programs, Program(mask, mem))
      end

      mask = map(x -> x == 'X' ? nothing : parse(Bool, x), collect(value))
      @assert length(mask) == MEM_SIZE
      mem = []
    else
      address = parse(Int, match(r"mem\[(\d+)\]", key).captures[1])
      value = parse(Int, value)
      push!(mem, (address, value))
    end
  end

  push!(programs, Program(mask, mem))
  return programs
end

@assert parse_input(example) ==
        [Program([repeat([nothing], 29); 1; repeat([nothing], 4); 0; nothing], [(8, 11), (7, 101), (8, 0)])]

test_program = "
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
mem[0] = 0
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
mem[1] = 1
mem[2] = 2
"

@assert parse_input(test_program) ==
        [Program(repeat([nothing], MEM_SIZE), [(0, 0)]), Program(repeat([nothing], MEM_SIZE), [(1, 1), (2, 2)])]

issomething(x) = !isnothing(x)

function binary_to_decimal(binary::Vector{Bool})
  decimal = 0
  for (i, x) in enumerate(reverse(binary))
    decimal += 2^(i - 1) * x
  end
  return decimal
end

@assert binary_to_decimal(Bool[0, 0, 0]) == 0
@assert binary_to_decimal(Bool[0, 0, 1]) == 1
@assert binary_to_decimal(Bool[0, 1, 0]) == 2
@assert binary_to_decimal(Bool[1, 0, 0, 0, 0, 0, 0]) == 64
@assert binary_to_decimal(Bool[1, 1, 0, 0, 1, 0, 1]) == 101

function write(programs::Vector{Program})
  memory = Dict{Integer, Integer}()
  for program in programs
    for cmd in program.mem
      bits = convert(Vector{Bool}, reverse(digits(cmd[2], base = 2, pad = MEM_SIZE)))
      masked = issomething.(program.mask)
      bits[masked] = program.mask[masked]
      memory[cmd[1]] = binary_to_decimal(bits)
    end
  end
  return memory
end

@assert write(parse_input(example)) == Dict(7 => 101, 8 => 64)

function part1(input)
  program = parse_input(input)
  memory = write(program)
  sum(values(memory))
end

@assert part1(example) == 165

function write_addpresse(programs::Vector{Program})
  memory = Dict{Integer, Integer}()
  for program in programs
    bits = zeros(Bool, MEM_SIZE)
    bits[masked] = program.mask[masked]
    floating = isnothing.(program.mask)
    if any(floating)
      for i in findall(floating)

      end
    else
      value = binary_to_decimal(bits)
      memory[value] = value
    end
  end
  return memory
end

test = read("data/day-14.txt", String)
println("Part 1: $(part1(test))")
# println("Part 2: $(part2(test))")

@assert part1(test) == 6386593869035
# @assert part2(test) == 725850285300475
