#include <algorithm>
#include <fstream>
#include <cmath>
#include <iostream>
#include <string>
#include <vector>

int readCrabPositions(const std::string &filename, std::vector<unsigned int> &crabPositions)
{
    std::fstream fp;
    fp.open(filename);
    if (!fp.is_open()) {
        return -1;
    }

    std::string line;
    std::getline(fp, line);
    std::size_t lastDelim = 0;
    std::size_t delim = line.find(',', 0);
    while (delim != std::string::npos) {
        unsigned int fish = std::stoul(line.substr(lastDelim, delim - lastDelim));
        crabPositions.push_back(fish);
        lastDelim = delim + 1;
        delim = line.find(',', lastDelim);
    }
    crabPositions.push_back(std::stoul(line.substr(lastDelim, delim)));

    fp.close();
    return 0;
}

void findOptimalPosition(std::vector<unsigned int> crabPositions, unsigned int &position, unsigned int &fuelCost)
{
    position = 0;
    fuelCost = 0;
    std::sort(crabPositions.begin(), crabPositions.end());
    unsigned int half = crabPositions.size() / 2;
    position = crabPositions[half];
    for (auto crab : crabPositions) {
        fuelCost += crab > position ? crab - position : position - crab;
    }
}

unsigned int sumTo(unsigned int n) {
    unsigned int sum = 0;
    for (unsigned int i = 1; i <= n; ++i) sum += i;
    return sum;
}

void findOptimalPosition2(std::vector<unsigned int> crabPositions, unsigned int &position, unsigned int &fuelCost)
{
    position = 0;
    fuelCost = 0;
    unsigned int sum = 0;
    for (auto crab : crabPositions) {
        sum += crab;
    }
    position = std::round(sum / crabPositions.size());
    for (auto crab : crabPositions) {
        fuelCost += crab > position ? sumTo(crab - position) : sumTo(position - crab);
    }
}

int main(int argc, char *argv[])
{
    if (argc < 2) {
        std::cerr << "missing filename\n";
        return 1;
    }

    const std::string inputFile(argv[1]);
    std::vector<unsigned int> crabPositions;
    if (0 != readCrabPositions(inputFile, crabPositions)) {
        std::cerr << "failed to open file: " << inputFile << '\n';
        return -1;
    }

    unsigned int position;
    unsigned int fuelCost;
    findOptimalPosition(crabPositions, position, fuelCost);
    std::cout << "optimal position: " << position << ", fuel cost: " << fuelCost << '\n';

    findOptimalPosition2(crabPositions, position, fuelCost);
    std::cout << "optimal position: " << position << ", fuel cost: " << fuelCost << '\n';

    return 0;
}
