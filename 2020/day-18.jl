# https://adventofcode.com/2020/day/18

function parse_expression(chars::Vector{Char})
    expression = []
    while !isempty(chars)
        ch = popfirst!(chars)

        if ch == ' '
            continue
        elseif isdigit(ch)
            digit = parse(Int, ch)
            if !isempty(expression) && isa(expression[end], Number)
                expression[end] = 10 * expression[end] + digit
            else
                push!(expression, digit)
            end
        elseif ch == '*'
            push!(expression, :*)
        elseif ch == '+'
            push!(expression, :+)
        elseif ch == '('
            partial, chars = parse_expression(chars)
            push!(expression, partial)
        elseif ch == ')'
            return expression, chars
        else
            error("invalid symbol: '$(ch)'")
        end
    end
    return expression, chars
end

function parse_expression(string::AbstractString)
    expression, _ = parse_expression(collect(string))
    return expression
end

@assert parse_expression("1 + 2") == [1, :+, 2]
@assert parse_expression("1 + (2 * 3)") == [1, :+, [2, :*, 3]]

function evaluate(expression)
    result = 0
    op = (+)
    for x in expression
        if isa(x, Symbol)
            op = (x == :+) ? (+) : (*)
            continue
        elseif isa(x, Vector)
            x = evaluate(x)
        end
        result = op(result, x)
    end
    return result
end

parse_eval(expr) = evaluate(parse_expression(expr))

@assert parse_eval("1 + (2 * 3) + (4 * (5 + 6))") == 51
@assert parse_eval("2 * 3 + (4 * 5)") == 26
@assert parse_eval("5 + (8 * 3 + 9 + 3 * 4 * 3)") == 437
@assert parse_eval("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))") == 12240
@assert parse_eval("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2") == 13632

function part1(input)
    return reduce(+, map(parse_eval, split(input, '\n', keepempty=false)))
end

function bracket(expr)
    new_expr = []
    i = 1
    while i <= length(expr)
        x = expr[i]
        if isa(x, Vector)
            x = bracket(x)
        end

        if x == :+
            y = expr[i + 1]
            if isa(y, Vector)
                y = bracket(y)
            end

            new_expr[end] = [new_expr[end], :+, y]
            i += 2
        else
            push!(new_expr, x)
            i += 1
        end
    end
    return length(new_expr) == 1 ? first(new_expr) : new_expr
end

@assert bracket([2, :*, 3, :+, 4]) == [2, :*, [3, :+, 4]]
@assert bracket([2, :*, [3, :*, 4], :+, 5]) == [2, :*, [[3, :*, 4], :+, 5]]
@assert bracket([2, :+, 3, :+, 4, :*, 5]) == [[[2, :+, 3], :+, 4], :*, 5]
@assert bracket([5, :+, [8, :*, 3, :+, 9, :+, 3, :*, 4, :*, 3]]) == [5, :+, [8, :*, [[3, :+, 9], :+, 3], :*, 4, :*, 3]]

parse_bracket_eval(expr) = evaluate(bracket(parse_expression(expr)))

@assert parse_bracket_eval("1 + (2 * 3) + (4 * (5 + 6))") == 51
@assert parse_bracket_eval("2 * 3 + (4 * 5)") == 46
@assert parse_bracket_eval("5 + (8 * 3 + 9 + 3 * 4 * 3)") == 1445
@assert parse_bracket_eval("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))") == 669060
@assert parse_bracket_eval("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2") == 23340

function part2(input)
    return reduce(+, map(parse_bracket_eval, split(input, '\n', keepempty=false)))
end

test = read("data/day-18.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 1408133923393
@assert result2 == 314455761823725
