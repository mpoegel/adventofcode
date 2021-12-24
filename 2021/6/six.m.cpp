#include <algorithm>
#include <cstdint>
#include <fstream>
#include <iostream>
#include <numeric>
#include <string>
#include <vector>

int readLanturnFish(const std::string &filename, std::vector<unsigned int> &lanturnFish)
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
        lanturnFish.push_back(fish);
        lastDelim = delim + 1;
        delim = line.find(',', lastDelim);
    }
    lanturnFish.push_back(std::stoul(line.substr(lastDelim, delim)));

    fp.close();
    return 0;
}

void countFishAfter(std::vector<unsigned int> lanturnFish, unsigned int days, unsigned int &numFish)
{
    for (unsigned int d = 0; d < days; ++d) {
        numFish = lanturnFish.size();
        for (unsigned int f = 0; f < numFish; ++f) {
            if (lanturnFish[f] == 0) {
                lanturnFish.push_back(8);
                lanturnFish[f] = 6;
            } else {
                --lanturnFish[f];
            }
        }
    }
    numFish = lanturnFish.size();
}

void countFishFaster(const std::vector<unsigned int> &lanturnFish, unsigned int days, uint64_t &numFish)
{
    constexpr int NUM_FISH_COUNTS = 9;
    uint64_t fishCounts[NUM_FISH_COUNTS] = {0};
    for (const auto& f : lanturnFish) {
        ++fishCounts[f];
    }

    numFish = 0;
    for (unsigned int d = 0; d < days; ++d) {
        std::rotate(fishCounts, fishCounts + 1, fishCounts + NUM_FISH_COUNTS);
        fishCounts[6] += fishCounts[8];
    }
    numFish = std::accumulate(fishCounts, fishCounts + NUM_FISH_COUNTS, uint64_t(0));
}

int main(int argc, char *argv[])
{
    if (argc < 2) {
        std::cerr << "missing filename\n";
        return 1;
    }

    const std::string inputFile(argv[1]);
    std::vector<unsigned int> lanturnFish;
    if (0 != readLanturnFish(inputFile, lanturnFish)) {
        std::cerr << "failed to open file: " << inputFile << '\n';
        return -1;
    }

    unsigned int days = 80;
    unsigned int numFish;
    countFishAfter(lanturnFish, days, numFish);
    std::cout << "Number of lanturn fish after " << days << " days: " << numFish << '\n';

    days = 256;
    uint64_t numFish2;
    countFishFaster(lanturnFish, days, numFish2);
    std::cout << "Number of lanturn fish after " << days << " days: " << numFish2 << '\n';

    return 0;
}
