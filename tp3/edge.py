class Edge(object):
    def __init__(self, src, dst, weight):
        self.src = src
        self.dst = dst
        self.weight = weight
        self.originalsrc = src
        self.originaldst = dst
