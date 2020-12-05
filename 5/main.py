import sys

def getRowCol(string):
    data = list(string[:-3])
    x = 128
    mn, mx = 1, 128
    row, col = 0, 0
    while data:
        mid = (mx - mn)//2
        let = data.pop(0)
        if let == 'B':
            mn = mx-mid
        else:
            mx = mn + mid
        if mn == mx:
            row = mn
            break
    data = list(string[-3:])
    mn = 1
    mx = 8
    while data:
        mid = (mx - mn)//2
        let = data.pop(0)
        if let == 'R':
            mn = mx-mid
        else:
            mx = mn + mid
        if mn == mx:
            col = mn
            break
    return row-1, col-1, (row-1)*8+(col-1)


seats = {}
with open(sys.argv[1]) as fil:
    ids = []
    for line in fil.readlines():
        line = line.strip()
        if not line:
            continue
        seatdata = getRowCol(line)
        seats[seatdata[2]] = seatdata[:2]
        #ids += [seatdata[2]]
    print(max(seats.keys()))

# check the back row
myseat = -1
prev = False
for i in range(0, 1028):
    if i not in seats:
        if i+1 in seats and i-1 in seats:
            myseat = i
print(myseat)