package mcfp

var R4S Machiner = r4s{}

type r4s struct {
	BASE
}

func (r4s) Model() string {
	return "R4S"
}

func (r4s) RootDevPath() string {
	return "/dev/mmcblk1p2"
}
