package hexagon

type cardinal int

const (
  N cardinal = iota
  NE
  NW
)

type Edge struct {
  q, r int
  d cardinal
}
