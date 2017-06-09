import heapq

def dijkstra(g, n):
    distances = {}
    padres = {}
    visto = {}
    cola = []
    for v in g:
        distances[v] = float("inf")
        padres[v] = None
        visto[v] = False
    distances[n] = 0
    heapq.heappush(cola, (distances[n], n))
    while cola:
        distanceN, n = heapq.heappop(cola)
        visto[n] = True
        for e in g.adj_e(n):
            v = e.dst
            if not visto[v] and distances[v] > distanceN + e.weight:
                distances[v] = distanceN + e.weight
                padres[v] = n
                heapq.heappush(cola, (distances[v], v))
    return padres
