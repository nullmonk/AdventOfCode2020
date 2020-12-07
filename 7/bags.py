import sys
from collections import defaultdict


#class Bags(dict):
#    def __init__(self):

bag_colors = {}

class Bag(object):
    def __init__(self, color, line, bagmap={}):
        self._line = line
        self.color = color
        self.bags = defaultdict(int)
        self.bagmap = bagmap

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
            i = (count * self.bagmap[b].getCount())
            fcount += i
        return fcount + 1
    
    def getCountTest(self):
        print(self.color, "bags contain ", end="")
        sm = 0
        for k, v in self.bags.items():
            print(v, k, "bags, ", end="")
            sm += v
        print("total:", sm)
    
    def getBagString(self):
        string = []
        for bag_name in self.bags:
            string += ["{}: {}".format(bag_name, self.bagmap[bag_name].getBagString())]
        return ", ".join(string)


with open(sys.argv[1]) as fil:
    for line in fil.readlines():
        color, rest = line.split(" bags contain ", maxsplit=1)
        b = Bag(color, line.strip(), bag_colors)
        bag_colors[b.color] = b
        if "no other bags" in rest:
            continue
        b.addSubBags(*rest.split(", "))


count = 0
for b, v in bag_colors.items():
    if "shiny gold" in v.getBagString():
        count += 1


print(count)
print(bag_colors["shiny gold"].getCount() - 1)