import numpy as np


# This function prints the moments, bins and masses for a random variable rv
def evaluate(rv, k=20):
    name, params, args = rv.dist.name, rv.kwds, rv.args   
    parts = []
    if args:
        parts += [repr(a) for a in args]
    if params:
        parts += [f"{k}={v!r}" for k, v in params.items()]
    print(f"{name.capitalize()}({', '.join(parts)})") 
    qs = np.linspace(0, 1, k+1)
    raw_bins = rv.ppf(qs)
    bins = np.unique(raw_bins)
    masses = np.diff(rv.cdf(bins))
    m,v,k = rv.stats(moments='mvk')
    print(f"mean: {m}, variance: {v}, kurtosis: {k+3}")
    print("bins:")
    print(np.array2string(bins, separator=', '))
    print("masses:")
    print(np.array2string(masses, separator=', '))
    print()

# These are the different test values

from scipy.stats import bernoulli
evaluate(bernoulli(0.2),k=2)
evaluate(bernoulli(0.9),k=2)

from scipy.stats import beta
evaluate(beta(0.6,3.0),k=20)
evaluate(beta(2.0,5.0),k=20)

from scipy.stats import binom
evaluate(binom(30,0.1),k=100)
evaluate(binom(300,0.2),k=100)
evaluate(binom(5000,0.9),k=100)
evaluate(binom(5000,0.001),k=2000)

from scipy.stats import cauchy
evaluate(cauchy(loc=-2.0, scale=7.0),k=20)

from scipy.stats import chi2
evaluate(chi2(1.0))
evaluate(chi2(6.0))

from scipy.stats import gamma
erlang = lambda  lambd, k: gamma(k,scale=1/lambd)
evaluate(erlang(2.0,5))
evaluate(erlang(0.6,3))

from scipy.stats import expon
evaluate(expon(scale=1/2.0))

from scipy.stats import f
evaluate(f(2.3,3.7))
evaluate(f(0.6,10.0))

from scipy.stats import gamma
evaluate(gamma(2.0,scale=4.0))
evaluate(gamma(0.6,scale=10.0))

from scipy.stats import geom
evaluate(geom(0.2))

from scipy.stats import gumbel_r
evaluate(gumbel_r(loc=1.5,scale=2.0))

from scipy.stats import invgauss 
ig = lambda nu,lambd: invgauss(nu/lambd, scale=lambd)
evaluate(ig(2.3,3.2))

from scipy.stats import laplace
evaluate(laplace(loc=1.5,scale=4.0))

from scipy.stats import logser
evaluate(logser(0.85),k=100)

exit()
from scipy.stats import logistic
evaluate(logistic(loc=2.5,scale=3.0))

from scipy.stats import lognorm
evaluate(lognorm(3.0,loc=0.0))

from scipy.stats import nbinom
evaluate(nbinom(29.5,0.2),k=80)

from scipy.stats import norm
evaluate(norm(loc=-3.0,scale=5.0))

from scipy.stats import pareto
evaluate(pareto(1.5,scale=2.0))

from scipy.stats import poisson
evaluate(poisson(15.0),k=50)
evaluate(poisson(55.0),k=50)

from scipy.stats import powerlaw
evaluate(powerlaw(2.3,scale=1.5))

from scipy.stats import rayleigh
evaluate(rayleigh(scale=2.3)) 

from scipy.stats import t
evaluate(t(6.0))

from scipy.stats import triang
tria = lambda a,b,c: triang((c-a)/(b-a), loc=a, scale=b-a)
evaluate(tria(-2.3,3.2,1.0))

from scipy.stats import uniform 
uni = lambda a,b: uniform(loc = a, scale=b-a)
evaluate(uni(-2.3,3.2))

from scipy.stats import weibull_min
evaluate(weibull_min(3.2,scale=2.3))
