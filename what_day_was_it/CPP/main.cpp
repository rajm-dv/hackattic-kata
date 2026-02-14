#include <iostream>
#include <string>

using ll = long long;

void countDays(std::string input) {
  std::string days[7] = {"Thursday", "Friday",  "Saturday", "Sunday",
                         "Monday",   "Tuesday", "Wednesday"};
  int n = std::stoi(input);
  int idx = ((n % 7) + 7) % 7;
  
  std::cout << days[idx] << "\n";
}

int main(void) {
  std::string line;
  while (std::getline(std::cin, line)) {
    countDays(line);
  }
}
