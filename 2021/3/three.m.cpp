#include <fstream>
#include <functional>
#include <iostream>
#include <string>
#include <unordered_set>
#include <vector>

int readDiagnostic(const std::string& filename, std::vector<std::string> &diagnostic)
{
    std::ifstream fp;
    fp.open(filename);
    if (!fp.is_open()) {
        return -1;
    }
    while (fp.good()) {
        std::string reading;
        fp >> reading;
        diagnostic.push_back(reading);
    }
    fp.close();
    return 0;
}

void calculatePowerConsumption(const std::vector<std::string> &diagnostic,
                               unsigned int &gammaRate,
                               unsigned int &epsilonRate)
{
    auto totalReadings = diagnostic.size();
    auto readingLength = diagnostic[0].size();
    std::vector<unsigned int> bitSum(readingLength, 0);

    for (auto reading : diagnostic) {
        for (std::size_t i = 0; i < readingLength; ++i) {
            bitSum[i] += reading[i] == '1' ? 1 : 0;
        }
    }

    std::string gammaRateStr(readingLength, '0');
    std::string epsilonRateStr(readingLength, '1');
    unsigned int halfNumReadings = totalReadings / 2;
    for (std::size_t i = 0; i < readingLength; ++i) {
        gammaRateStr[i] = bitSum[i] > halfNumReadings ? '1' : '0';
        epsilonRateStr[i] = bitSum[i] > halfNumReadings ? '0' : '1';
    }

    gammaRate = std::stoul(gammaRateStr, nullptr, 2);
    epsilonRate = std::stoul(epsilonRateStr, nullptr, 2);
}

unsigned int setRemoveIf(std::unordered_set<std::string> &set, std::function<bool(const std::string& item)> predicate)
{
    unsigned int count = 0;
    for (auto itr = set.begin(); itr != set.end(); ) {
        if (predicate(*itr)) {
            itr = set.erase(itr);
            ++count;
        } else {
            ++itr;
        }
    }
    return count;
}

void calculateOxygenAndCO2Scrubbing(const std::vector<std::string> &diagnostic,
                                    unsigned int &oxygen,
                                    unsigned int &scrubbing)
{
    const auto totalReadings = diagnostic.size();
    double halfNumReadings = totalReadings / 2.0;
    auto readingLength = diagnostic[0].size();
    std::unordered_set<std::string> oxygenReadings(diagnostic.begin(), diagnostic.end());
    std::unordered_set<std::string> scrubbingReadings(oxygenReadings);

    char removeBit;
    unsigned int bitIndex = 0;
    while (oxygenReadings.size() > 1 && bitIndex < readingLength) {
        unsigned int sum = 0;
        halfNumReadings = oxygenReadings.size() / 2.0;
        for (auto reading : oxygenReadings) sum += reading[bitIndex] == '1' ? 1 : 0;
        removeBit = sum >= halfNumReadings ? '0' : '1';
        auto removePredicate = [removeBit, bitIndex](const std::string& reading) { return reading[bitIndex] == removeBit; };
        setRemoveIf(oxygenReadings, removePredicate);
        ++bitIndex;
    }
    oxygen = std::stoul(*oxygenReadings.begin(), nullptr, 2);

    bitIndex = 0;
    while (scrubbingReadings.size() > 1 && bitIndex < readingLength) {
        unsigned int sum = 0;
        halfNumReadings = scrubbingReadings.size() / 2.0;
        for (auto reading : scrubbingReadings) sum += reading[bitIndex] == '1' ? 1 : 0;
        removeBit = sum >= halfNumReadings ? '1' : '0';
        auto removePredicate = [removeBit, bitIndex](const std::string& reading) { return reading[bitIndex] == removeBit; };
        setRemoveIf(scrubbingReadings, removePredicate);
        ++bitIndex;
    }
    scrubbing = std::stoul(*scrubbingReadings.begin(), nullptr, 2);
}

int main(int argc, char *argv[]) {
    if (argc < 2) {
        std::cerr << "missing filename\n";
        return 1;
    }

    const std::string inputFile(argv[1]);
    std::vector<std::string> diagnostic;
    if (0 != readDiagnostic(inputFile, diagnostic)) {
        std::cerr << "failed to open input file: " << inputFile << '\n';
        return 1;
    }

    unsigned int gamma;
    unsigned int epsilon;
    calculatePowerConsumption(diagnostic, gamma, epsilon);
    std::cout << "gamma=" << gamma << " epsilon=" << epsilon << " answer=" << gamma * epsilon << '\n';

    unsigned int oxygenLevel;
    unsigned int co2Scrubbing;
    calculateOxygenAndCO2Scrubbing(diagnostic, oxygenLevel, co2Scrubbing);
    std::cout << "oxygen=" << oxygenLevel << " CO2=" << co2Scrubbing << " life support rating=" << oxygenLevel * co2Scrubbing << '\n';

    return 0;
}
