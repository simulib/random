#include <stdio.h>
#include "philox.h"

int main() {
    philox4x64_key_t k = {{0, 0}};
    philox4x64_ctr_t c = {{0, 0, 0, 0}};

    philox4x64_ctr_t r = philox4x64(c, k);

    printf("%016llx %016llx %016llx %016llx\n",
        (unsigned long long)r.v[0],
        (unsigned long long)r.v[1],
        (unsigned long long)r.v[2],
        (unsigned long long)r.v[3]
    );
}
