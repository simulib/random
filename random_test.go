// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

package random

import (
	"fmt"
	"testing"
)
import "math"
import "github.com/dfiems/simulib/random/rng/pcg64dxsm"

func TestBernoulli(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Bernoulli with p=0.2
	s, _ := r.GetBernoulliSampler(0.2)
	momentTestInt(t, s, 0.2, 0.16, 3.25, 100000)

	// Bernoulli with p=0.9
	s, _ = r.GetBernoulliSampler(0.9)
	momentTestInt(t, s, 0.9, 0.09, 8.1111, 100000)
}

func TestBeta(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Beta distribution: a = 0.6, b = 3.0
	s, _ := r.GetBetaSampler(0.6, 3.0)
	histogramTest(t,
		s,
		[]float64{0., 0.00200725, 0.00640778, 0.01269435, 0.02071111, 0.03040901,
			0.04180077, 0.05494688, 0.06995306, 0.08697426, 0.10622473, 0.1279955,
			0.1526833, 0.18083841, 0.21324767, 0.25108782, 0.29623711, 0.35199888,
			0.42517108, 0.53472078, 1.},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		1000000)
	momentTest(t, s, 0.16666666666666666, 0.030193236714975844, 4.480519480519481, 10000)

	// Beta distribution: a = 2.0, b = 5.0
	s, _ = r.GetBetaSampler(2.0, 5.0)
	histogramTest(t,
		s,
		[]float64{0., 0.06284989, 0.09259526, 0.1173795, 0.13988069, 0.16116292,
			0.18180347, 0.20218104, 0.22258353, 0.24325963, 0.26444998, 0.28641175,
			0.30944443, 0.33392413, 0.36035769, 0.38947949, 0.42244752, 0.4613036,
			0.51031631, 0.58180341, 1.},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		1000000)
	momentTest(t, s, 0.2857142857142857, 0.025510204081632654, 2.88, 10000)
}

