# https://adventofcode.com/2020/day/14

import Base: ==

example1 = "
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
"

const MEM_SIZE = 36

BitMask = Vector{Union{Nothing,Bool}}
Binary = Vector{Union{Bool}}

Base.repr(obj::BitMask) = join([Dict(nothing => 'X', true => 1, false => 0)[x] for x in obj])

struct Program
    mask::BitMask
    mem::Vector{Tuple{Integer,Integer}}
end

(==)(x::Program, y::Program) = (x.mask == y.mask) && (x.mem == y.mem)

function parse_input(string)
    programs = Program[]
    mask, mem = [], []
    for (i, row) in enumerate(split(string, '\n', keepempty=false))
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

@assert parse_input(example1) ==
        [Program([repeat([nothing], 29); 1; repeat([nothing], 4); 0; nothing], [(8, 11), (7, 101), (8, 0)])]
@assert repr(parse_input(example1)[1].mask) == "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"

test_program = "
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
mem[0] = 0
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
mem[1] = 1
mem[2] = 2
"

@assert parse_input(test_program) ==
        [Program(repeat([nothing], MEM_SIZE), [(0, 0)]), Program(repeat([nothing], MEM_SIZE), [(1, 1), (2, 2)])]

function binary_to_decimal(binary::Binary)
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

function apply_mask(bits::Binary, mask::BitMask)::Binary
    idx = map(x -> !isnothing(x), mask)
    bits[idx] = mask[idx]
    return bits
end

@assert apply_mask(Bool[0, 1, 0, 1], [nothing, nothing, nothing, true]) == Bool[0, 1, 0, 1]
@assert apply_mask(Bool[0, 0, 0, 0], [nothing, true, nothing, false]) == Bool[0, 1, 0, 0]
@assert apply_mask(Bool[0, 0, 0, 0], [true, true, true, nothing]) == Bool[1, 1, 1, 0]
@assert apply_mask(Bool[0, 1, 0, 1], [true, nothing, true, nothing]) == Bool[1, 1, 1, 1]

function tobinary(value::Integer)::Binary
    return convert(Vector{Bool}, reverse(digits(value, base=2, pad=MEM_SIZE)))
end

@assert tobinary(42) ==
        Bool[0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0]

function write(programs::Vector{Program})
    memory = Dict{Integer,Integer}()
    for program in programs
        for cmd in program.mem
            bits = tobinary(cmd[2])
            bits = apply_mask(bits, program.mask)
            memory[cmd[1]] = binary_to_decimal(bits)
        end
    end
    return memory
end

@assert write(parse_input(example1)) == Dict(7 => 101, 8 => 64)

function part1(input)
    program = parse_input(input)
    memory = write(program)
    sum(values(memory))
end

@assert part1(example1) == 165

"""
Apply mask:

```
address: 000000000000000000000000000000101010  (decimal 42)
mask:    000000000000000000000000000000X1001X
result:  000000000000000000000000000000X1101X
```
"""
function apply_xmask(bits::Binary, mask::BitMask)::BitMask
    @assert length(bits) == length(mask)
    bits = convert(BitMask, bits)
    bits[mask.==true] .= true
    bits[isnothing.(mask)] .= nothing
    return bits
end

@assert apply_xmask(tobinary(42), [repeat([false], 30); nothing; true; false; false; true; nothing]) ==
        [repeat([false], 30); nothing; true; true; false; true; nothing]

Base.repr(obj::Binary) = join([Dict(true => 1, false => 0)[x] for x in obj])

function write_ver2(programs::Vector{Program}; verbose=false)
    memory = Dict{Integer,Integer}()
    for program in programs
        for cmd in program.mem
            bits = tobinary(cmd[1])

            verbose && println("address: $(repr(bits)) (decimal $(binary_to_decimal(bits)))")
            verbose && println("mask:    $(repr(program.mask))")

            bits = apply_xmask(bits, program.mask)
            verbose && println("result:  $(repr(bits))\n")

            if any(isnothing.(bits))
                idx = findall(isnothing, bits)
                for replacement in Iterators.product(repeat([0:1], length(idx))...)
                    bits[idx] .= convert(Vector{Bool}, collect(replacement))
                    address = binary_to_decimal(convert(Binary, bits))
                    memory[address] = cmd[2]
                    verbose && println("$(join(convert(Vector{Int}, bits))) (decimal $(address))")
                end
            else
                address = binary_to_decimal(bits)
                memory[address] = cmd[2]
                verbose && println("$(join(convert(Vector{Int}, bits))) (decimal $(address))")
            end
        end
        verbose && println()
    end
    return memory
end

example2 = "
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
"

function part2(input)
    program = parse_input(input)
    memory = write_ver2(program)
    sum(values(memory))
end

@assert part2(example2) == 208

test = read("data/day-14.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 6386593869035
@assert result2 == 4288986482164
