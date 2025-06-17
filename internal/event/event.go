package event

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/osm/rvspub/internal/buffer"
	"github.com/osm/rvspub/internal/charset"
)

const (
	eyesModelChecksum   = 6967
	playerModelChecksum = 13845
	clAccelTreshhold    = 100
)

var signatureV2 = []byte{
	0x45, 0xe4, 0xb0, 0x25, 0x3a, 0x4f, 0xf5, 0x08,
	0x2a, 0x04, 0x21, 0xb3, 0x92, 0x91, 0xa5, 0xf6,
}

type Event struct {
	Version           string        `json:"version"`
	ExternalIPAddress net.IP        `json:"external_ip_address"`
	InternalIPAddress net.IP        `json:"internal_ip_address"`
	SessionID         string        `json:"session_id"`
	Hostname          string        `json:"hostname"`
	Signature         string        `json:"signature"`
	Timestamp         time.Time     `json:"timestamp"`
	Duration          time.Time     `json:"duration"`
	Elapsed           time.Duration `json:"elapsed"`
	ServerAddress     string        `json:"server_address"`
	ServerHostname    string        `json:"server_hostname"`
	UserID            uint16        `json:"user_id"`
	Name              string        `json:"name"`
	EyesModel         uint16        `json:"eyes_model"`
	PlayerModel       uint16        `json:"player_model"`
	Frags             int16         `json:"frags"`
	CLAccel           uint16        `json:"cl_accel"`
	Cheats            string        `json:"cheats"`
	Spectator         bool          `json:"is_spectator"`
}

func FromBuffer(b *buffer.Buffer) (Event, error) {
	var entry Event
	var err error

	if err = entry.parseVersion(b); err != nil {
		return entry, err
	}

	if err = entry.parseExternalIPAddress(b); err != nil {
		return entry, err
	}

	// The external IP address is also stored as string, but we can skip
	// that since the same value is parsed above.
	if err = b.SkipBytes(24); err != nil {
		return entry, readError("skip ip address string", b.Offset(), err)
	}

	if err = entry.parseInternalIPAddress(b); err != nil {
		return entry, err
	}

	if err = entry.parseSessionID(b); err != nil {
		return entry, err
	}

	if err = entry.parseHostname(b); err != nil {
		return entry, err
	}

	if err = entry.parseSignature(b); err != nil {
		return entry, err
	}

	if err = entry.parseTimestamp(b); err != nil {
		return entry, err
	}

	if err = entry.parseDuration(b); err != nil {
		return entry, err
	}

	entry.Elapsed = entry.Duration.Sub(entry.Timestamp)

	// TODO: Figure out what these bytes are.
	if err = b.SkipBytes(6); err != nil {
		return entry, readError("skip 6 unknown bytes", b.Offset(), err)
	}

	if err = entry.parseServerAddress(b); err != nil {
		return entry, err
	}

	if err = entry.parseServerHostname(b); err != nil {
		return entry, err
	}

	if err = entry.parseUserID(b); err != nil {
		return entry, err
	}

	if err = entry.parseName(b); err != nil {
		return entry, err
	}

	if err = entry.parseEyesModel(b); err != nil {
		return entry, err
	}

	if err = entry.parsePlayerModel(b); err != nil {
		return entry, err
	}

	if err = entry.parseFrags(b); err != nil {
		return entry, err
	}

	if err = entry.parseCLAccel(b); err != nil {
		return entry, err
	}

	// TODO: Figure out what these bytes are.
	if err = b.SkipBytes(2); err != nil {
		return entry, readError("skip 2 unknown bytes", b.Offset(), err)
	}

	if err = entry.parseCheats(b); err != nil {
		return entry, err
	}

	return entry, nil
}

func (e *Event) parseVersion(b *buffer.Buffer) error {
	val, err := b.ReadBytes(2)
	if err != nil {
		return readError("parse version", b.Offset(), err)
	}

	e.Version = fmt.Sprintf("%d.%d", val[0], val[1])
	return nil
}

func (e *Event) parseExternalIPAddress(b *buffer.Buffer) error {
	val, err := b.ReadBytes(4)
	if err != nil {
		return readError("parse external ip address", b.Offset(), err)
	}

	e.ExternalIPAddress = net.IPv4(val[3], val[2], val[1], val[0])
	return nil
}

func (e *Event) parseInternalIPAddress(b *buffer.Buffer) error {
	val, err := b.ReadBytes(4)
	if err != nil {
		return readError("parse internal ip address", b.Offset(), err)
	}

	e.InternalIPAddress = net.IPv4(val[0], val[1], val[2], val[3])
	return nil
}

func (e *Event) parseSessionID(b *buffer.Buffer) error {
	val, err := b.ReadBytes(4)
	if err != nil {
		return readError("parse session ID", b.Offset(), err)
	}

	e.SessionID = hex.EncodeToString([]byte{val[3], val[2], val[1], val[0]})
	return nil
}

func (e *Event) parseHostname(b *buffer.Buffer) error {
	var hostnameSize int
	if e.Version == "2.1" {
		hostnameSize = 80
	} else {
		hostnameSize = 64
	}

	val, err := b.ReadBytes(hostnameSize)
	if err != nil {
		return readError("parse hostname", b.Offset(), err)
	}

	e.Hostname = charset.Parse(val)
	return nil
}

