#include <fstream>
#include <iostream>
#include <string>
#include <unordered_set>
#include <vector>

const unsigned int BINGO_BOARD_SIZE = 5;

struct BingoBoard {
    std::vector<std::unordered_set<unsigned int>> rows;
    std::vector<std::unordered_set<unsigned int>> columns;
    bool isWinner;

    BingoBoard();

    void reset();
    bool callNumber(unsigned int num);

    void printByRow();
    void printByColumn();
};

BingoBoard::BingoBoard()
: rows(BINGO_BOARD_SIZE, std::unordered_set<unsigned int>())
, columns(BINGO_BOARD_SIZE, std::unordered_set<unsigned int>())
, isWinner(false)
{
}

void BingoBoard::reset()
{
    for (std::size_t i = 0; i < rows.size(); ++i) rows[i].clear();
    for (std::size_t i = 0; i < columns.size(); ++i) columns[i].clear();
}

bool BingoBoard::callNumber(unsigned int num)
{
    if (isWinner) {
        return false;
    }
    bool isWinnerByRow = false;
    for (std::size_t i = 0; i < rows.size(); ++i) {
        if (rows[i].erase(num) == 1) {
            isWinnerByRow = rows[i].size() == 0;
            break;
        }
    }

    bool isWinnerByCol = false;
    for (std::size_t i = 0; i < columns.size(); ++i) {
        if (columns[i].erase(num) == 1) {
            isWinnerByCol = columns[i].size() == 0;
            break;
        }
    }
    isWinner = isWinnerByCol || isWinnerByRow;
    return isWinner;
}

void BingoBoard::printByRow()
{
    for (auto row : rows) {
        for (auto num : row) {
            std::cout << num << " ";
        }
        std::cout << '\n';
    }
}

void BingoBoard::printByColumn()
{
    for (auto col : columns) {
        for (auto num : col) {
            std::cout << num << " ";
        }
        std::cout << '\n';
    }
}

int loadBingo(const std::string &filename,
              std::vector<unsigned int> &drawnNumbers,
              std::vector<BingoBoard> &boards)
{
    std::ifstream fp;
    fp.open(filename);
    if (!fp.is_open()) {
        return -1;
    }

    std::string drawnNumStr;
    std::getline(fp, drawnNumStr);
    std::size_t index = 0;
    while (index < drawnNumStr.size()) {
        auto delim = drawnNumStr.find(',', index);
        if (delim != std::string::npos) {
            drawnNumbers.push_back(std::stoul(drawnNumStr.substr(index, delim - index)));
        } else {
            drawnNumbers.push_back(std::stoul(drawnNumStr.substr(index, std::string::npos)));
            break;
        }
        index = delim + 1;
    }
    // skip blank line
    std::getline(fp, drawnNumStr);

    std::string line;
    BingoBoard curr;
    unsigned int currRow = 0;
    do {
        std::getline(fp, line);
        if (line.length() == 0) {
            boards.push_back(curr);
            curr.reset();
            currRow = 0;
        } else {
            std::size_t index = 0;
            unsigned int currCol = 0;
            while (index != std::string::npos) {
                while (index < line.length() && line[index] == ' ') index++;
                std::size_t delim = line.find(' ', index);
                unsigned int num;
                if (delim != std::string::npos) {
                    num = std::stoul(line.substr(index, delim - index));
                } else {
                    num = std::stoul(line.substr(index, delim));
                }
                curr.rows[currRow].insert(num);
                curr.columns[currCol++].insert(num);
                index = delim;
            }
            ++currRow;
        }
    } while (fp.good());
    boards.push_back(curr);

    fp.close();
    return 0;
}

void getWinningScore(const std::vector<unsigned int> &drawnNumbers,
                     std::vector<BingoBoard> boards,
                     unsigned int &winningSum,
                     unsigned int &winningDraw)
{
    winningSum = 0;
    winningDraw = 0;
    for (unsigned int draw : drawnNumbers) {
        for (std::size_t i = 0; i < boards.size(); ++i) {
            if (boards[i].callNumber(draw)) {
                for (auto row : boards[i].rows) {
                    for (unsigned int num : row) winningSum += num;
                }
                winningDraw = draw;
                return;
            }
        }
    }

    std::cout << "no winner found\n";
    for (auto board : boards) {
        board.printByRow();
        std::cout << '\n';
    }
}

void getLosingScore(const std::vector<unsigned int> &drawnNumbers,
                    std::vector<BingoBoard> boards,
                    unsigned int &losingSum,
                    unsigned int &losingDraw)
{
    losingSum = 0;
    losingDraw = 0;
    unsigned int numWins = 0;
    for (unsigned int draw : drawnNumbers) {
        for (std::size_t i = 0; i < boards.size(); ++i) {
            if (boards[i].callNumber(draw) && ++numWins == boards.size()) {
                for (auto row : boards[i].rows) {
                    for (unsigned int num : row) losingSum += num;
                }
                losingDraw = draw;
                return;
            }
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

    std::vector<unsigned int> drawnNumbers;
    std::vector<BingoBoard> boards;
    if (0 != loadBingo(inputFile, drawnNumbers, boards)) {
        std::cerr << "failed to open file: " << inputFile << '\n';
        return 1;
    }

    unsigned int winningSum;
    unsigned int winningDraw;
    getWinningScore(drawnNumbers, boards, winningSum, winningDraw);
    std::cout << "sum=" << winningSum << " draw=" << winningDraw << " score=" << winningDraw * winningSum << '\n';

    unsigned int losingSum;
    unsigned int losingDraw;
    getLosingScore(drawnNumbers, boards, losingSum, losingDraw);
    std::cout << "sum=" << losingSum << " draw=" << losingDraw << " score=" << losingDraw * losingSum << '\n';

    return 0;
}
