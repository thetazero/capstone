window.onload = () => {
  let loadElem = document.getElementById("loading")
  loadElem.innerText = "loaded"
  loadElem.classList = ""
}

let vSize = 256
function setup() {
  let cnv = createCanvas(vSize, vSize)
  cnv.parent('vector-controls')
  noLoop()
}

function draw() {
  background(24);
  noFill()
  stroke(200)
  let qp = [q[0] + p[0], q[1] + p[1]]
  let qm = [q[0] - p[0], q[1] - p[1]]
  let m = vSize / 2
  let scale = m / Math.max(CartesianSize(qp), CartesianSize(qm)) * 0.9
  console.log(scale, CartesianSize(p))
  circle(m, vSize / 2, 2 * CartesianSize(p) * scale)
  line(m + q[0] * scale, m - q[1] * scale, m + qp[0] * scale, m - qp[1] * scale)
  line(m + q[0] * scale, m - q[1] * scale, m + qm[0] * scale, m - qm[1] * scale)
  fill(200)
  circle(m + q[0] * scale, m - q[1] * scale, 8)
}

function CartesianSize(arr) {
  let size = 0;
  for (let i = 0; i < arr.length; i++) {
    size += arr[i] ** 2
  }
  return size ** .5
}