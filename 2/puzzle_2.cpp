// g++ -std=c++20 -o puzzle_2 puzzle_2.cpp && ./puzzle_2

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

std::vector<Cubes> find_fewest_cubes_to_make_a_possible_game(std::vector<Game> games)
{
    std::vector<Cubes> min_cubes;

    std::for_each(
        games.begin(),
        games.end(),
        [&min_cubes](Game g)
        {
            Cubes cubes;

            auto it = std::max_element(
                g.cubes.begin(),
                g.cubes.end(),
                [](Cubes lhs, Cubes rhs)
                {
                    return lhs.red < rhs.red;
                });
            cubes.red = it->red;

            it = std::max_element(
                g.cubes.begin(),
                g.cubes.end(),
                [](Cubes lhs, Cubes rhs)
                {
                    return lhs.green < rhs.green;
                });
            cubes.green = it->green;

            it = std::max_element(
                g.cubes.begin(),
                g.cubes.end(),
                [](Cubes lhs, Cubes rhs)
                {
                    return lhs.blue < rhs.blue;
                });
            cubes.blue = it->blue;

            min_cubes.push_back(cubes);
        });

    return min_cubes;
}

std::vector<int> compute_power_of_games(std::vector<Cubes> min_cubes_for_each_game)
{
    std::vector<int> power_of_games;

    std::transform(
        min_cubes_for_each_game.begin(),
        min_cubes_for_each_game.end(),
        std::back_inserter(power_of_games),
        [](Cubes cubes)
        {
            return cubes.red * cubes.green * cubes.blue;
        });

    return power_of_games;
}

int main()
{
    auto puzzle = read_puzzle();
    auto games = parse_games(puzzle);
    auto min_cubes_for_each_game = find_fewest_cubes_to_make_a_possible_game(games);
    auto power_of_games = compute_power_of_games(min_cubes_for_each_game);
    int summa = std::accumulate(power_of_games.begin(), power_of_games.end(), 0);

    std::cout << "Day 2 Puzzle 2: " << summa << std::endl;

    return 0;
}
