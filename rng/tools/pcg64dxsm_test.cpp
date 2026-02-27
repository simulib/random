#include <iostream>
#include <cstdint>

#include "pcg_random.hpp"
#include "pcg_extras.hpp"

int main() {
    using pcg_dxsm_128_64 = pcg_engines::setseq_dxsm_128_64;

    uint64_t seed1 = 9876ULL;
    uint64_t seed2 = 54321ULL;

    __uint128_t seed = ((__uint128_t)seed1 << 64) | seed2;
    pcg_dxsm_128_64 rng(seed);

    //std::cout << rng <<"\n";

    printf("pcg64dxsmdata := [...]uint64{");
    for (int i=0; i<10; i++){
        for (int j=0; j < 5; j++) std::cout << rng() << ",";
        std::cout << "\n";
                if (i % 2) rng.advance((__uint128_t)1ULL << 64);
                else rng.advance((__uint128_t)1ULL << 96);
    }
    std::cout << rng() << "}\n";
    
    return 0;
}