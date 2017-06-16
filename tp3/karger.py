from random import randint
import copy

def karger(original):
    g = copy.deepcopy(original)
    while g.V() > 2:
        randE = g.get_rand_edge()
        g.move_node(randE.dst, randE.src)
    g.to_bi()
    return g
