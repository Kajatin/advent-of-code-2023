// g++ -std=c++20 -o puzzle_1 puzzle_1.cpp && ./puzzle_1

#include <regex>
#include <fstream>
#include <numeric>
#include <iostream>

struct Cubes
{
    int red;
    int green;
    int blue;
};

struct Game
{
    int id;
    std::vector<Cubes> cubes;
};

std::vector<std::string> read_puzzle()
{
    std::string line;
    std::vector<std::string> lines;
    std::ifstream puzzle("puzzle_input.txt");

    while (getline(puzzle, line))
    {
        lines.push_back(line);
    }

    puzzle.close();

    return lines;
}

std::vector<Game> parse_games(std::vector<std::string> puzzle)
{
    std::vector<Game> games;

    for (auto game_str : puzzle)
    {
        Game game;

        // Parse id
        int id = 0;
        std::cmatch m;
        std::regex game_id("Game\\s+(\\d+)");
        if (std::regex_search(game_str.c_str(), m, game_id))
        {
            id = std::stoi(m[1]);
        }
        game.id = id;

        auto colon = game_str.find_first_of(":");
        auto sub_game_str = game_str.substr(colon + 2);
        while (true)
        {
            Cubes cubes;

            auto game_split = sub_game_str.find_first_of(";");
            auto sub_game_split_str = sub_game_str.substr(0, game_split);
            sub_game_str = sub_game_str.substr(game_split + 2);

            cubes.red = 0;
            std::regex red_num("(\\d+)\\s+red");
            if (std::regex_search(sub_game_split_str.c_str(), m, red_num))
            {
                if (2 <= m.size())
                {
                    cubes.red = std::stoi(m[1]);
                }
            }

            cubes.green = 0;
            std::regex green_num("(\\d+)\\s+green");
            if (std::regex_search(sub_game_split_str.c_str(), m, green_num))
            {
                if (2 <= m.size())
                {
                    cubes.green = std::stoi(m[1]);
                }
            }

            cubes.blue = 0;
            std::regex blue_num("(\\d+)\\s+blue");
            if (std::regex_search(sub_game_split_str.c_str(), m, blue_num))
            {
                if (2 <= m.size())
                {
                    cubes.blue = std::stoi(m[1]);
                }
            }

            game.cubes.push_back(cubes);

            if (std::string::npos == game_split)
            {
                break;
            }
        }

        games.push_back(game);
    }

    return games;
}

std::vector<Game> find_possible_games(std::vector<Game> games)
{
    std::vector<Game> possible_games;

    std::for_each(
        games.begin(),
        games.end(),
        [&possible_games](Game g)
        {
            if (std::none_of(
                    g.cubes.begin(),
                    g.cubes.end(),
                    [](Cubes c)
                    {
                        // Make sure there are never more cubes shown than possible
                        return !(c.red <= 12 && c.green <= 13 && c.blue <= 14);
                    }))
            {
                possible_games.push_back(g);
            }
        });

    return possible_games;
}

int main()
{
    auto puzzle = read_puzzle();
    auto games = parse_games(puzzle);
    auto possible_games = find_possible_games(games);
    int summa = std::accumulate(possible_games.begin(), possible_games.end(), 0, std::bind(std::plus<int>(), std::placeholders::_1, std::bind(&Game::id, std::placeholders::_2)));

    std::cout << "Day 2 Puzzle 1: " << summa << std::endl;

    return 0;
}
