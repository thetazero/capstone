import matplotlib.pyplot as plt
import numpy as np


def ns_pn(p, q, debth):
    x = []
    ps = SizeSquared(p)
    for i in range(debth*2+1):
        n = i-debth
        x.append(1-SizeSquared(p)/SizeSquared(q+n*p))
    return x


def ns_an(nu, p, q, pn):
    lambdas = np.zeros(len(pn))
    constants = np.zeros(len(pn))
    for i in range(len(lambdas)):
        lambdas[i] = 1/pn[i]
        n = i-int(len(pn)/2)
        constants[i] = nu*SizeSquared(q+n*p)/pn[i]
    return [constants, lambdas]


def computeAtLambda(la, constants, lamdas):
    cur = 0
    for i in reversed(range(len(constants))):
        u = constants[i]+lamdas[i]*la
        cur += u
        cur = 1/cur
    return cur


def SizeSquared(p):
    sum = 0
    for v in p:
        sum += v*v
    return sum


def plot(nu, p, q, debth, start, end, step):
    start, end, step = float(start), float(end), float(step)
    pn = ns_pn(p, q, debth)
    [constants, lambdas] = ns_an(nu, p, q, pn)
    print(constants, lambdas)
    pConst = constants[debth+1:]
    nConst = list(reversed(constants[:debth]))
    pLam = lambdas[debth+1:]
    nLam = list(reversed(lambdas[:debth]))
    p0Const = constants[debth]
    p0Lam = lambdas[debth]
    print(p0Const, p0Lam)
    print(pConst, pLam)
    print(nConst, nLam)
    x = list(map(lambda x: float(x)*step,
                 range(int(start/step), int(end/step))))
    y = list(map(lambda x: computeAtLambda(x, pConst, pLam) +
                 computeAtLambda(x, nConst, nLam)+x*p0Lam+p0Const, x))
    # print(x, y)
    plt.plot(x, y)
    plt.title(label='f(λ,ν)+g(λ,ν)+aₒ(λ,ν)')
    plt.axhline(y=0, color="k")
    plt.axvline(x=0, color="k")
    plt.show()


plot(0.05, np.array([3, 1]), np.array([-1, 2]), 10, 0.0, 1.0, 0.001)
