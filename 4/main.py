import re

validation = {
    "byr": lambda x: 1920 <= int(x) <= 2002,
    "iyr": lambda x: 2010 <= int(x) <= 2020,
    "eyr": lambda x: 2020 <= int(x) <= 2030,
    "hgt": lambda x: (x.endswith("in") and 59 <= int(x.strip("in")) <= 76) or (x.endswith("cm") and 150 <= int(x.strip("cm")) <= 193),
    "hcl": lambda x: bool(re.match("^#[0-9a-f]{6}$", x)),
    "ecl": lambda x: bool(re.match("^(amb|blu|brn|gry|grn|hzl|oth)$", x)),
    "pid": lambda x: bool(re.match("^[0-9]{9}$", x))
}

count = 0
count2 = 0
with open("data.txt") as fil:
    for r in fil.read().split("\n\n"):
        rcd = {}
        for k, v in [x.split(":") for x in r.strip().replace("\n", " ").split(" ")]:
            rcd[k] = v
        if all([i in rcd for i in validation.keys()]):
            count += 1
            if all([fn(rcd[i]) for i, fn in validation.items()]):
                count2 += 1
print("Part 1:", count)
print("Part 2:", count2)
