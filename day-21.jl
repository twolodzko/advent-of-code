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

function map_candidates(food_list)
    allergen_names, allergen_counts, ingredient_names, ingredient_counts = count_relations(food_list)

    candidates = Dict()
    best = maximum(allergen_counts, dims=1)[:]
    for (i, row) in enumerate(eachrow(allergen_counts))
        for (j, v) in enumerate(best)
            if row[j] == best[j]
                allergen = allergen_names[j]
                ingredient = ingredient_names[i]

                if allergen in keys(candidates)
                    push!(candidates[allergen], ingredient)
                else
                    candidates[allergen] = Set([ingredient])
                end
            end
        end
    end
    return candidates
end

function reduce_candidates(candidates)
    matches = Dict()

    while !isempty(candidates)
        for (allergen, ingredients) in candidates
            if length(ingredients) == 1
                ingredient = first(ingredients)
                matches[allergen] = ingredient
                pop!(candidates, allergen)
                for (k, v) in candidates
                    if ingredient in v
                        pop!(v, ingredient)
                    end
                end
                break
            end
        end
    end
    return matches
end

function part2(input)
    food_list = parse_input(input)
    candidates = map_candidates(food_list)
    allergens = reduce_candidates(candidates)
    sorted = sort(collect(allergens), by=first)
    return join([food for (_, food) in sorted], ',')
end

@assert part2(example) == "mxmxvkd,sqjhc,fvjkl"

test = read("data/day-21.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 2412
@assert result2 == "mfp,mgvfmvp,nhdjth,hcdchl,dvkbjh,dcvrf,bcjz,mhnrqp"
