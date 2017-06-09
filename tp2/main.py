from graph import *
import sys

n = sys.argv[1]
alg = sys.argv[2]

if len(sys.argv) > 3:
    d = sys.argv[3]
else:
    d = False
g = create_complete_digraph(int(n), d)
if alg == "dj":
    print dijkstra(g, 0)
if alg == "bf":
    print bellmanford(g, 0)
if alg == "fw":
    print floydwarshall(g)
