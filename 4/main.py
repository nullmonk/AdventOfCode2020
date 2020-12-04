with open("data.txt") as fil:
    count = 0
    for r in fil.read().split("\n\n"):
        rcd = {}
        for k, v in [x.split(":") for x in r.strip().replace("\n", " ").split(" ")]:
            rcd[k] = v
        if all([i in rcd for i in ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]]):
            count += 1
    print(count)