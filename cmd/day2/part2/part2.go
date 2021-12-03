package part2

type SubmarineControls struct {
	Aim, Horizontal, Depth int
}

func (c SubmarineControls) Forward(x int) {

	c.Aim = x

}
