# https://adventofcode.com/2020/day/19

example1 = """
0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"
"""

example2 = """
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
"""

function preprocess_rule(rule)
    if '"' in rule
        return strip(replace(rule, '"' => ""))
    end

    if '|' in rule
        rule = join(map(s -> "(?:$s)", split(rule, '|')), '|')
    end
    rule = replace(rule, r" +" => " ")
    rule = replace(rule, r"(\d+)" => s"(?:\1)")
    rule = "(?:$rule)"
    return rule
end

function parse_input(string)
    rules = Dict()
    for row in split(string, '\n', keepempty=false)
        key, rule = split(row, ":")
        rules[strip(key)] = preprocess_rule(rule)
    end
    return rules
end

function reduce_rule!(rules, name)
    replacement = pop!(rules, name)

    for (key, rule) in rules
        if key == name
            continue
        end
        rules[key] = replace(rule, "(?:$name)" => "(?:$replacement)")
        #rules[key] = replace(rules[key], r"\(\s+([a-z]+)\s+\)" => s"\1")
    end

    return rules
end

function reduce_rules(rules, ignore=["0"])
    select = map(s -> !isnothing(match(r"^[a-z ]+$", s)), values(rules))
    terminal_nodes = collect(keys(rules))[select]

    while length(terminal_nodes) > 0
        value = popfirst!(terminal_nodes)
        reduce_rule!(rules, value)

        select = map(s -> isnothing(match(r"[0-9]", s)), values(rules))
        terminal_nodes = collect(keys(rules))[select]
        terminal_nodes = [x for x in terminal_nodes if !(x in ignore)]
    end

    return rules
end

function remove_extra_characters(str)
    str = replace(str, " " => "")
    str = replace(str, r"\(?:([a-z]+)\)\(?:([a-z]+)\)" => s"\1\2")
    str = replace(str, r"(?<!\|)\(\?:([a-z]+)\)(?!\|)" => s"\1")
    return str
end

reduce_and_parse(input) = remove_extra_characters(reduce_rules(parse_input(input))["0"])

function isvalid(message, rule)
    return !isnothing(match(Regex("^$rule\$"), message))
end

@assert isvalid("aab", reduce_and_parse(example1))
@assert isvalid("aba", reduce_and_parse(example1))
@assert !isvalid("aaa", reduce_and_parse(example1))
@assert !isvalid("abaa", reduce_and_parse(example1))
@assert isvalid("aaaabb", reduce_and_parse(example2))
@assert isvalid("abbabb", reduce_and_parse(example2))
@assert isvalid("aabaab", reduce_and_parse(example2))
@assert isvalid("abaaab", reduce_and_parse(example2))

example3 = """
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
"""

function count_valid_messages(rule, messages)
    return sum(isvalid.(split(messages, '\n'), rule))
end

function part1(input)
    rules, messages = split(input, "\n\n", keepempty=false)
    rule = reduce_and_parse(rules)
    return count_valid_messages(rule, messages)
end

@assert part1(example3) == 2

example4 = """
42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
"""

function part2(input)
    rules, messages = split(input, "\n\n", keepempty=false)
    rules = reduce_rules(parse_input(rules), ["0", "8", "11", "42", "31"])

    # 8: 42 | 42 8
    # 11: 42 31 | 42 11 31

    rules["8"] = "(?:(?:42)+)"
    rules["11"] = "((?:42)(?1)?(?:31))"
    
    for key in ["31", "42", "8", "11"]
        reduce_rule!(rules, key)
    end
    rule = remove_extra_characters(rules["0"])
    matched_messages = count_valid_messages(rule, messages)

    return matched_messages
end

@assert part1(example4) == 3
@assert part2(example4) == 12

test = read("data/day-19.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 168
@assert result2 == 277