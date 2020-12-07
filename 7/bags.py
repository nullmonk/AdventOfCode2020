import sys
from collections import defaultdict


class Bags(dict):
    def __init__(self):
        self._bags = {}
    
    def __getitem__(self, bag):
        return self._bags[str(bag)]
    
    def newBag(self, color, line=""):
        b = Bag(color, line, parent=self)
        self._bags[str(color)] = b
        return b

    def items(self):
        return self._bags.items()


class Bag(object):
    def __init__(self, color, line, parent=None):
        self._line = line
        self.color = color
        self.bags = defaultdict(int)
        self.parent = parent

    def addSubBags(self, *bags):
        for bag in bags:
            bag = bag.strip(" .\n").split()
            color = " ".join(bag[1:-1])
            count = int(bag[0])
            self.bags[color] = count
        
    def __str__(self):
        return self.color

    def __hash__(self):
        return hash(self.color)

    def getCount(self):
        fcount = 0
        for b, count in self.bags.items():
            i = (count * (self.parent[b].getCount()+1))
            fcount += i
        return fcount

    def getBagString(self):
        string = []
        for bag_name in self.bags:
            string += ["{}: {}".format(bag_name, self.parent[bag_name].getBagString())]
        return ", ".join(string)


def main():
    bags = Bags()
    with open(sys.argv[1]) as fil:
        for line in fil.readlines():
            color, rest = line.split(" bags contain ", maxsplit=1)
            b = bags.newBag(color, line)
            if "no other bags" in rest:
                continue
            b.addSubBags(*rest.split(", "))

    count = 0
    for b, v in bags.items():
        if "shiny gold" in v.getBagString():
            count += 1


    print(count)
    print(bags["shiny gold"].getCount())

main()