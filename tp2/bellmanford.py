def bellmanford(g, n):
    distances = {}
    padres = {}
    for v in g:
        distances[v] = float("inf")
        padres[v] = None
    distances[n] = 0
    for _ in range(1, g.V()):
        for e in g.iter_edges():
            if distances[e.dst] > distances[e.src] + e.weight:
                distances[e.dst] = distances[e.src] + e.weight
                padres[e.dst] = e.src
    return padres
