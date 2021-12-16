package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

const DEBUG = true

type Binary []byte

func (b Binary) String() string {
	s := ""
	for _, d := range b {
		s += fmt.Sprintf("%v", d)
	}
	return s
}

func byteToBin(b byte) Binary {
	var bytes []byte
	for j := 0; j < 8; j++ {
		bytes = append(bytes, b>>(7-j)&1)
	}
	return bytes
}

func hexToBinary(s string) (Binary, error) {
	var binary []byte
	h, err := hex.DecodeString(s)
	if err != nil {
		return binary, err
	}
	for _, b := range h {
		binary = append(binary, byteToBin(b)...)
	}
	return binary, nil
}

func pow(x, n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= x
	}
	return result
}

func (b Binary) ToInt() int {
	n := len(b)
	num := 0
	for i := n - 1; i >= 0; i-- {
		j := n - i - 1
		num += int(b[i]) * pow(2, j)
	}
	return num
}

type Packet struct {
	version int
	typeId  int
	payload interface{}
}

type PacketReader struct {
	message  Binary
	position int
	end      int
}

func NewPacketReader(str string) (PacketReader, error) {
	bin, err := hexToBinary(str)
	if err != nil {
		log.Fatal(err)
	}
	return PacketReader{bin, 0, len(bin)}, nil
}

func (r *PacketReader) ReadPacket() (Packet, error) {

	if DEBUG {
		fmt.Printf("reading packet: %v\n", r.message[r.position:r.end])
	}

	version, typeId, err := r.ReadHeader()
	if err != nil {
		return Packet{}, err
	}
	payload, err := r.ReadPayload(typeId)
	if err != nil {
		return Packet{}, err
	}
	return Packet{version, typeId, payload}, nil
}

func (r *PacketReader) ReadHeader() (int, int, error) {
	if len(r.message) < 6 {
		return 0, 0, fmt.Errorf("invalid packet: %v", r.message[r.position:])
	}
	version := r.message[r.position : r.position+3]
	typeId := r.message[r.position+3 : r.position+6]
	r.position += 6
	if DEBUG {
		fmt.Printf("   version: %d (%v)\n", version.ToInt(), version)
		fmt.Printf("   typeId: %d (%v)\n", typeId.ToInt(), typeId)
	}
	return version.ToInt(), typeId.ToInt(), nil
}

func (r *PacketReader) ReadPayload(typeId int) (interface{}, error) {
	if typeId == 4 {
		return r.ReadLiteralValue()
	} else {
		return r.ReadOperator()
	}
}

func (r *PacketReader) HasNext() bool {
	return r.position < r.end
}

func (r *PacketReader) Next() (byte, error) {
	if r.HasNext() {
		val := r.message[r.position]
		r.position++
		return val, nil
	}
	return 0, fmt.Errorf("end of input")
}

func (r *PacketReader) ReadLiteralValue() (int, error) {
	if r.position+5 > len(r.message) {
		return 0, fmt.Errorf("invalid message: %v", r.message)
	}

	var groups Binary
	for r.HasNext() {
		start := r.position
		r.position++

		group := Binary{}
		for r.position < start+5 {
			group = append(group, r.message[r.position])
			r.position++
		}
		if DEBUG {
			fmt.Printf("      group: %v\n", group)
		}
		groups = append(groups, group...)

		// last group identifier
		if r.message[start] == 0 {
			break
		}
	}
	value := groups.ToInt()
	if DEBUG {
		fmt.Printf("   literal value: %d\n", value)
	}
	return value, nil
}

func (r *PacketReader) ReadInputLength() (byte, int, error) {
	lengthTypeId := r.message[r.position]
	r.position++

	start := r.position
	var bits Binary
	var fieldLength int

	if DEBUG {
		fmt.Printf("   length type ID: %d\n", lengthTypeId)
	}

	if lengthTypeId == 0 {
		// total length in bits
		fieldLength = 15
	} else {
		// number of sub-packets immediately contained
		fieldLength = 11
	}

	for r.HasNext() {
		if r.position >= start+fieldLength {
			break
		}
		bits = append(bits, r.message[r.position])
		r.position++
	}

	if DEBUG {
		fmt.Printf("   length of sub-packets: %d\n", bits.ToInt())
	}

	return lengthTypeId, bits.ToInt(), nil
}

func (r *PacketReader) ReadOperator() (interface{}, error) {
	lengthTypeId, size, err := r.ReadInputLength()
	if err != nil {
		return nil, err
	}
	if lengthTypeId == 0 {
		return r.ReadPacketsUntil(r.position + size)
	} else {
		return r.ReadNPackets(size)
	}
}

func (r *PacketReader) ReadPacketsUntil(end int) ([]Packet, error) {
	var packets []Packet
	for r.HasNext() {
		packet, err := r.ReadPacket()
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)

		// 3 + 3 + 5 = 11 is the smallest literal package
		if r.end-r.position < 11 {
			break
		}
	}
	return packets, nil
}

func (r *PacketReader) ReadNPackets(n int) ([]Packet, error) {
	var packets []Packet
	for len(packets) < n {
		packet, err := r.ReadPacket()
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}
	return packets, nil
}

func parse(str string) (Packet, error) {
	reader, err := NewPacketReader(str)
	if err != nil {
		return Packet{}, err
	}
	return reader.ReadPacket()
}

func main() {

	for _, example := range []string{
		"D2FE28",
		"38006F45291200",
		"EE00D40C823060",
		"8A004A801A8002F478",
		"620080001611562C8802118E34",
		"C0015000016115A2E0802F182340",
		"A0016C880162017C3686B18A3D4780",
	} {
		fmt.Printf("%v\n", example)
		packet, err := parse(example)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" => %v\n", packet)
		fmt.Println()
	}

}
