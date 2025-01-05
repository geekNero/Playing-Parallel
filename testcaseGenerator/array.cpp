#include <iostream>
#include <cstdlib>
#include <ctime>

using namespace std;

int main() {  
    srand(time(0));
    int n;
    cin >> n;
    cout << n << endl;
    for (int i = 0; i < n; i++) {
        cout << (rand() % 2000001) - 1000000 << " ";
    }
    cout << endl;
    return 0;
}