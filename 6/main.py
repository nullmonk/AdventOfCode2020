from collections import Counter
import sys
part1, part2 = 0, 0
for group in open(sys.argv[1]).read().split("\n\n"):
    part1 += len(set("".join(group.split())))
    part2 += len([i for i in Counter(group.strip()).values() if i == group.strip().count("\n") + 1])
print(part1, part2)
