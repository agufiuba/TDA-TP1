import random
from edge import Edge
from karger import karger

class Graph(object):
    def __init__(g, V):
        """ Construye el grafo con V nodos """
        g._a = [[] for _ in range(V)]
        g._i = [[] for _ in range(V)]
        g._e = []

    def V(g):
        """ Cantidad de nodos """
        return len(g._a)

    def E(g):
        """ Cantidad de aristas """
        return len(g._e)

    def adj_e(g, v):
        """ Itera sobre los aristas de v """
        return iter(g._i[v])

    def adj(g, v):
        """ Itera sobre los nodos adyacentes a v """
        return iter(g._a[v])

    def add_edge(g, u, v, weight=0):
        """ Anade y devuelve una arista """
        e = Edge(u, v, weight)
        g._i[u].append(e)
        g._a[u].append(v)
        g._e.append(e)
        return e

    def __iter__(g):
        """ Itera los nodos """
        return iter(range(g.V()))

    def iter_edges(g):
        """ Itera sobre todas las aristas """
        return iter(g._e)

    def has_node(g, v):
        return v < len(g._a)

    def get_edge(g, s, d):
        return filter(lambda x: x.dst == d, g.adj_e(s))[0]

    def get_edges(g, *s):
        print s

    def move_node(g, s, d):
        g._e = filter(lambda e: not(e.src == s and e.dst == d) and not(e.src == d and e.dst == s), g._e)
        for e in filter(lambda e: e.src == s or e.dst == s, g._e):
            if e.src == s:
                e.src = d
            else:
                e.dst = d
        for e in g._e:
            if e.src > s:
                e.src-=1
            if e.dst > s:
                e.dst-=1
        for adjs in g._a:
            for a in adjs:
                if a > s:
                    a-=1
        g._a.pop()
        g._i.pop()

    def get_rand_edge(g):
        return g._e[random.randint(0, len(g._e)-1)]

    def to_bi(g):
        for i in range(g.E()/2):
            g._e = filter(lambda e: not e.src == g._e[i].dst and not e.dst == g._e[i].src, g._e)

def create_graph_2n(n, d = False):
    if n > 4:
        g = Graph(n)
        for i in range(n):
            if i != n - 1:
                if d:
                    print "{0} <--> {1}".format(i, i+1)
                g.add_edge(i, i+1)
                g.add_edge(i+1, i)
                if i != n-2:
                    if d:
                        print "{0} <--> {1}".format(i, i+2)
                    g.add_edge(i, i+2)
                    g.add_edge(i+2, i)
                else:
                    if d:
                        print "{0} <--> {1}".format(i, 0)
                    g.add_edge(i, 0)
                    g.add_edge(0, i)
            else:
                if d:
                    print "{0} <--> {1}".format(i, 0)
                    print "{0} <--> {1}".format(i, 1)
                g.add_edge(i, 0)
                g.add_edge(i, 1)
                g.add_edge(0, i)
                g.add_edge(1, i)
        return g
    print "El tamano del grafo tiene que ser mayor a 4"
    return None
