import sys

class Numbers(object):
    def __init__(self, mx=25):
        self.max = mx
        self._ar = []

    def append(self, number):
        if isinstance(number, str):
            number = int(number.strip())
        self._ar.append(number)

    def isValid(self, number):
        rng = self._ar[-self.max:]
        for i in rng:
            j = number - i
            if i != j and j in rng:
                return True
        return False
    
    def check(self, number):
        if isinstance(number, str):
            number = int(number.strip())
        if len(self._ar) < self.max:
            self.append(number)
            return True
        if self.isValid(number):
            self.append(number)
            return True
        return False

    def checkContiguous(self, number):
        if isinstance(number, str):
            number = int(number.strip())
        for i in range(len(self._ar)):
            for j in range(len(self._ar)):
                nm = self._ar[i:j].copy()
                snm = sum(nm)
                if snm == number:
                    return nm
                if snm > number:
                    break
        return "404"
                
def main():
    nums = Numbers(25)
    invalid = 0
    with open(sys.argv[1]) as fil:
        for line in fil.readlines():
            line = line.strip()
            if not line:
                continue
            if not nums.check(line):
                invalid = int(line.strip())
                break
    
    print("Part 1:", invalid)
    n = nums.checkContiguous(invalid)
    print(n)
    print("Part 2:", min(n) + max(n))
    
main()