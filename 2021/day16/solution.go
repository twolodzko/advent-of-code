package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

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

func (p *Packet) VersionNumberSum() int {
	total := p.version
	switch obj := p.payload.(type) {
	case []Packet:
		for _, packet := range obj {
			total += packet.VersionNumberSum()
		}
	default:
	}
	return total
}

func (p *Packet) Value() (int, error) {
	switch p.typeId {
	case 0:
		return p.Sum()
	case 1:
		return p.Prod()
	case 2:
		return p.Min()
	case 3:
		return p.Max()
	case 4:
		switch obj := p.payload.(type) {
		case int:
			return p.payload.(int), nil
		default:
			return 0, fmt.Errorf("invalid type: %T", obj)
		}
	case 5:
		return p.Greater()
	case 6:
		return p.Less()
	case 7:
		return p.Equal()
	default:
		return 0, fmt.Errorf("invalid typeId: %v", p.typeId)
	}
}

func (p *Packet) PayloadToPackets() ([]Packet, error) {
	switch obj := p.payload.(type) {
	case []Packet:
		return obj, nil
	default:
		return nil, fmt.Errorf("invalid type: %T", obj)
	}
}

func (p *Packet) Sum() (int, error) {
	result := 0
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	for _, packet := range packets {
		val, err := packet.Value()
		if err != nil {
			return 0, err
		}
		result += val
	}
	return result, nil
}

func (p *Packet) Prod() (int, error) {
	result := 1
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	for _, packet := range packets {
		val, err := packet.Value()
		if err != nil {
			return 0, err
		}
		result *= val
	}
	return result, nil
}

func (p *Packet) Min() (int, error) {
	result := math.MaxInt
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	for _, packet := range packets {
		val, err := packet.Value()
		if err != nil {
			return 0, err
		}
		if val < result {
			result = val
		}
	}
	return result, nil
}

func (p *Packet) Max() (int, error) {
	result := 0
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	for _, packet := range packets {
		val, err := packet.Value()
		if err != nil {
			return 0, err
		}
		if val > result {
			result = val
		}
	}
	return result, nil
}

func (p *Packet) Greater() (int, error) {
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	if len(packets) != 2 {
		return 0, fmt.Errorf("expected two packets, got %v", packets)
	}
	x, err := packets[0].Value()
	if err != nil {
		return 0, err
	}
	y, err := packets[1].Value()
	if err != nil {
		return 0, err
	}
	if x > y {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (p *Packet) Less() (int, error) {
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	if len(packets) != 2 {
		return 0, fmt.Errorf("expected two packets, got %v", packets)
	}
	x, err := packets[0].Value()
	if err != nil {
		return 0, err
	}
	y, err := packets[1].Value()
	if err != nil {
		return 0, err
	}
	if x < y {
		return 1, nil
	} else {
		return 0, nil
	}
}

func (p *Packet) Equal() (int, error) {
	packets, err := p.PayloadToPackets()
	if err != nil {
		return 0, err
	}
	if len(packets) != 2 {
		return 0, fmt.Errorf("expected two packets, got %v", packets)
	}
	x, err := packets[0].Value()
	if err != nil {
		return 0, err
	}
	y, err := packets[1].Value()
	if err != nil {
		return 0, err
	}
	if x == y {
		return 1, nil
	} else {
		return 0, nil
	}
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
	if r.end-r.position < 6 {
		return 0, 0, fmt.Errorf("packet to short: %v", r.message[r.position:])
	}
	version := r.message[r.position : r.position+3]
	typeId := r.message[r.position+3 : r.position+6]
	r.position += 6
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

func (r *PacketReader) ReadByte() (byte, error) {
	if r.position < r.end {
		val := r.message[r.position]
		r.position++
		return val, nil
	}
	return 0, io.EOF
}

func (r *PacketReader) ReadLiteralValue() (int, error) {
	if r.position+5 > r.end {
		return 0, fmt.Errorf("invalid message: %v", r.message)
	}

	var groups Binary
	for {
		start, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		group := Binary{}
		for i := 0; i < 4; i++ {
			val, err := r.ReadByte()
			if err != nil {
				return 0, err
			}
			group = append(group, val)
		}
		groups = append(groups, group...)

		// last group identifier
		if start == 0 {
			break
		}
	}
	value := groups.ToInt()
	return value, nil
}

func (r *PacketReader) ReadInputLength() (byte, int, error) {
	lengthTypeId, err := r.ReadByte()
	if err != nil {
		return 0, 0, err
	}

	var bits Binary
	var fieldLength int

	if lengthTypeId == 0 {
		// total length in bits
		fieldLength = 15
	} else {
		// number of sub-packets immediately contained
		fieldLength = 11
	}

	for i := 0; i < fieldLength; i++ {
		val, err := r.ReadByte()
		if err != nil {
			return 0, 0, err
		}
		bits = append(bits, val)
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

func (r *PacketReader) FollowOnlyZeros() bool {
	for i := r.position; i < r.end; i++ {
		if r.message[i] != 0 {
			return false
		}
	}
	return true
}

func (r *PacketReader) ReadPacketsUntil(end int) ([]Packet, error) {
	var packets []Packet
	start := r.position
	for r.position < start+end {
		packet, err := r.ReadPacket()
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)

		if r.FollowOnlyZeros() {
			break
		}
	}
	return packets, nil
}

func (r *PacketReader) ReadNPackets(n int) ([]Packet, error) {
	var packets []Packet
	for {
		packet, err := r.ReadPacket()
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)

		if len(packets) == n {
			break
		}
		if r.FollowOnlyZeros() {
			break
		}
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

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	err = scanner.Err()
	return line, err
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
		"C200B40A82",
		"04005AC33890",
		"880086C3E88112",
		"CE00C43D881120",
		"D8005AC2A8F0",
		"F600BC2D8F",
		"9C005AC2F8F0",
		"9C0141080250320F1802104A08",
	} {
		fmt.Printf("%v\n", example)
		packet, err := parse(example)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" => %v\n", packet)
		fmt.Printf(" checksum = %d\n", packet.VersionNumberSum())
		val, err := packet.Value()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" value = %d\n", val)
		fmt.Println()
	}

	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	message, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	packet, err := parse(message)
	if err != nil {
		log.Fatal(err)
	}

	result1 := packet.VersionNumberSum()
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2, err := packet.Value()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Puzzle 2: %v\n", result2)

}
