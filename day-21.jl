# https://adventofcode.com/2020/day/21

example = "
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
"

function parse_input(input)
    out = []
    for row in split(input, '\n', keepempty=false)
        ingredients, allergens = split(row, " (contains ", keepempty=false)
        allergens = replace(allergens, ")" => "")
        push!(out, [split(ingredients, " ", keepempty=false), split(allergens, ", ", keepempty=false)])
    end
    return out
end

function count_relations(food_list)
    allergen_names = Set()
    ingredient_names = Set()
    for (ingredients, allergens) in food_list
        union!(allergen_names, allergens)
        union!(ingredient_names, ingredients)
    end
    allergen_names = collect(allergen_names)
    ingredient_names = collect(ingredient_names)

    allergen_counts = zeros(Integer, (length(ingredient_names), length(allergen_names)))
    ingredient_counts = zeros(Integer, length(ingredient_names))

    for (ingredients, allergens) in food_list
        for ingredient in ingredients
            i = findfirst(ingredient .== ingredient_names)
            ingredient_counts[i] += 1
            for allergen in allergens
                j = findfirst(allergen .== allergen_names)
                allergen_counts[i, j] += 1
            end
        end
    end
    return allergen_names, allergen_counts, ingredient_names, ingredient_counts
end

function part1(input)
    food_list = parse_input(input)
    allergen_names, allergen_counts, ingredient_names, ingredient_counts = count_relations(food_list)

    count_impossible = 0
    for (i, ingredient) in enumerate(ingredient_names)
        possible_allergens = allergen_names[allergen_counts[i, :] .== maximum(allergen_counts, dims=1)[:]]
        if isempty(possible_allergens)
            count_impossible += ingredient_counts[i]
        end
    end
    return count_impossible
end

@assert part1(example) == 5

function part2(input)
    food_list = parse_input(input)
    allergen_names, allergen_counts, ingredient_names, ingredient_counts = count_relations(food_list)

    keep_indexes = ones(Bool, length(ingredient_names))
    for (i, ingredient) in enumerate(ingredient_names)
        possible_allergens = allergen_names[allergen_counts[i, :] .== maximum(allergen_counts, dims=1)[:]]
        if isempty(possible_allergens)
            keep_indexes[i] = false
        end
    end
    ingredient_counts = ingredient_counts[keep_indexes]
    ingredient_names = ingredient_names[keep_indexes]
    allergen_counts = allergen_counts[keep_indexes, :]

    food_allergens = Dict()
    while !all(allergen_counts .== 0)
        j = argmax(maximum(allergen_counts, dims=1)[:])
        for i in sortperm(allergen_counts[:, j], rev=true)
            if allergen_counts[i, j] != maximum(allergen_counts[i, :])
                # there is a better candidate for this row
                continue
            end

            food_allergens[ingredient_names[i]] = allergen_names[j]
            allergen_counts[i, :] .= 0
            allergen_counts[:, j] .= 0
            break
        end
    end

    sorted = sort(collect(food_allergens), by=pair -> pair[2])
    return join([food for (food, _) in sorted], ',')
end

@assert part2(example) == "mxmxvkd,sqjhc,fvjkl"

test = read("data/day-21.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 2412
@assert result2 != "mfp,mhnrqp,dcvrf,mgvfmvp,dvkbjh,nhdjth,bcjz,hcdchl"
@assert result2 != "mfp,dvkbjh,hcdchl,mgvfmvp,nhdjth,dcvrf"
@assert result2 != "mfp,mgvfmvp,bcjz,hcdchl,dvkbjh,dcvrf,nhdjth,mhnrqp"
