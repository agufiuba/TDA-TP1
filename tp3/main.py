from graph import *
from stock import *
from approx_subset_sum import *
import sys

if sys.argv[1] == "stock":
    n = int(sys.argv[2])
    test = []
    for _ in range(n):
        test.append(random.randint(-25, 25))
    print test
    c, v = profit(test)
    print "Conviente comprar el dia " + str(c+1) + " y vender el dia " + str(v+1)
if sys.argv[1] == "karger":
    n = int(sys.argv[2])
    if len(sys.argv) > 3:
        d = sys.argv[3]
    else:
        d = False
    k = karger(create_graph_2n(n, d))
    print "El corte minimo aproximado es: " + str(k.E())
    if d:
        print "Y las aristas son: "
        for e in k.iter_edges():
            print str(e.originalsrc) + "---" + str(e.originaldst)
if sys.argv[1] == "subsetsum":
    size = int(sys.argv[2])
    t = int(sys.argv[3])
    approx = float(sys.argv[4])
    test = []
    for _ in range(size):
        test.append(random.randint(0, 25))
    print test
    print "La aproximacion es: " + str(approx_subset_sum(test, t, approx))
