
# https://adventofcode.com/2020/day/7

example1 = "
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
"

struct Bag
    kind::AbstractString
    quantity::Integer

    Bag(kind, quantity) = new(kind, quantity)

    function Bag(string::AbstractString)
        m = match(r"(\d*?) ?([a-z ]+?) (?:bag|bags)", string)
        if isnothing(m)
            error("invalid bag description")
        end
        quantity, kind = m.captures

        if kind == "no other"
            return nothing
        end
        return new(String(kind), quantity == "" ? 1 : parse(Int, quantity))
    end
end

function rule_parser(rule)
    bag, contains = split(rule, " contain ", limit=2)
    bag = Bag(bag)
    contains = [Bag(elem) for elem in split(contains, ',')]
    return bag, filter(x -> !isnothing(x), contains)
end

function rules_to_revdict(input)
    bags = Dict()

    for row in split(input, '\n', keepempty=false)
        bag, contains = rule_parser(row)

        for elem in contains
            if isnothing(elem)
                continue
            end

            if elem.kind in keys(bags)
                push!(bags[elem.kind], bag.kind)
            else
                bags[elem.kind] = [bag.kind]
            end
        end
    end

    return bags
end

@assert rules_to_revdict(example1)["shiny gold"] == ["bright white", "muted yellow"]

function search_for_possible_bag_colors(bags, target="shiny gold")
    possible_colors = Set()
    to_search = [target]

    while !isempty(to_search)
        found = []
        for elem in to_search
            try
                # we already have those on the list, ignore in case there are cycles
                for x in bags[elem]
                    push!(found, x)
                    push!(possible_colors, x)
                end
            catch KeyError
            end

        end
        to_search = found
    end
    return possible_colors
end

function part1(input)
    bags = rules_to_revdict(input)
    possible_colors = search_for_possible_bag_colors(bags)
    return length(possible_colors)
end

@assert part1(example1) == 4

function rules_to_dict(input)
    rules = Dict{String,Array{Bag}}()
    for row in split(input, '\n', keepempty=false)
        bag, contains = rule_parser(row)
        rules[bag.kind] = contains
    end
    return rules
end

@assert rules_to_dict(example1)["shiny gold"] == Bag[Bag("dark olive", 1), Bag("vibrant plum", 2)]

function count_bags(rules, bag)
    if isempty(rules[bag.kind])
        return bag.quantity
    else
        n = 0
        for elem in rules[bag.kind]
            n += count_bags(rules, elem)
        end
        return bag.quantity + bag.quantity * n
    end
end

@assert count_bags(rules_to_dict(example1), Bag("faded blue", 1)) == 1
@assert count_bags(rules_to_dict(example1), Bag("faded blue", 3)) == 3
@assert count_bags(rules_to_dict(example1), Bag("vibrant plum", 1)) == 12
@assert count_bags(rules_to_dict(example1), Bag("vibrant plum", 2)) == 24

function part2(input)
    rules = rules_to_dict(input)
    return count_bags(rules, Bag("shiny gold", 1)) - 1
end

example2 = "
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
"

@assert part2(example2) == 126

test = read("data/day-07.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 222
@assert result2 == 13264
