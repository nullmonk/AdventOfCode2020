import sys, math
from collections import defaultdict

def countPerms(nxt):
    prev = nxt.pop(0)
    i = 0
    print("r")
    ans = 1
    while True:
        if i >= len(nxt):
            return 1

        if nxt[i] <= prev + 3:
            ans += 1
        i += 1
    return ans * countPerms(nxt)


with open(sys.argv[1]) as fil:
    adapters = sorted([0]+[int(i.strip()) for i in fil.readlines()])

adapters.append(max(adapters) + 3)
voltages = defaultdict(int)


print(adapters)
for i in range(len(adapters) - 1):
    voltages[adapters[i+1]-adapters[i]] += 1


print(voltages)
print("Part 1:", voltages[3] * voltages[1])
print(countPerms(adapters.copy()))