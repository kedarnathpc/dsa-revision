func numWaterBottles(numBottles int, numExchange int) int {
    bottles, emptyBottles := numBottles, numBottles

    for emptyBottles >= numExchange {
        exchange := emptyBottles / numExchange
        bottles += exchange
        emptyBottles = exchange + (emptyBottles % numExchange)
    }

    return bottles
}
