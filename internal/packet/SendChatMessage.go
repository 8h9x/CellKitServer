package packet

import "image/color"

type ChatMessageAuthor struct {
	Nickname string
	Color    color.Color
}

type SendChatMessage struct {
	Author      ChatMessageAuthor
	TextContent string
}

func (p *SendChatMessage) Build() []byte {
	writer := NewWriter(make([]byte, 5678))

	writer.WriteUint8(0) // flags
	writer.WriteColor(p.Author.Color)
	writer.WriteStringUTF8(p.Author.Nickname)
	writer.WriteStringUTF8(p.TextContent)

	return writer.Buffer()
}