func (e *Event) parseSignature(b *buffer.Buffer) error {
	if e.Version == "2.1" {
		e.Signature = hex.EncodeToString(signatureV2)
		return nil
	}

	val, err := b.ReadBytes(16)
	if err != nil {
		return readError("parse signature", b.Offset(), err)
	}

	e.Signature = hex.EncodeToString(val)
	return nil
}

func (e *Event) parseTimestamp(b *buffer.Buffer) error {
	val, err := b.ReadBytes(4)
	if err != nil {
		return readError("parse timestamp", b.Offset(), err)
	}

	e.Timestamp = time.Unix(int64(binary.LittleEndian.Uint32(val)), 0).UTC()
	return nil
}

func (e *Event) parseDuration(b *buffer.Buffer) error {
	val, err := b.ReadBytes(4)
	if err != nil {
		return readError("parse duration", b.Offset(), err)
	}

	e.Duration = time.Unix(int64(binary.LittleEndian.Uint32(val)), 0).UTC()
	return nil
}

func (e *Event) parseServerAddress(b *buffer.Buffer) error {
	val, err := b.ReadBytes(24)
	if err != nil {
		return readError("parse server address", b.Offset(), err)
	}

	e.ServerAddress = charset.Parse(val)
	return nil
}

func (e *Event) parseServerHostname(b *buffer.Buffer) error {
	val, err := b.ReadBytes(64)
	if err != nil {
		return readError("parse server hostname", b.Offset(), err)
	}

	e.ServerHostname = charset.Parse(val)
	return nil
}

func (e *Event) parseUserID(b *buffer.Buffer) error {
	val, err := b.ReadBytes(2)
	if err != nil {
		return readError("parse user ID", b.Offset(), err)
	}

	e.UserID = binary.LittleEndian.Uint16(val)
	return nil
}

func (e *Event) parseName(b *buffer.Buffer) error {
	val, err := b.ReadBytes(32)
	if err != nil {
		return readError("parse user name", b.Offset(), err)
	}

	e.Name = charset.Parse(val)
	return nil
}

func (e *Event) parseEyesModel(b *buffer.Buffer) error {
	val, err := b.ReadBytes(2)
	if err != nil {
		return readError("parse eyes model", b.Offset(), err)
	}

	e.EyesModel = binary.LittleEndian.Uint16(val)
	return nil
}

func (e *Event) parsePlayerModel(b *buffer.Buffer) error {
	val, err := b.ReadBytes(2)
	if err != nil {
		return readError("parse player model", b.Offset(), err)
	}

	e.PlayerModel = binary.LittleEndian.Uint16(val)
	return nil
}

func (e *Event) parseFrags(b *buffer.Buffer) error {
	val, err := b.ReadBytes(2)
	if err != nil {
		return readError("parse frags", b.Offset(), err)
	}

	e.Frags = int16(binary.LittleEndian.Uint16(val))
	return nil
}

func (e *Event) parseCLAccel(b *buffer.Buffer) error {
	val, err := b.ReadBytes(2)
	if err != nil {
		return readError("parse cl_accel", b.Offset(), err)
	}

	e.CLAccel = binary.LittleEndian.Uint16(val)
	return nil
}

func (e *Event) parseCheats(b *buffer.Buffer) error {
	val, err := b.ReadBytes(4)
	if err != nil {
		return readError("parse cheats", b.Offset(), err)
	}

	var cheats []string

	if e.EyesModel != eyesModelChecksum {
		cheats = append(cheats, "E!")
	}

	if e.PlayerModel != playerModelChecksum {
		cheats = append(cheats, "P?")
	}

	if e.CLAccel > clAccelTreshhold {
		cheats = append(cheats, "S!")
	}

	// TODO: Determine the meaning of all abbreviations.
	c := binary.LittleEndian.Uint32(val)
	switch c {
	case 0x00:
	case 0x02:
		cheats = append(cheats, "R!")
	case 0x04:
		cheats = append(cheats, "A!")
	case 0x06:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "A!")
	case 0x08:
		cheats = append(cheats, "B!")
	case 0x0a:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "B!")
	case 0x0c:
		cheats = append(cheats, "A!")
		cheats = append(cheats, "B!")
	case 0x0e:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "A!")
		cheats = append(cheats, "B!")
	case 0x10:
		cheats = append(cheats, "F!")
	case 0x12:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "F!")
	case 0x14:
		cheats = append(cheats, "A!")
		cheats = append(cheats, "F!")
	case 0x16:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "A!")
		cheats = append(cheats, "F!")
	case 0x18:
		cheats = append(cheats, "B!")
		cheats = append(cheats, "F!")
	case 0x1a:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "B!")
		cheats = append(cheats, "F!")
	case 0x1c:
		cheats = append(cheats, "A!")
		cheats = append(cheats, "B!")
		cheats = append(cheats, "F!")
	case 0x1e:
		cheats = append(cheats, "R!")
		cheats = append(cheats, "A!")
		cheats = append(cheats, "B!")
		cheats = append(cheats, "F!")
	case 0x01, 0x03, 0x05, 0x07, 0x09, 0x0b, 0x0d, 0x0f,
		0x11, 0x13, 0x15, 0x17, 0x19, 0x1b, 0x1d, 0x1f:
		e.Spectator = true
	default:
		return fmt.Errorf("unknown cheat %q", c)
	}

	e.Cheats = strings.Join(cheats, " ")
	return nil
}

func readError(action string, offset int, err error) error {
	return fmt.Errorf("failed to %s, offset: %d, err: %v", action, offset, err)
}
