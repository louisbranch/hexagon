package hexagon

type side int

const (
  R side = iota
  L
)

type Vertex struct {
  q, r int
  s side
}
