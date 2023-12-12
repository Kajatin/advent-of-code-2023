import Foundation

struct Strength {
    let value: Int
    let name: String

    init(cards: [Card]) {
        let letters = cards.map({ $0.letter })
        var letterCounts = [String: Int]()
        for letter in letters {
            letterCounts[letter, default: 0] += 1
        }

        if letterCounts.values.contains(5) {
            self.value = 7
            self.name = "Five of a kind"
        } else if letterCounts.values.contains(4) {
            self.value = 6
            self.name = "Four of a kind"
        } else if letterCounts.values.contains(3) && letterCounts.values.contains(2) {
            self.value = 5
            self.name = "Full house"
        } else if letterCounts.values.contains(3) && !letterCounts.values.contains(2) {
            self.value = 4
            self.name = "Three of a kind"
        } else if letterCounts.values.filter({ $0 == 2 }).count == 2 {
            self.value = 3
            self.name = "Two pairs"
        } else if letterCounts.values.contains(2) && !letterCounts.values.contains(3) {
            self.value = 2
            self.name = "One pair"
        } else {
            self.value = 1
            self.name = "High card"
        }

    }
}

struct Card {
    let value: Int
    let letter: String

    static func parse_value(_ letter: String) -> Int {
        switch letter {
        case "A":
            return 14
        case "K":
            return 13
        case "Q":
            return 12
        case "J":
            return 11
        case "T":
            return 10
        default:
            return Int(letter)!
        }
    }
}

struct Hand {
    let bid: Int
    let cards: [Card]
    let strength: Strength
}

var hands: [Hand] = []

if let filePath = Bundle.main.path(forResource: "puzzle_input", ofType: "txt") {
    do {
        let contents = try String(contentsOfFile: filePath)

        let handsStringArray = contents.components(separatedBy: "\n")
        for handString in handsStringArray {
            let handStringArray = handString.components(separatedBy: " ")
            var cards: [Card] = []
            for cardString in handStringArray[0] {
                let value = Card.parse_value(String(cardString))
                let letter = String(cardString)
                cards.append(Card(value: value, letter: letter))
            }
            let bid = Int(handStringArray[1])!
            let strength = Strength(cards: cards)
            hands.append(Hand(bid: bid, cards: cards, strength: strength))
        }
    } catch {
        print("Error reading file: \(error.localizedDescription)")
    }
}

func handSorting(_ lhs: Hand, _ rhs: Hand) -> Bool {
    if (lhs.strength.value == rhs.strength.value) {
        for i in 0..<lhs.cards.count {
            if (lhs.cards[i].value == rhs.cards[i].value) {
                continue
            } else {
                return lhs.cards[i].value < rhs.cards[i].value
            }
        }
    }

    return lhs.strength.value < rhs.strength.value
}

hands.sort(by: { handSorting($0, $1) })
var summa = 0
for (index, hand) in hands.enumerated() {
    summa += hand.bid * (index + 1)
    print("\(index + 1): \(hand.bid) * \(index + 1) = \(hand.bid * (index + 1))")
}
print("Day 7 Puzzle 1: \(summa)")
