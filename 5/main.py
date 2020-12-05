seats = []
for line in open("sample.txt").readlines():
    line = line.strip().replace("F", "0").replace("B", '1').replace("L", "0").replace("R", "1")
    seats.append((int(line[:-3], 2)*8)+int(line[-3:], 2))
myseat = [i for i in range(1028) if i not in seats and i + 1 in seats and i - 1 in seats]
print("Part 1: {}\nPart 2: {}".format(max(seats), myseat[0]))