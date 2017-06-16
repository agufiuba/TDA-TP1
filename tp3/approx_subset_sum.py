def trim(es, approx):
    es = sorted(es)
    last = es[0]
    new = [last]
    for e in es[1:]:
        if e > last*(1+approx):
            new.append(e)
            last = e
    return new

def approx_subset_sum(es, t, approx = 0.1):
    L = [[] for _ in range(len(es)+1)]
    L[0].append(0)
    for i in range(len(es)):
        auxL = []
        for l in L[i]:
            auxL.append(l + es[i])
        L[i+1] = L[i] + auxL
        L[i+1] = trim(L[i+1], approx/(2*len(es)))
        for j,_ in enumerate(L[i+1]):
            L[i+1] = filter(lambda x: x <= t, L[i+1])
    return max(L[len(es)])
