#include <fstream>
#include <iostream>
#include <string>
#include <vector>

int loadFile(const std::string& filename, std::vector<int>& data)
{
    std::ifstream fp;
    fp.open(filename);
    if (!fp.is_open()) {
        return -1;
    }
    while (fp.good()) {
        std::string line;
        std::getline(fp, line);
        data.push_back(std::stoi(line));
    }
    fp.close();
    return 0;
}

unsigned int countIncreasingMeasures(const std::vector<int>& data)
{
    unsigned int count = 0;
    for (std::size_t i = 1; i < data.size(); ++i) {
        if (data[i] > data[i-1]) {
            ++count;
        }
    }
    return count;
}

unsigned int countSlidingWindowIncreases(const std::vector<int>& data)
{
    const unsigned int WINDOW_SIZE = 3;
    unsigned int count = 0;
    if (data.size() < WINDOW_SIZE) {
        return count;
    }

    unsigned int prevSum = data[0] + data[1] + data[2];
    unsigned int currSum = prevSum;
    for (std::size_t i = WINDOW_SIZE; i < data.size(); ++i) {
        currSum += data[i] - data[i - WINDOW_SIZE];
        if (currSum > prevSum) {
            ++count;
        }
        prevSum = currSum;
    }
    return count;
}

int main(int argc, char *argv[])
{
    if (argc < 2) {
        std::cerr << "missing filename\n";
        return 1;
    }

    const std::string inputFile(argv[1]);
    std::vector<int> data;
    int rc = loadFile(inputFile, data);
    if (0 != rc) {
        std::cerr << "failed to open file: " << inputFile << '\n';
        return 1;
    }
    std::cout << "read " << data.size() << " lines\n";

    unsigned int numIncreasingMeasures = countIncreasingMeasures(data);
    std::cout << "number increasing measures: " << numIncreasingMeasures << '\n';

    unsigned int numSlidingIncreases = countSlidingWindowIncreases(data);
    std::cout << "number increasing with sliding window: " << numSlidingIncreases << '\n';

    return 0;
}