func TestBinomial(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Binomial distribution: p = 0.1, N = 30, this is a test for Inversion
	s, _ := r.GetBinomialSampler(30, 0.1)
	histogramTestInt(t,
		s,
		[]float64{-1., 0., 1., 2., 3., 4., 5., 6., 7., 30.},
		[]float64{0.04239116, 0.14130386, 0.22765622, 0.23608793, 0.17706595, 0.10230477,
			0.04736332, 0.01804317, 0.00778362},
		1000000)
	momentTestInt(t, s, 3.0, 2.7, 3.1703703703703705, 100000)

	// Binomial distribution: p = 0.2, N = 300, this is a test for BTPE
	s, _ = r.GetBinomialSampler(300, 0.2)
	histogramTestInt(t,
		s,
		[]float64{-1., 44., 46., 47., 48., 49., 50., 51., 52., 53., 54., 55.,
			56., 57., 58., 59., 60., 61., 62., 63., 64., 65., 66., 67.,
			68., 69., 70., 71., 72., 73., 75., 77., 300.},
		[]float64{0.01060503, 0.01243882, 0.00976182, 0.01286323, 0.01653844, 0.02075574,
			0.02543596, 0.03044978, 0.0356205, 0.0407327, 0.04554656, 0.04981655,
			0.05331245, 0.0558402, 0.05725986, 0.05749845, 0.05655585, 0.05450342,
			0.05147545, 0.04765501, 0.04325608, 0.03850447, 0.03361957, 0.02879912,
			0.02420796, 0.01997156, 0.01617415, 0.0128607, 0.01004192, 0.01350253,
			0.00741665, 0.00697946},
		1000000)
	momentTestInt(t, s, 60.0, 48.0, 3.000833, 100000)

	// Binomial distribution: p = 0.9, N = 5000, this is a test for BTPE
	s, _ = r.GetBinomialSampler(5000, 0.9)
	histogramTestInt(t,
		s,
		[]float64{-1.000e+00, 4.450e+03, 4.456e+03, 4.460e+03, 4.463e+03, 4.465e+03,
			4.467e+03, 4.469e+03, 4.470e+03, 4.471e+03, 4.473e+03, 4.474e+03,
			4.475e+03, 4.476e+03, 4.477e+03, 4.478e+03, 4.479e+03, 4.480e+03,
			4.481e+03, 4.482e+03, 4.483e+03, 4.484e+03, 4.485e+03, 4.486e+03,
			4.487e+03, 4.488e+03, 4.489e+03, 4.490e+03, 4.491e+03, 4.492e+03,
			4.493e+03, 4.494e+03, 4.495e+03, 4.496e+03, 4.497e+03, 4.498e+03,
			4.499e+03, 4.500e+03, 4.501e+03, 4.502e+03, 4.503e+03, 4.504e+03,
			4.505e+03, 4.506e+03, 4.507e+03, 4.508e+03, 4.509e+03, 4.510e+03,
			4.511e+03, 4.512e+03, 4.513e+03, 4.514e+03, 4.515e+03, 4.516e+03,
			4.517e+03, 4.518e+03, 4.519e+03, 4.520e+03, 4.521e+03, 4.522e+03,
			4.523e+03, 4.524e+03, 4.525e+03, 4.526e+03, 4.527e+03, 4.528e+03,
			4.530e+03, 4.531e+03, 4.533e+03, 4.535e+03, 4.537e+03, 4.540e+03,
			4.543e+03, 4.549e+03, 5.000e+03},
		[]float64{0.01053354, 0.01057849, 0.01125033, 0.01138381, 0.00925807, 0.0107607,
			0.01240282, 0.00685857, 0.00731724, 0.01606613, 0.00877362, 0.00928141,
			0.00979773, 0.01032077, 0.01084857, 0.01137901, 0.01190986, 0.01243877,
			0.0129633, 0.0134809, 0.01398899, 0.01448492, 0.01496603, 0.01542965,
			0.01587315, 0.01629394, 0.01668949, 0.0170574, 0.01739536, 0.01770122,
			0.017973, 0.01820891, 0.01840736, 0.018567, 0.0186867, 0.01876562,
			0.01880315, 0.01879897, 0.01875304, 0.01866558, 0.01853711, 0.01836841,
			0.01816051, 0.01791472, 0.01763256, 0.01731581, 0.01696642, 0.01658655,
			0.0161785, 0.01574473, 0.01528781, 0.01481038, 0.01431517, 0.01380493,
			0.01328244, 0.01275043, 0.01221164, 0.01166872, 0.01112425, 0.01058071,
			0.01004044, 0.00950569, 0.00897853, 0.00846088, 0.0079545, 0.01444265,
			0.00651788, 0.01171133, 0.01006468, 0.00856811, 0.01038195, 0.0078707,
			0.01007908, 0.00906861},
		1000000)
	momentTestInt(t, s, 4500.0, 450.0, 3.00102, 10000000)

	// Binomial distribution: p = 0.001, N = 5000, this is a test for waiting time
	s, _ = r.GetBinomialSampler(5000, 0.001)
	histogramTestInt(t,
		s,
		[]float64{-1.0e+00, 0.0e+00, 1.0e+00, 2.0e+00, 3.0e+00, 4.0e+00, 5.0e+00,
			6.0e+00, 7.0e+00, 8.0e+00, 9.0e+00, 1.0e+01, 1.1e+01, 1.2e+01,
			1.3e+01, 1.4e+01, 5.0e+03},
		[]float64{0.00672111, 0.0336392, 0.08416534, 0.14035982, 0.17552003, 0.17555517,
			0.14629597, 0.1044762, 0.06527148, 0.03624016, 0.01810557, 0.00822157,
			0.00342154, 0.00131413, 0.00046858, 0.00022414},
		1000000)
	momentTestInt(t, s, 5.0, 4.995, 3.1990002002002003, 10000000)
}

