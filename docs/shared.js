
function PolynomialString(vector) {
  let poly = ""
  for (let i = 0; i < vector.length; i++) {
    if (vector[i] == 0) {
      continue
    }
    if (vector[i] != 1 || i == 0) {
      let num = `${vector[i]}`
      let index = num.indexOf("e")
      if (index != -1) {
        num = `${num.slice(0, index)}\\cdot 10^{${num.slice(index + 2, num.length)}}`
      }
      poly += num
    }
    if (i > 1) {
      poly += `x^{${i}}`
    } else if (i == 1) {
      poly += `x`
    }
    poly += "+"
  }
  return poly.replace(/\++$/g, "")
}

function parseVector(text) {
  text = text.replace(/</g, "[").replace(/>/, "]").replace(/\(/g, "[").replace(/\)/g, "]")
  return JSON.parse(text)
}

function setPQ() {
  try {
    p = parseVector(document.querySelector("#p").value)
  } catch {

  }
  try {
    q = parseVector(document.querySelector("#q").value)
  } catch { 
    
  }
  draw()
}

function makeRationalFunc(top, bot) {
  top = PolynomialString(top)
  bot = PolynomialString(bot)
  return `\\frac{${top}}{${bot}}`
}

function makeErrorBounds(name, cur, next) {
  cur = PolynomialString(cur)
  next = PolynomialString(next)
  return `\\left|y-${name}(x)\\right|<\\frac{1}{(${cur})(${next})}`
}

function resetCalc() {
  calculator.getExpressions().forEach(expression_state => {
    calculator.removeExpression(expression_state);
  });
}