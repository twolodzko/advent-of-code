
# For example, suppose you have the following list:

# 1-3 a: abcde
# 1-3 b: cdefg
# 2-9 c: ccccccccc

# Each line gives the password policy and then the password. The password policy
# indicates the lowest and highest number of times a given letter must appear for
#     the password to be valid. For example, 1-3 a means that the password must
#     contain a at least 1 time and at most 3 times.

# In the above example, 2 passwords are valid. The middle password, cdefg, is not;
# it contains no instances of b, but needs at least 1. The first and third passwords
# are valid: they contain one a or nine c, both within the limits of their respective
# policies.

function getpattern(row)
    i, j, char, password = match(r"^(\d+)\-(\d+) ([^:]+): (.*)$", row).captures
    return parse(Int32, i), parse(Int32, j), char[1], password
end

function part1(inputs)
    count = 0
    for row in split(inputs, '\n')
        if strip(row) == ""
            continue
        end

        min, max, char, password = getpattern(row)

        # num_found = length([x for x in eachmatch(Regex(char), password)])
        num_found = sum(collect(password) .== char)
        if min <= num_found <= max
            count += 1
        end
    end
    return count
end


str = "
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
"

@assert part1(str) == 2

# Each policy actually describes two positions in the password, where 1 means the first
# character, 2 means the second character, and so on. (Be careful; Toboggan Corporate
# Policies have no concept of "index zero"!) Exactly one of these positions must contain
# the given letter. Other occurrences of the letter are irrelevant for the purposes of
# policy enforcement.

# Given the same example list from above:

#     1-3 a: abcde is valid: position 1 contains a and position 3 does not.
#     1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
#     2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

# How many passwords are valid according to the new interpretation of the policies?

function part2(inputs)
    count = 0
    for row in split(inputs, '\n')
        if strip(row) == ""
            continue
        end

        i, j, char, password = getpattern(row)

        if xor(password[i] == char, password[j] == char)
            count += 1
        end
    end
    return count
end


str = "
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
"

@assert part2(str) == 1
