#include <iostream>
#include <cstdint>
#include "philox.h" // make sure this is in your include path

int main() {
    // Seed / stream values
    uint64_t key0 = 12345; // stream selector
    uint64_t key1 = 67890; // additional key word

    // Counter-based RNG: counter = {c0, c1, c2, c3}, key = {key0, key1}
    philox4x64_ctr_t counter = {{0, 0, 0, 0}}; 
    philox4x64_key_t key = {{key0, key1}};

    philox4x64_ctr_t counterT = {{1, 2, 3, 4}}; 
    philox4x64_key_t keyT = {{0,0}};
    philox4x64_ctr_t rndT = philox4x64(counterT, keyT);
    std::cout << rndT.v[0]<< " " << rndT.v[1]<<" " << rndT.v[2]<<" " << rndT.v[3] << std::endl;

    std::cout << "philox4x64data = [...]uint64{";
    for (int i = 0; i < 25; i++) {
        // Generate a block (4x64-bit numbers)
        philox4x64_ctr_t rnd = philox4x64(counter, key);

        // Print one of the outputs (say rnd.v[0])
        for (int j = 0; j < 4; j++) {
           std::cout << rnd.v[j] << ((i*4+j == 99) ? "}\n" : ",");
        }
        std::cout << std::endl;
        // Increment the counter
        counter.v[0]++;
        if (counter.v[0] == 0) {
            counter.v[1]++;
            if (counter.v[1] == 0) {
                counter.v[2]++;
                if (counter.v[2] == 0) {
                    counter.v[3]++;
                }
            }
        }
    }

    return 0;
}