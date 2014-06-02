package hexagon

// Hexagon (6-sided polygons) using cubic coordinates
type Hex struct {
  x, y, z int
}

// Creates a hexagon cubic coordinates from axial coordinates
func NewHex(r, q int) *Hex {
  return &Hex{ x: q, z: r, y: -q - r }
}

// Converts hexagon cubic coordinates to axial ones
func (h *Hex) AxialCoordinates() (r, q int) {
  r = h.z
  q = h.x
  return
}

// Finds all six adjacent cubic coordinates
func (h *Hex) NeighborCoordinates() (coords [6][3]int) {
  neighbors := [6][3]int{
    {+1, -1,  0}, {+1,  0, -1}, { 0, +1, -1},
    {-1, +1,  0}, {-1,  0, +1}, { 0, -1, +1},
  }
  for i, value := range neighbors {
    coords[i] = [3]int{
      h.x + value[0],
      h.y + value[1],
      h.z + value[2],
    }
  }
  return
}
