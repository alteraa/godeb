#include <iostream>
#include <thread>
#include <chrono>

int main() {
    while (true) {
        std::cout << "Deb package test program running..." << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(5));
    }
    return 0;
}
