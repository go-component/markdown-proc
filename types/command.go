package types

type Processing interface {
	Process() error
}


type Command struct {
	Output       string
	ImageDirname string
	Filename     string
	Processing   Processing
}
