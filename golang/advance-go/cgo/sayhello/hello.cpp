# include <iostream>

extern "C" {
    # include "hello.h"
}

void sayHello(const char* s){
    std::cout << s;
    std::cout << "This is a Test";
}
