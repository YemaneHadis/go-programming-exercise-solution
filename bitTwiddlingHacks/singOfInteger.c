#include<stdio.h>
int main(int argc, char const *argv[])
{
    int sing1;
    int sing2;

    int u = 12;
    int v = -12;

     sing1 = u >> (sizeof(int) * __CHAR_BIT__-1);
     sing2 = v >> (sizeof(int) * __CHAR_BIT__-1);

     printf("Sing of U\t%d\n", sing1);
     printf("Sing of V\t%d\n", sing2);



    return 0;
}
