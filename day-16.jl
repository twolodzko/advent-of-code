# https://adventofcode.com/2020/day/16

# You collect the rules for ticket fields, the numbers on your ticket, and the numbers on other
# nearby tickets for the same train service (via the airport security cameras) together into
# a single document you can reference (your puzzle input).
#
# The rules for ticket fields specify a list of fields that exist somewhere on the ticket and
# the valid ranges of values for each field. For example, a rule like class: 1-3 or 5-7 means
# that one of the fields in every ticket is named class and can be any value in the ranges
# 1-3 or 5-7 (inclusive, such that 3 and 5 are both valid in this field, but 4 is not).
#
# Start by determining which tickets are completely invalid; these are tickets that contain
# values which aren't valid for any field. Ignore your ticket for now.

example1 = "
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
"

split_lines(str) = split(str, '\n', keepempty=false)

function parse_range(string::AbstractString)::AbstractRange
    from, to = parse.(Int, split(strip(string), '-', keepempty=false))
    return from:to
end

@assert parse_range(" 9-13233 ") == 9:13233

function parse_rule(string::AbstractString)
    m = match(r"^([a-z ]*): ([\d or-]*)$", string)
    if isnothing(m)
        error("parsing error")
    end
    name, rule_string = m.captures
    ranges = map(parse_range, split(rule_string, "or", keepempty=false))
    return name, ranges
end

@assert parse_rule("row: 6-11 or 33-44") == ("row", [6:11, 33:44])
@assert parse_rule("foo: 1-100") == ("foo", [1:100])
@assert parse_rule("bar: 6-11 or 33-44 or 55-66 ") == ("bar", [6:11, 33:44, 55:66])

function parse_ticket(string::AbstractString)
    return parse.(Int, split(string, ',', keepempty=false))
end

function parse_input(input::AbstractString)
    rules, my_ticket, nearby_tickets = split(input, "\n\n", keepempty=false)
    rules = Dict(map(parse_rule, split_lines(rules)))
    my_ticket = parse_ticket(split_lines(my_ticket)[2])
    nearby_tickets = map(parse_ticket, split_lines(nearby_tickets)[2:end])
    return rules, my_ticket, nearby_tickets
end

function collect_rules(rules::AbstractDict)
    ranges = []
    for rule in values(rules)
        for range in rule
            push!(ranges, range)
        end
    end
    return ranges
end

@assert collect_rules(parse_input(example1)[1]) == [1:3, 5:7, 6:11, 33:44, 13:40, 45:50]

function invalid_fields(ticket, valid_ranges)
    values = Integer[]
    for value in ticket
        valid = false
        for range in valid_ranges
            valid |= (value in range)
        end
        !valid && push!(values, value)
    end
    return values
end

@assert invalid_fields([7, 3, 47], [1:3, 5:7, 6:11, 33:44, 13:40, 45:50]) == []
@assert invalid_fields([40, 4, 50], [1:3, 5:7, 6:11, 33:44, 13:40, 45:50]) == [4]
@assert invalid_fields([55, 2, 20], [1:3, 5:7, 6:11, 33:44, 13:40, 45:50]) == [55]
@assert invalid_fields([38, 6, 12], [1:3, 5:7, 6:11, 33:44, 13:40, 45:50]) == [12]

function part1(input)
    rules, my_ticket, nearby_tickets = parse_input(input)
    valid_ranges = collect_rules(rules)

    invalid = Integer[]
    for ticket in nearby_tickets
        append!(invalid, invalid_fields(ticket, valid_ranges))
    end
    return sum(invalid)
end

@assert part1(example1) == 71

example2 = "
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
"

function isvalid(value, rule)
    for range in rule
        if value in range
            return true
        end
    end
    return false
end

@assert isvalid(9, [0:13, 16:19])
@assert !isvalid(15, [0:13, 16:19])

function trypop!(collection, key)
    try
        pop!(collection, key)
    catch KeyError
    end
    return collection
end

@assert trypop!(Set([:a, :b, :c]), :c) == Set([:a, :b])
@assert trypop!(Set([:a, :b, :c]), :d) == Set([:a, :b, :c])

function guess_fields(rules, nearby_tickets)
    number_of_fields = length(rules)
    valid_ranges = collect_rules(rules)

    valid_tickets = filter(ticket -> isempty(invalid_fields(ticket, valid_ranges)), nearby_tickets)
    @assert !isempty(valid_tickets)

    # per each field, leave only the rules that are consistent with the data
    field_candidates = [Set(keys(rules)) for _ = 1:number_of_fields]
    for ticket in valid_tickets
        for (i, field) in enumerate(ticket)
            for (name, rule) in rules
                if !isvalid(field, rule)
                    pop!(field_candidates[i], name)
                end
            end
        end
    end

    # eliminate the rules, assumming that there is a unique mapping
    field_names = Array{String,1}(undef, number_of_fields)
    for _ = 1:number_of_fields
        pos = findfirst(map(length, field_candidates) .== 1)
        name = first(field_candidates[pos])  # extract the only element
        field_names[pos] = name
        map(field -> trypop!(field, name), field_candidates)
    end
    return field_names
end

@assert guess_fields(parse_input(example2)[[1, 3]]...) == ["row", "class", "seat"]

function part2(input)
    rules, my_ticket, nearby_tickets = parse_input(input)
    ticket_fields = guess_fields(rules, nearby_tickets)

    result = 1
    for (i, name) in enumerate(ticket_fields)
        if startswith(name, "departure")
            result *= my_ticket[i]
        end
    end
    return result
end

test = read("data/day-16.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 19093
@assert result2 == 5311123569883
