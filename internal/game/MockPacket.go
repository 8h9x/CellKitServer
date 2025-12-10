package game

type MockPacket struct {
	Raw []byte
}

func (mp *MockPacket) Build() []byte {
	return mp.Raw
}
