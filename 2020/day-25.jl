# https://adventofcode.com/2020/day/25

"""
Single transformation step
"""
function transform(value::Integer, subject_number::Integer)
    return rem(value * subject_number, 20201227)
end

"""
Using the `subject_number` and `loop_size`, calculate the `key`
"""
function get_key(subject_number::Integer, loop_size::Integer)
    value = 1
    for _ in 1:loop_size
        value = transform(value, subject_number)
    end
    return value
end

# example: subject number -> public key
@assert get_key(7, 8) == 5764801
@assert get_key(7, 11) == 17807724
# example: public key -> encryption key
@assert get_key(17807724, 8) == 14897079
@assert get_key(5764801, 11) == 14897079

"""
Using the `subject_number` and `public_key`, guess the `loop_size`
"""
function get_loop_size(subject_number::Integer, public_key::Integer)
    loop_size = 0
    value = 1
    while value != public_key
        loop_size += 1
        value = transform(value, subject_number)
    end
    return loop_size
end

# example: (subject number, public key) -> loop size
@assert get_loop_size(7, 5764801) == 8
@assert get_loop_size(7, 17807724) == 11
# example: (public key, encryption key) -> loop size
@assert get_loop_size(17807724, 14897079) == 8
@assert get_loop_size(5764801, 14897079) == 11

function part1(card_key, door_key)
    card_loop_size = get_loop_size(7, card_key)
    door_loop_size = get_loop_size(7, door_key)

    card_encryption_key = get_key(card_key, door_loop_size)
    door_encryption_key = get_key(door_key, card_loop_size)

    @assert card_encryption_key == door_encryption_key

    return card_encryption_key
end

@assert part1(5764801, 17807724) == 14897079

test = (14205034, 18047856)
println("Part 1: $(result1 = part1(test...))")
# println("Part 2: $(result2 = part2(test...))")

@assert result1 == 297257
# @assert result2 == 418819514477