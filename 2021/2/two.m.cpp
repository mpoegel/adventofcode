#include <fstream>
#include <iostream>
#include <string>
#include <vector>

enum class Direction {
    Unknown = -1,
    Forward,
    Down,
    Up,
};

Direction stringToDirection(const std::string& dir)
{
    if (dir == "forward") {
        return Direction::Forward;
    }
    if (dir == "down") {
        return Direction::Down;
    }
    if (dir == "up") {
        return Direction::Up;
    }
    return Direction::Unknown;
}

struct Instruction {
    Direction dir;
    unsigned int distance;
};

int readInstructions(const std::string& filename, std::vector<Instruction>& instructions)
{
    std::ifstream fp;
    fp.open(filename);
    if (!fp.is_open()) {
        return -1;
    }
    while (fp.good()) {
        std::string dirStr;
        unsigned int distance;
        fp >> dirStr;
        fp >> distance;
        instructions.emplace_back(Instruction{stringToDirection(dirStr), distance});
    }
    fp.close();
    return 0;
}

void calculateEndPosition(const std::vector<Instruction>& instructions,
                          unsigned int& horizontal,
                          unsigned int& depth)
{
    horizontal = 0;
    depth = 0;
    for (auto instruction : instructions) {
        switch (instruction.dir) {
            case Direction::Forward:
                horizontal += instruction.distance;
                break;
            case Direction::Up:
                depth -= instruction.distance;
                break;
            case Direction::Down:
                depth += instruction.distance;
                break;
        }
    }
}

void calculateTrueEndPosition(const std::vector<Instruction>& instructions,
                              unsigned int& horizontal,
                              unsigned int& depth)
{
    horizontal = 0;
    depth = 0;
    int aim = 0;
    for (auto instruction : instructions) {
        switch (instruction.dir) {
            case Direction::Forward:
                horizontal += instruction.distance;
                depth += aim * instruction.distance;
                break;
            case Direction::Up:
                aim -= instruction.distance;
                break;
            case Direction::Down:
                aim += instruction.distance;
                break;
        }
    }
}

int main(int argc, char *argv[])
{
    if (argc < 2) {
        std::cerr << "missing filename\n";
        return 1;
    }

    const std::string inputFile(argv[1]);
    std::vector<Instruction> instructions;
    if (0 != readInstructions(inputFile, instructions)) {
        std::cerr << "failed to open input file: " << inputFile << '\n';
        return 1;
    }

    unsigned int horizontal;
    unsigned int depth;
    calculateEndPosition(instructions, horizontal, depth);
    std::cout << "horizontal=" << horizontal << " depth=" << depth << " answer=" << horizontal*depth << '\n';

    calculateTrueEndPosition(instructions, horizontal, depth);
    std::cout << "horizontal=" << horizontal << " depth=" << depth << " answer=" << horizontal*depth << '\n';

    return 0;
}
