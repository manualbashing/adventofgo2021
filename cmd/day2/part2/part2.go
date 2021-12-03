package part2

type SubmarineControls struct {
	Aim, Horizontal, Depth int
}

func (c *SubmarineControls) Down(x int) {

	c.Aim += x

}

func (c *SubmarineControls) Up(x int) {

	c.Aim -= x

}

func (c *SubmarineControls) Forward(x int) {

	c.Horizontal += x
	c.Depth += c.Aim * x

}
