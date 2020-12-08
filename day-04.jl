# https://adventofcode.com/2020/day/4

inputs = "
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
"

struct Field
    name::String
    value::String

    Field(string) = new(split(strip(string), ":")...)
end

function extract_rows(input)
    entries = [split(replace(row, r"\n" => " "), r"\s+") for row in split(input, "\n\n")]
    return [[Field(field) for field in entry if field != ""] for entry in entries]
end

function part1(input)
    expected = sort(["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"])
    result = 0
    for row in extract_rows(input)
        fields = sort(map(x -> x.name, row))

        if filter(x -> x != "cid", fields) == expected
            result += 1
        end
    end
    return result
end

@assert part1(inputs) == 2


function isvalid(field::Field)
    function isvalidhgt(x:: String)
        m = match(r"(\d+)(cm|in)", x)
        if m === nothing
            return false
        end

        value, units = m.captures
        value = parse(Int, value)

        if units == "cm"
            return 150 <= value <= 193
        else
            return 59 <= value <= 76
        end
    end

    return Dict(
        "byr" => x -> 1920 <= parse(Int, x) <= 2002,
        "iyr" => x -> 2010 <= parse(Int, x) <= 2020,
        "eyr" => x -> 2020 <= parse(Int, x) <= 2030,
        "hgt" => isvalidhgt,
        "hcl" => x -> match(r"^#[a-f0-9]{6}$", x) !== nothing,
        "ecl" => x -> x in Set(["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]),
        "pid" => x -> match(r"^[0-9]{9}$", x) !== nothing,
        "cid" => x -> true,
    )[field.name](field.value)
end


function part2(input)
    expected = sort(["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"])
    result = 0
    for row in extract_rows(input)
        fields = sort(map(x -> x.name, row))

        if filter(x -> x != "cid", fields) != expected
            continue
        end

        ok = true
        for field in row
            ok = ok && isvalid(field)
        end

        result += ok
    end
    return result
end


inputs = "
eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
"

@assert part2(inputs) == 4

test = read("data/day-04.txt", String)
println("Part 1: $(part1(test))")
println("Part 2: $(part2(test))")

@assert part1(test) == 208
@assert part2(test) == 167
