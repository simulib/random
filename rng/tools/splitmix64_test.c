
#include <stdio.h>
#include "splitmix64.c"

int main(){
      x = 987654321; // seed
      printf("splitmix64data := uint64[...]{");
      for (int i=0; i<40; i++){
         for (int j=0; j < 5; j++) printf("%llu,",next());
         printf("\n");
      }
      printf("%llu}\n",next());
}