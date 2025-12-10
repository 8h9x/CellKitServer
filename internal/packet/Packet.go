package packet

type Packet interface {
	Build() []byte
}
