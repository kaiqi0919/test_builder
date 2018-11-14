#include <stdlib.h>
#include <unistd.h>

void main(){
    chdir("../bin");
    system("docm_generator_achievement.docm");
}
