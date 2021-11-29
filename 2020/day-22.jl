# https://adventofcode.com/2020/day/22

example = "
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
"

function read_input(input)
    deck1, deck2 = split(input, "\n\n", keepempty=false)
    deck1 = parse.(Int, split(deck1, '\n', keepempty=false)[2:end])
    deck2 = parse.(Int, split(deck2, '\n', keepempty=false)[2:end])
    return deck1, deck2
end

@assert read_input(example) == ([9, 2, 6, 3, 1], [5, 8, 4, 7, 10])

function combat(deck1, deck2)
    player = 1
    while !isempty(deck1) && !isempty(deck2)
        x = popfirst!(deck1)
        y = popfirst!(deck2)
        if x > y
            append!(deck1, [x, y])
        else
            append!(deck2, [y, x])
        end
    end
    return isempty(deck1) ? deck2 : deck1
end

score(deck) = sum([i * x for (i, x) in enumerate(reverse(deck))])

function part1(input)
    deck1, deck2 = read_input(input)
    return score(combat(deck1, deck2))
end

@assert part1(example) == 306

function recursive_combat(deck1, deck2; verbose=false)
    previous_rounds = Set()
    deck1 = copy(deck1)
    deck2 = copy(deck2)

    verbose && println("Game\n")

    while !isempty(deck1) && !isempty(deck2)
        verbose && println("Round")
        verbose && println("Player's 1 deck: $(deck1)")
        verbose && println("Player's 2 deck: $(deck2)")

        game_state = hash((deck1, deck2))
        if game_state in previous_rounds
            verbose && println("We saw those cards!")
            return deck1, []
        else
            push!(previous_rounds, game_state)
        end

        x = popfirst!(deck1)
        y = popfirst!(deck2)

        verbose && println("Player 1 plays: $x")
        verbose && println("Player 2 plays: $y")

        if (length(deck1) >= x) && (length(deck2) >= y)
            # sub-game
            verbose && println("Playing a sub-game to determine the winner...")
            winner = findfirst(.!isempty.(recursive_combat(deck1[1:x], deck2[1:y])))
            verbose && println("...anyway, back\n")
        else
            winner = x > y ? 1 : 2
            verbose && println("Player $winner wins\n")
        end

        if winner == 1
            append!(deck1, [x, y])
        else
            append!(deck2, [y, x])
        end
    end
    return deck1, deck2
end

function part2(input)
    deck1, deck2 = read_input(input)
    deck1, deck2 = recursive_combat(deck1, deck2)
    winner = isempty(deck1) ? deck2 : deck1
    return score(winner)
end

@assert part2(example) == 291

test = read("data/day-22.txt", String)
println("Part 1: $(result1 = part1(test))")
println("Part 2: $(result2 = part2(test))")

@assert result1 == 32472
@assert result2 == 36463
