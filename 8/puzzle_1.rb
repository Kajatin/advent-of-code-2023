file = File.open("puzzle_input.txt")
puzzle = file.read.split("\n")
file.close

Node = Struct.new(:field, :left, :right) do
    def to_s
        "Field: #{field}, Left: #{left}, Right: #{right}"
    end
end

instruction = ""
nodes = []

puzzle.each_with_index { |line, index|
    if index == 0 then
        instruction = line
        next
    end

    # If line is empty, skip.
    if line == "" then
        next
    end

    split = line.split(" = ")
    lr_split = split[1].split(", ")
    node = Node.new(split[0], lr_split[0].slice(1..-1), lr_split[1].slice(0..-2))
    nodes.push(node)
}

nodes.each_with_index { |node, index|
    nodes[index].right = nodes.map(&:field).index(nodes[index].right)
    nodes[index].left = nodes.map(&:field).index(nodes[index].left)
    nodes[index] = node
}

steps = 0
reached = false
current = nodes.map(&:field).index("AAA")
until reached do
    instruction.each_char { |char|
        if nodes[current].field == "ZZZ" then
            reached = true
            break
        end

        if char == "R" then
            current = nodes[current].right
        elsif char == "L" then
            current = nodes[current].left
        end

        steps += 1
    }
end

puts "Day 8, Puzzle 1: #{steps}"
