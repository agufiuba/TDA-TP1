def floydwarshall(g):
    fw = [[float("inf") for _ in range(g.V())] for _ in range(g.V())]
    for i in range(g.V()):
        fw[i][i] = 0
    for e in g.iter_edges():
        fw[e.src][e.dst] = e.weight
    for i in range(g.V()):
        for j in range(g.V()):
            for k in range(g.V()):
                if fw[j][k] > fw[j][i] + fw[i][k]:
                    fw[j][k] = fw[j][i] + fw[i][k]
    return fw
