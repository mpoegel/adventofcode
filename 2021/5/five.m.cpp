#include <fstream>
#include <iostream>
#include <string>
#include <vector>

struct Point {
    unsigned int x;
    unsigned int y;
};

struct Line {
    Point start;
    Point end;

    bool isHorizontal() const;
    bool isVertical() const;
    void print() const;

    void pointsOnLine(std::vector<Point> &points) const;
};

bool Line::isHorizontal() const
{
    return start.y == end.y;
}

bool Line::isVertical() const
{
    return start.x == end.x;
}

void Line::print() const
{
    std::cout << start.x << ',' << start.y << " -> " << end.x << ',' << end.y << '\n';
}

void Line::pointsOnLine(std::vector<Point> &points) const
{
    points.push_back(start);
    Point p = start;
    if (isHorizontal()) {
        while (p.x != end.x) {
            p.x += p.x < end.x ? 1 : -1;
            points.push_back(p);
        }
    }
    else if (isVertical()) {
        while (p.y != end.y) {
            p.y += p.y < end.y ? 1 : -1;
            points.push_back(p);
        }
    }
    else {
        while (p.x != end.x && p.y != end.y) {
            p.x += p.x < end.x ? 1 : -1;
            p.y += p.y < end.y ? 1 : -1;
            points.push_back(p);
        }
    }
}

int loadLines(const std::string &filename, std::vector<Line> &lines)
{
    std::ifstream fp;
    fp.open(filename);
    if (!fp.is_open()) {
        return -1;
    }

    std::string line;
    while (fp.good()) {
        std::getline(fp, line);
        auto delim = line.find(',', 0);
        unsigned int x = std::stoul(line.substr(0, delim));
        auto nextDelim = line.find(' ', delim + 1);
        unsigned int y = std::stoul(line.substr(delim + 1, nextDelim - delim));
        Point start{x, y};

        delim = line.find(' ', nextDelim + 1);
        nextDelim = line.find(',', delim);
        x = std::stoul(line.substr(delim + 1, nextDelim - delim));
        y = std::stoul(line.substr(nextDelim + 1, std::string::npos));
        Point end{x, y};

        lines.push_back(Line{start, end});
    }

    return 0;
}

void countOverlappingVents(const std::vector<Line> lines, unsigned int &overlaps, bool useDiagonal)
{
    overlaps = 0;
    unsigned int maxX = 0;
    unsigned int maxY = 0;
    for (auto ln : lines) {
        maxX = std::max(maxX, ln.start.x);
        maxY = std::max(maxY, ln.start.y);
        maxX = std::max(maxX, ln.end.x);
        maxY = std::max(maxY, ln.end.y);
    }

    std::vector<std::vector<unsigned int>> seafloor(maxX + 1, std::vector<unsigned int>(maxY + 1, 0));
    for (auto ln : lines) {
        if (useDiagonal || ln.isVertical() || ln.isHorizontal()) {
            std::vector<Point> points;
            ln.pointsOnLine(points);
            for (auto p : points) {
                seafloor[p.x][p.y]++;
                if (seafloor[p.x][p.y] == 2) {
                    overlaps++;
                }
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
    std::vector<Line> lines;
    if (0 != loadLines(inputFile, lines)) {
        std::cerr << "failed to open file: " << inputFile << '\n';
        return -1;
    }

    unsigned int overlaps;
    countOverlappingVents(lines, overlaps, false);
    std::cout << "number of overlapping vents: " << overlaps << '\n';

    countOverlappingVents(lines, overlaps, true);
    std::cout << "number of overlapping vents (with diagnals): " << overlaps << '\n';

    return 0;
}
