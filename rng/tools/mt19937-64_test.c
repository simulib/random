
#include <stdio.h>
#include "mt19937-64.h"

int main(){
     init_genrand64(987654321); 
      printf("mt19937x64data := uint64[...]{");
      for (int i=0; i<40; i++){
         for (int j=0; j < 5; j++) printf("%llu,",genrand64_int64());
         printf("\n");
      }
      printf("%llu}\n",genrand64_int64());
}