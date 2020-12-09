function makeComplexRatFunc(top, bot) {
  return λ => {
    let t = 0
    let b = 0
    for (let i = 0; i < top.length; i++) {
      t = math.add(t, math.multiply(top[i], math.pow(λ, i)))
    }
    for (let i = 0; i < bot.length; i++) {
      b = math.add(b, math.multiply(bot[i], math.pow(λ, i)))
    }
    return math.divide(t, b)
  }
}