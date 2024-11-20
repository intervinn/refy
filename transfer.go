package refy

import "slices"

type Packet struct {
	Value string
}

func NewPacket(value string) *Packet {
	return &Packet{
		Value: value,
	}
}

type Register struct {
	Packets map[string][]*Packet
}

func NewRegister() *Register {
	return &Register{
		Packets: map[string][]*Packet{},
	}
}

func (r *Register) Get(key string) []*Packet {
	return slices.Clone(r.Packets[key])
}

func (r *Register) Clean(key string) {
	entry := r.Packets[key]
	r.Packets[key] = slices.Delete(entry, 0, len(entry))
}

func (r *Register) Add(key string, packet *Packet) {
	r.Packets[key] = append(r.Packets[key], packet)
}
