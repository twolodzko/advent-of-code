
# find the two entries that sum to 2020 and then multiply those two numbers together

function part1(list_of_numbers, total=2020)
    for (i, x) in enumerate(list_of_numbers)
        y = total - x
        if y in list_of_numbers[i+1:end]
            return x * y
        end
    end
end

lst = [1721
       979
       366
       299
       675
       1456]

@assert part1(lst) == 514579

# Using the above example again, the three entries that sum to 2020 are
# 979, 366, and 675. Multiplying them together produces the answer, 241861950.

function part2(list_of_numbers)
    for (i, x) in enumerate(list_of_numbers)
        partial = 2020 - x
        tmp = filter(y -> y <= partial, list_of_numbers[i+1:end])
        y = part1(tmp, partial)
        if y !== nothing
            return x * y
        end
    end
end

@assert part2(lst) == 241861950