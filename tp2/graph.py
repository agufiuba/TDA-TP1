from edge import Edge
from dijkstra import dijkstra
from bellmanford import bellmanford
from floydwarshall import floydwarshall
from random import randint

class Graph(object):
    def __init__(g, V):
        """ Construye el grafo con V nodos """
        g._a = [[] for _ in range(V)]
        g._i = [[] for _ in range(V)]
        g._e = [] * V

    def V(g):
        """ Cantidad de nodos """
        return len(g._a)

    def E(g):
        """ Cantidad de aristas """
        return sum(len(x) for x in g._i)

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

    def distance_to(g, path, d):
        suma = 0
        dst = d
        while True:
            src = path[dst]
            if src == dst or src == None:
                break
            suma += g.get_edge(src, dst).weight
            dst = src
        return suma

def create_complete_digraph(n, d = False):
    g = Graph(n)
    for i in range(n):
        for j in range(n):
            if i != j:
                w = randint(2, 25)
                g.add_edge(i, j, w)
                if d:
                    print "{0} -> {1}: {2}".format(i, j, w)
    return g

g = create_complete_digraph(7)