func TestCauchy(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetCauchySampler(-2.0, 7.0)
	histogramTest(t,
		s,
		[]float64{math.Inf(-1), -46.1962606, -23.54378476, -15.73827354, -11.63467344,
			-9., -7.0857977, -5.56667815, -4.27443787, -3.10869108,
			-2., -0.89130892, 0.27443787, 1.56667815, 3.0857977,
			5., 7.63467344, 11.73827354, 19.54378476, 42.1962606, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		1000000)
}

func TestChiSq(t *testing.T) {
	rng := pcg64dxsm.New(1234, 4756)
	r := New(rng)

	// ChiSq distribution: k = 1.0
	s, _ := r.GetChiSqSampler(1.0)
	histogramTest(t,
		s,
		[]float64{0., 0.00393214, 0.01579077, 0.03576578, 0.06418475, 0.10153104,
			0.14847186, 0.20590013, 0.2749959, 0.35731717, 0.45493642, 0.57065186,
			0.7083263, 0.87345714, 1.07419417, 1.3233037, 1.64237442, 2.07225086,
			2.70554345, 3.84145882, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 1.0, 2.0, 15.0, 100000)
	// ChiSq distribution: k = 6.0
	s, _ = r.GetChiSqSampler(6.0)
	histogramTest(t,
		s,
		[]float64{0., 1.63538289, 2.20413066, 2.66127318, 3.07008841, 3.45459884,
			3.82755159, 4.19726953, 4.57015381, 4.95187661, 5.34812063, 5.76519934,
			6.21075719, 6.69476088, 7.23113533, 7.84080412, 8.55805972, 9.44610313,
			10.64464068, 12.59158724, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 6.0, 12.0, 5.0, 100000)
}

func TestErlang(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Erlang distribution: lambda = 2.0, k = 5
	s, _ := r.GetErlangSampler(2.0, 5)
	histogramTest(t,
		s,
		[]float64{0., 0.98507478, 1.21629551, 1.39251486, 1.54476981, 1.68430019,
			1.81680454, 1.94581074, 2.07386794, 2.20308795, 2.33545444, 2.47305393,
			2.61830906, 2.77428557, 2.94518066, 3.13721535, 3.36048939, 3.633484,
			3.99679479, 4.57675951, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 2.5, 1.25, 4.2, 100000)
	// Erlang distribution: lambda = 0.6, k = 3
	s, _ = r.GetErlangSampler(0.6, 3)
	histogramTest(t,
		s,
		[]float64{0., 1.36281908, 1.83677555, 2.21772765, 2.558407,
			2.87883236, 3.18962632, 3.49772461, 3.80846151, 4.12656384,
			4.45676719, 4.80433278, 5.175631, 5.5789674, 6.02594611,
			6.53400343, 7.13171643, 7.87175261, 8.8705339, 10.49298937, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 5.0, 8.333333333333334, 5.0, 100000)
}

func TestExponential(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Erlang distribution: lambda = 2.0, k = 5
	s, _ := r.GetExponentialSampler(2.0)
	histogramTest(t,
		s,
		[]float64{0., 0.02564665, 0.05268026, 0.08125946, 0.11157178, 0.14384104,
			0.17833747, 0.21539146, 0.25541281, 0.2989185, 0.34657359, 0.39925385,
			0.45814537, 0.52491106, 0.6019864, 0.69314718, 0.80471896, 0.94855999,
			1.15129255, 1.49786614, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 0.5, 0.25, 9.0, 100000)
}

func TestF(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// F distribution: d1 = 2.3, d2 = 3.7
	s, _ := r.GetFSampler(2.3, 3.7)
	histogramTest(t,
		s,
		[]float64{0., 0.06991629, 0.13461185, 0.20175348, 0.27322467, 0.35044009,
			0.43483266, 0.52805414, 0.63213903, 0.74970173, 0.88421889, 1.04047001,
			1.22526876, 1.44875226, 1.72681908, 2.08616395, 2.57596451, 3.29989848,
			4.52979609, 7.36645617, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	// no moment test as the variance is infinite
	// F distribution: d1 = 0.6, d2 = 10.0
	s, _ = r.GetFSampler(0.6, 10.0)
	histogramTest(t,
		s,
		[]float64{0.00000000e+00, 1.14891129e-04, 1.15832560e-03, 4.47872215e-03,
			1.17053236e-02, 2.47055406e-02, 4.55974334e-02, 7.68039064e-02,
			1.21153667e-01, 1.82044830e-01, 2.63701839e-01, 3.71582564e-01,
			5.13043678e-01, 6.98480705e-01, 9.43407796e-01, 1.27257535e+00,
			1.72906465e+00, 2.39773038e+00, 3.48199367e+00, 5.69815086e+00, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 1.25, 7.465277777777779, 156.62790697674416, 100000)
}

func TestGamma(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Gamma distribution: alpha = 2.0, theta = 4.0
	s, _ := r.GetGammaSampler(2.0, 4.0)
	histogramTest(t,
		s,
		[]float64{0., 1.42144604, 2.12724643, 2.73295445, 3.29755324,
			3.84511505, 4.38939684, 4.94017477, 5.50568537, 6.09389285,
			6.71338796, 7.37426767, 8.08925298, 8.8753784, 9.75686593,
			10.77053812, 11.97723339, 13.48976617, 15.55888068, 18.97545807, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 8.0, 32.0, 6.0, 100000)
	// Gamma distribution: alpha = 0.6, theta = 10.0
	s, _ = r.GetGammaSampler(0.6, 10.0)
	histogramTest(t,
		s,
		[]float64{0., 0.05644834, 0.18060438, 0.35894098, 0.58803353,
			0.8677172, 1.19988825, 1.58821632, 2.03822579, 2.55764048,
			3.15702017, 3.85081382, 4.65909246, 5.61049129, 6.74747967,
			8.13653578, 9.88991799, 12.21960081, 15.60503416, 21.59012054, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 6.0, 60.0, 13.0, 100000)
}

func TestGeometric(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)

	// Geometric distribution: alpha = 2.0, theta = 4.0
	s, _ := r.GetGeometricSampler(0.2)
	histogramTestInt(t,
		s,
		[]float64{0., 1., 2., 3., 4., 5., 6., 7., 8., 9., 11., 14., math.Inf(1)},
		[]float64{0.2, 0.16, 0.128, 0.1024, 0.08192, 0.065536,
			0.0524288, 0.04194304, 0.03355443, 0.04831838, 0.04191888, 0.04398047},
		10000000)
	momentTestInt(t, s, 5.0, 20.0, 9.05, 100000)
}

func TestGumbel(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetGumbelSampler(1.5, 2.0)
	histogramTest(t,
		s,
		[]float64{math.Inf(-1), -0.6943774, -0.16806489, 0.21932612, 0.54823001,
			0.84673148, 1.12874648, 1.40275851, 1.67484314, 1.95002135,
			2.23302584, 2.52887427, 2.84345398, 3.18430198, 3.56186087,
			3.99179865, 4.49987997, 5.13392159, 6.00073465, 7.4403905, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 2.6544313298030655, 6.579736267392906, 5.4, 100000)
}

func TestInverseGaussian(t *testing.T) {
	rng := pcg64dxsm.New(123, 5678)
	r := New(rng)
	s, _ := r.GetInverseGaussianSampler(2.3, 3.2)
	momentTest(t, s, 2.3, 3.8021874999999987, 13.781249999999998, 100000)
	histogramTest(t,
		s,
		[]float64{0., 0.53157286, 0.67185587, 0.79336305, 0.90955907, 1.02583937,
			1.14543927, 1.27088354, 1.40455204, 1.54900421, 1.70725408, 1.88309156,
			2.08154116, 2.30960088, 2.5775448, 2.90143489, 3.30852406, 3.85075833,
			4.64725865, 6.08939875, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		1000000)

}

func TestLaplace(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetLaplaceSampler(1.5, 4.0)
	histogramTest(t,
		s,
		[]float64{math.Inf(-1), -7.71034037, -4.93775165, -3.31589122, -2.16516293,
			-1.27258872, -0.5433025, 0.07330022, 0.60742579, 1.07855794,
			1.5, 1.92144206, 2.39257421, 2.92669978, 3.5433025,
			4.27258872, 5.16516293, 6.31589122, 7.93775165, 10.71034037, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 1.5, 32, 6.0, 100000)
}

func TestLogarithmic(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetLogarithmicSampler(0.85)
	histogramTestInt(t,
		s,
		[]float64{0., 1., 2., 3., 4., 5., 6., 7., 8., 9., 10., 12., 13., 17.,
			math.Inf(1)},
		[]float64{0.44804757, 0.19042022, 0.10790479, 0.0687893, 0.04677673, 0.03313351,
			0.02414013, 0.01795422, 0.01356541, 0.01037754, 0.01426715, 0.00490239,
			0.01134222, 0.0083788},
		1000000)
	momentTestInt(t, s, 2.9870, 10.9912, 20.0852, 100000)
}

func TestLogistic(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetLogisticSampler(2.5, 3.0)
	histogramTest(t,
		s,
		[]float64{math.Inf(-1), -6.33331694, -4.09167373, -2.70380317, -1.65888308,
			-0.79583687, -0.04189358, 0.64288237, 1.28360468, 1.89798791,
			2.5, 3.10201209, 3.71639532, 4.35711763, 5.04189358,
			5.79583687, 6.65888308, 7.70380317, 9.09167373, 11.33331694, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 2.5, 29.608813203268074, 4.2, 100000)
}

func TestLognormal(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetLognormalSampler(0.0, 3.0)
	histogramTest(t,
		s,
		[]float64{0.00000000e+00, 7.19361911e-03, 2.13937876e-02, 4.46321794e-02,
			8.00692254e-02, 1.32196047e-01, 2.07380183e-01, 3.14754786e-01,
			4.67647108e-01, 6.85927021e-01, 1.00000000e+00, 1.45788104e+00,
			2.13836455e+00, 3.17707639e+00, 4.82206152e+00, 7.56452269e+00,
			1.24891929e+01, 2.24053589e+01, 4.67425412e+01, 1.39012086e+02, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 90.01713130052181, 65651866.05340294, 4312295840576304.0, 100000)
}

func TestNegativeBinomial(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetNegativeBinomialSampler(29.5, 0.2)
	histogramTestInt(t,
		s,
		[]float64{-1., 70., 75., 78., 81., 83., 85., 86., 88., 89., 91., 92.,
			93., 94., 95., 96., 97., 98., 99., 100., 101., 102., 103., 104.,
			105., 106., 107., 108., 109., 110., 111., 112., 113., 114., 115., 116.,
			117., 118., 119., 120., 121., 122., 123., 124., 125., 126., 127., 128.,
			129., 130., 131., 132., 133., 135., 136., 137., 138., 139., 140., 142.,
			143., 145., 146., 148., 150., 152., 154., 157., 160., 164., 170., 178.,
			math.Inf(1)},
		[]float64{0.0140568, 0.01325281, 0.01166398, 0.01504103, 0.01211475, 0.01390623,
			0.00765021, 0.01673021, 0.0090886, 0.01963058, 0.01053825, 0.01101417,
			0.01148286, 0.01194218, 0.01239001, 0.0128243, 0.01324305, 0.01364435,
			0.0140264, 0.01438747, 0.014726, 0.01504054, 0.01532978, 0.01559257,
			0.01582793, 0.01603503, 0.01621319, 0.01636194, 0.01648093, 0.01657002,
			0.0166292, 0.01665863, 0.01665863, 0.01662966, 0.01657232, 0.01648733,
			0.01637555, 0.01623794, 0.01607556, 0.01588956, 0.01568118, 0.0154517,
			0.01520248, 0.01493491, 0.01465044, 0.01435051, 0.01403659, 0.01371016,
			0.01337268, 0.0130256, 0.01267036, 0.01230835, 0.02351037, 0.01119513,
			0.01081924, 0.01044292, 0.01006727, 0.00969334, 0.01827659, 0.00859129,
			0.01611459, 0.00753579, 0.01406421, 0.0127742, 0.01155321, 0.01040561,
			0.01362339, 0.01147304, 0.01239538, 0.01352973, 0.01116977, 0.01234946},
		1000000)
	momentTestInt(t, s, 118.0, 590, 3.20508, 100000)
}

func TestNormal(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetNormalSampler(-3.0, 5.0)
	histogramTest(t,
		s,
		[]float64{math.Inf(-1), -11.22426813, -9.40775783, -8.18216695, -7.20810617,
			-6.37244875, -5.62200256, -4.92660233, -4.26673552, -3.62830673,
			-3., -2.37169327, -1.73326448, -1.07339767, -0.37799744,
			0.37244875, 1.20810617, 2.18216695, 3.40775783, 5.22426813, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, -3.0, 25.0, 3.0, 100000)
}

func TestPareto(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetParetoSampler(2.0, 1.5)
	histogramTest(t,
		s,
		[]float64{2., 2.06957384, 2.14553197, 2.22886644, 2.32079442,
			2.42282746, 2.53686858, 2.66535109, 2.81144222, 2.97935007,
			3.1748021, 3.4058197, 3.6840315, 4.02702785, 4.46288633,
			5.0396842, 5.84803548, 7.08439046, 9.28317767, 14.73612599, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
}

func TestPoisson(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetPoissonSampler(15.0)
	histogramTestInt(t,
		s,
		[]float64{-1., 8., 9., 10., 11., 12., 13., 14., 15., 16., 17., 18., 19., 20., 21.,
			22., 23., math.Inf(1)},
		[]float64{0.03744649, 0.03240717, 0.04861075, 0.06628739, 0.08285923, 0.09560681,
			0.10243587, 0.10243587, 0.09603362, 0.08473555, 0.07061296, 0.05574707,
			0.04181031, 0.0298645, 0.02036216, 0.01327967, 0.01946457},
		1000000)
	momentTestInt(t, s, 15.0, 15.0, 3.066666, 100000)
	s, _ = r.GetPoissonSampler(55.0)
	histogramTestInt(t,
		s,
		[]float64{-1., 40., 42., 44., 45., 46., 47., 48., 49., 50., 51., 52., 53., 54., 55.,
			56., 57., 58., 59., 60., 61., 62., 63., 64., 65., 66., 67., 68., 71., math.Inf(1)},
		[]float64{0.02125359, 0.02031268, 0.03314629, 0.02250674, 0.02691024, 0.0314907,
			0.0360831, 0.04050144, 0.04455158, 0.04804582, 0.05081769, 0.05273534,
			0.05371192, 0.05371192, 0.05275278, 0.05090181, 0.04826896, 0.04499648,
			0.04124678, 0.03718972, 0.03299088, 0.02880156, 0.02475134, 0.02094344,
			0.01745287, 0.01432698, 0.011588, 0.02211631, 0.01589304},
		1000000)
	momentTestInt(t, s, 55.0, 55.0, 3.01818, 100000)
}

func TestPower(t *testing.T) {
	rng := pcg64dxsm.New(123, 456)
	r := New(rng)
	s, _ := r.GetPowerSampler(1.5, 2.3)
	histogramTest(t,
		s,
		[]float64{0., 0.40778011, 0.55119929, 0.65746111, 0.74506002, 0.82096765,
			0.88869487, 0.95029835, 1.00710294, 1.06002008, 1.10970783, 1.15665935,
			1.2012552, 1.24379624, 1.28452506, 1.32364057, 1.36130823, 1.3976674,
			1.4328367, 1.46691811, 1.5},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 1.045455, 0.1105132, 2.6153756, 100000)
}

func TestRayleigh(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetRayleighSampler(2.3)
	histogramTest(t,
		s,
		[]float64{0., 0.73667025, 1.05580029, 1.31127811, 1.53650863, 1.74461352,
			1.94258099, 2.13487312, 2.3247656, 2.51497822, 2.70804305, 2.90658071,
			3.11357607, 3.33273432, 3.569038, 3.82975121, 4.12648193, 4.48012605,
			4.93572186, 5.62981771, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 2.88262251582565, 2.270487431254997, 3.2450893006876385, 100000)
}

func TestT(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetTSampler(6.0)
	histogramTest(t,
		s,
		[]float64{math.Inf(-1), -1.94318028e+00, -1.43975575e+00, -1.13415693e+00,
			-9.05703285e-01, -7.17558196e-01, -5.53380924e-01, -4.04313361e-01,
			-2.64834533e-01, -1.31075653e-01, 6.91908149e-17, 1.31075653e-01,
			2.64834533e-01, 4.04313361e-01, 5.53380924e-01, 7.17558196e-01,
			9.05703285e-01, 1.13415693e+00, 1.43975575e+00, 1.94318028e+00, math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 0.0, 1.5, 6.0, 100000)
}

func TestTriangular(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetTriangularSampler(-2.3, 3.2, 1.0)
	histogramTest(t,
		s,
		[]float64{-2.3, -1.34737206, -0.95278064, -0.65, -0.39474411,
			-0.16985916, 0.03345238, 0.22041663, 0.39443872, 0.55788383,
			0.71247407, 0.85950946, 1., 1.14208844, 1.29474411,
			1.46074729, 1.64436508, 1.85278064, 2.1, 2.42218254,
			3.2},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 0.6333333333333333, 1.2772222222222223, 2.4, 100000)
}

func TestUniform(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetUniformSampler(-2.3, 3.2)
	histogramTest(t,
		s,
		[]float64{-2.3, -2.025, -1.75, -1.475, -1.2, -0.925, -0.65, -0.375, -0.1,
			0.175, 0.45, 0.725, 1., 1.275, 1.55, 1.825, 2.1, 2.375,
			2.65, 2.925, 3.2},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 0.4500000000000002, 2.520833333333333, 1.8, 100000)
}
func TestWeibull(t *testing.T) {
	rng := pcg64dxsm.New(1423, 456)
	r := New(rng)
	s, _ := r.GetWeibullSampler(2.3, 3.2)
	histogramTest(t,
		s,
		[]float64{0., 0.90912117, 1.13845187, 1.30357289, 1.43933022, 1.55825341,
			1.66652957, 1.76780195, 1.86450283, 1.95843859, 2.05109508, 2.14382858,
			2.23801627, 2.335213, 2.43736487, 2.54716845, 2.66878466, 2.80952215,
			2.98483438, 3.2406716,
			math.Inf(1)},
		[]float64{0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05,
			0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05, 0.05},
		10000000)
	momentTest(t, s, 2.060003440954738, 0.499263764754041, 2.713882062033474, 100000)
}

// Helper functions

// histogramTest draws N samples from sampler, bins them according to the
// boundaries in bins, and compares the observed bin counts to their expected
// probabilities. The upper bound of the bin is included, the lower is not.
//
// For each bin i, the expected probability mass is masses[i]. The test checks
// whether the observed count in each bin lies within a 99.5% confidence
// interval.
func histogramTest(t *testing.T, sampler func() float64, bins []float64, masses []float64, N int64) {
	k := len(bins) - 1
	counts := make([]int64, k)
	// --- Sampling ---
	for i := int64(0); i < N; i++ {
		x := sampler()
		for b := 0; b < k; b++ {
			if x <= bins[b+1] {
				counts[b]++
				break
			}
		}
	}
	// --- Check bins ---
	var failures int
	for index, c := range counts {
		p := masses[index]
		N := float64(N)
		mean := N * p
		std := math.Sqrt(N * p * (1 - p))
		const z = 3.27 // 99.9 confidence interval
		lo := mean - z*std
		hi := mean + z*std
		if float64(c) < lo || float64(c) > hi {
			fmt.Println(index, c, masses[index]*N)
			failures++
		}
	}
	if failures > 0 {
		t.Fatalf("%d/%d bins outside 99.9%% CI", failures, k)
	}
}

func histogramTestInt(t *testing.T, sampler func() int64, bins []float64, masses []float64, N int64) {
	histogramTest(t, func() float64 { return float64(sampler()) }, bins, masses, N)
}

// quantileTest draws samples from sampler, counts them into bins defined
// by quantiles, and checks each bin count against a 99.5% confidence interval.
func quantileTest(t *testing.T, sampler func() float64, bins []float64, samples int64) {
	k := len(bins) - 1
	counts := make([]int64, k)
	// --- Sampling ---
	for i := int64(0); i < samples; i++ {
		x := sampler()
		for b := 0; b < k; b++ {
			if x < bins[b+1] {
				counts[b]++
				break
			}
		}
	}
	// --- Statistics ---
	p := 1.0 / float64(k)
	N := float64(samples)
	mean := N * p
	std := math.Sqrt(N * p * (1 - p))
	const z = 3.27 // 99.9 confidence interval
	lo := mean - z*std
	hi := mean + z*std

	// --- Check bins ---
	var failures int
	for index, c := range counts {
		if float64(c) < lo || float64(c) > hi {
			fmt.Println(index, c)
			failures++
		}
	}
	if failures > 0 {
		t.Fatalf("%d/%d bins outside 99.9%% CI", failures, k)
	}
}

// momentTest samples from a continuous distribution and validates the first two moments.
//
// For samples X_1,…,X_N, it computes the empirical mean and the empirical second
// central moment, using the theoretical mean μ. The variance of the empirical
// mean is σ²/N, where σ² is the theoretical variance. The variance of the
// empirical second central moment is (κ − 1) σ⁴ / N, where κ is the theoretical
// kurtosis.
//
// By the central limit theorem, both estimators are approximately normally
// distributed for large N. Confidence intervals are constructed using these
// asymptotic normal approximations. The test fails if the empirical values
// fall outside the 99.9% confidence interval.
func momentTest(t *testing.T, sampler func() float64, theoreticalMean, theoreticalVar, theoreticalKurtosis float64, N int64) {
	var mean float64
	var variance float64
	var x float64
	for i := int64(0); i < N; i++ {
		x = sampler()
		mean += (x - mean) / float64(i+1)
		variance += ((x-theoreticalMean)*(x-theoreticalMean) - variance) / float64(i+1)
	}
	// Mean CI (normal approx)
	z := 3.27 // 99.9%
	meanErr := z * math.Sqrt(theoreticalVar/float64(N))
	varErr := z * math.Sqrt((theoreticalKurtosis-1)/float64(N)) * theoreticalVar
	// --- Perform the test ---
	if math.Abs(mean-theoreticalMean) > meanErr {
		t.Errorf("Mean out of bounds: got %.6g, expected %.6g ± %.6g", mean, theoreticalMean, meanErr)
	}
	if math.Abs(variance-theoreticalVar) > varErr {
		t.Errorf("Variance out of bounds: got %.6g, expected %.6g ± %.6g", variance, theoreticalVar, varErr)
	}
}

// momentTestInt samples from a discrete distribution and validates the first two moments.
//
// For samples X_1,…,X_N, it computes the empirical mean and the empirical second
// central moment, using the theoretical mean μ. The variance of the empirical
// mean is σ²/N, where σ² is the theoretical variance. The variance of the
// empirical second central moment is (κ − 1) σ⁴ / N, where κ is the theoretical
// kurtosis.
//
// By the central limit theorem, both estimators are approximately normally
// distributed for large N. Confidence intervals are constructed using these
// asymptotic normal approximations. The test fails if the empirical values
// fall outside the 99.9% confidence interval.
func momentTestInt(t *testing.T, sampler func() int64, theoreticalMean, theoreticalVar float64, theoreticalKurtosis float64, N int64) {
	momentTest(t, func() float64 { return float64(sampler()) }, theoreticalMean, theoreticalVar, theoreticalKurtosis, N)
}
