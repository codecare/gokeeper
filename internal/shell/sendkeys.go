// Package shell provides a pure Go (no cgo, no third-party
// dependencies) way to inject keyboard events on Linux via /dev/uinput.
//
// The character map assumes a German ("de", QWERTZ) keyboard layout is
// active on the target system (i.e. the OS's configured XKB/console
// layout), since uinput sends physical key codes and the resulting
// character depends on that layout — not on this program.
package shell

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var mostGermanChars = "1234567890ß!\"§$%&/()=?\tqwertzuiopü+asdfghjklöä#<yxcvbnm,.-QWERTZUIOPÜ*ASDFGHJKLÖÄ'>YXCVBNM;:_@~|¹²³[]\\{} "

func VerifyImportantCharmapping() error {

	err := SendCharacters(mostGermanChars+"\n", false)
	if err != nil {
		return err
	}

	input := ReadInput()

	if mostGermanChars != input {
		message := fmt.Sprintf("expected: \n\t'%s'\nbut received \n\t'%s'", mostGermanChars, input)
		return errors.New(message)
	}
	fmt.Printf("received correct input: %s\n", input)

	return nil
}

// --- Linux input subsystem constants (from linux/input-event-codes.h) ---

const (
	evSyn = 0x00
	evKey = 0x01

	synReport = 0

	key1          = 2
	key2          = 3
	key3          = 4
	key4          = 5
	key5          = 6
	key6          = 7
	key7          = 8
	key8          = 9
	key9          = 10
	key0          = 11
	keyMinus      = 12 // "ß" key on German layout
	keyEqual      = 13 // "´ `" dead-key on German layout (unused below)
	keyTab        = 15
	keyQ          = 16
	keyW          = 17
	keyE          = 18
	keyR          = 19
	keyT          = 20
	keyY          = 21 // physical "Y" position -> produces 'z' on German layout
	keyU          = 22
	keyI          = 23
	keyO          = 24
	keyP          = 25
	keyLeftBrace  = 26 // "ü" key on German layout
	keyRightBrace = 27 // "+ * ~" key on German layout
	keyEnter      = 28
	keyA          = 30
	keyS          = 31
	keyD          = 32
	keyF          = 33
	keyG          = 34
	keyH          = 35
	keyJ          = 36
	keyK          = 37
	keyL          = 38
	keySemicolon  = 39 // "ö" key on German layout
	keyApostrophe = 40 // "ä" key on German layout
	keyGrave      = 41
	keyLeftShift  = 42
	keyBackslash  = 43 // "# '" key on German layout
	keyZ          = 44 // physical "Z" position -> produces 'y' on German layout
	keyX          = 45
	keyC          = 46
	keyV          = 47
	keyB          = 48
	keyN          = 49
	keyM          = 50
	keyComma      = 51
	keyDot        = 52
	keySlash      = 53 // "- _" key on German layout
	keySpace      = 57
	keyLeftAlt    = 56
	keyRightAlt   = 100 // AltGr
	key102nd      = 86  // extra key next to left shift on ISO keyboards ("< > |" on German layout)
)

// --- ioctl request numbers, computed from the standard Linux _IO macros
// and verified against <linux/uinput.h> ---

const (
	uiSetEvBit   = 0x40045564
	uiSetKeyBit  = 0x40045565
	uiDevCreate  = 0x5501
	uiDevDestroy = 0x5502
	uiDevSetup   = 0x405c5503
)

const uinputMaxNameSize = 80

// inputID mirrors struct input_id from <linux/input.h>.
type inputID struct {
	BusType uint16
	Vendor  uint16
	Product uint16
	Version uint16
}

// uinputSetup mirrors struct uinput_setup from <linux/uinput.h>.
type uinputSetup struct {
	ID           inputID
	Name         [uinputMaxNameSize]byte
	FFEffectsMax uint32
}

// inputEvent mirrors struct input_event from <linux/input.h> on 64-bit
// Linux (amd64/arm64), where struct timeval's two fields are 64-bit.
// (On a 32-bit kernel/userspace this layout differs — Sec/Usec would need
// to be int32 there.)
type inputEvent struct {
	Sec   int64
	Usec  int64
	Type  uint16
	Code  uint16
	Value int32
}

// keySpec describes which key + modifier combination produces a given
// rune on the active (German/QWERTZ) keyboard layout.
type keySpec struct {
	code  uint16
	shift bool
	altgr bool
}

var charMap = buildCharMap()

func buildCharMap() map[rune]keySpec {
	m := map[rune]keySpec{}

	set := func(r rune, code uint16, shift, altgr bool) {
		m[r] = keySpec{code: code, shift: shift, altgr: altgr}
	}

	set('1', key1, false, false)
	set('!', key1, true, false)
	set('¹', key1, false, true)
	set('2', key2, false, false)
	set('"', key2, true, false)
	set('²', key2, false, true)
	set('3', key3, false, false)
	set('§', key3, true, false)
	set('³', key3, false, true)
	set('4', key4, false, false)
	set('$', key4, true, false)
	set('5', key5, false, false)
	set('%', key5, true, false)
	set('6', key6, false, false)
	set('&', key6, true, false)
	set('7', key7, false, false)
	set('/', key7, true, false)
	set('{', key7, false, true)
	set('8', key8, false, false)
	set('(', key8, true, false)
	set('[', key8, false, true)
	set('9', key9, false, false)
	set(')', key9, true, false)
	set(']', key9, false, true)
	set('0', key0, false, false)
	set('=', key0, true, false)
	set('}', key0, false, true)
	set('ß', keyMinus, false, false)
	set('?', keyMinus, true, false)
	set('\\', keyMinus, false, true)

	set('\t', keyTab, false, false)
	set(' ', keySpace, false, false)
	set('\n', keyEnter, false, false)

	// Top letter row
	set('q', keyQ, false, false)
	set('Q', keyQ, true, false)
	set('@', keyQ, false, true)
	set('w', keyW, false, false)
	set('W', keyW, true, false)
	set('e', keyE, false, false)
	set('E', keyE, true, false)
	set('r', keyR, false, false)
	set('R', keyR, true, false)
	set('t', keyT, false, false)
	set('T', keyT, true, false)
	set('z', keyY, false, false) // German: Y-position key produces 'z'
	set('Z', keyY, true, false)
	set('u', keyU, false, false)
	set('U', keyU, true, false)
	set('i', keyI, false, false)
	set('I', keyI, true, false)
	set('o', keyO, false, false)
	set('O', keyO, true, false)
	set('p', keyP, false, false)
	set('P', keyP, true, false)
	set('ü', keyLeftBrace, false, false)
	set('Ü', keyLeftBrace, true, false)
	set('+', keyRightBrace, false, false)
	set('*', keyRightBrace, true, false)
	set('~', keyRightBrace, false, true)

	// Home row
	set('a', keyA, false, false)
	set('A', keyA, true, false)
	set('s', keyS, false, false)
	set('S', keyS, true, false)
	set('d', keyD, false, false)
	set('D', keyD, true, false)
	set('f', keyF, false, false)
	set('F', keyF, true, false)
	set('g', keyG, false, false)
	set('G', keyG, true, false)
	set('h', keyH, false, false)
	set('H', keyH, true, false)
	set('j', keyJ, false, false)
	set('J', keyJ, true, false)
	set('k', keyK, false, false)
	set('K', keyK, true, false)
	set('l', keyL, false, false)
	set('L', keyL, true, false)
	set('ö', keySemicolon, false, false)
	set('Ö', keySemicolon, true, false)
	set('ä', keyApostrophe, false, false)
	set('Ä', keyApostrophe, true, false)
	set('#', keyBackslash, false, false)
	set('\'', keyBackslash, true, false)

	// Bottom row
	set('<', key102nd, false, false)
	set('>', key102nd, true, false)
	set('|', key102nd, false, true)
	set('y', keyZ, false, false) // German: Z-position key produces 'y'
	set('Y', keyZ, true, false)
	set('x', keyX, false, false)
	set('X', keyX, true, false)
	set('c', keyC, false, false)
	set('C', keyC, true, false)
	set('v', keyV, false, false)
	set('V', keyV, true, false)
	set('b', keyB, false, false)
	set('B', keyB, true, false)
	set('n', keyN, false, false)
	set('N', keyN, true, false)
	set('m', keyM, false, false)
	set('M', keyM, true, false)
	set(',', keyComma, false, false)
	set(';', keyComma, true, false)
	set('.', keyDot, false, false)
	set(':', keyDot, true, false)
	set('-', keySlash, false, false)
	set('_', keySlash, true, false)

	return m
}

// device wraps an open /dev/uinput virtual keyboard.
type device struct {
	f *os.File
}

func ioctl(fd uintptr, req uintptr, arg uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, req, arg)
	if errno != 0 {
		return errno
	}
	return nil
}

func newKeyboardDevice() (*device, error) {
	f, err := os.OpenFile("/dev/uinput", os.O_WRONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("open /dev/uinput (needs root, or a udev rule granting access): %w", err)
	}

	fd := f.Fd()

	if err := ioctl(fd, uiSetEvBit, uintptr(evKey)); err != nil {
		f.Close()
		return nil, fmt.Errorf("UI_SET_EVBIT: %w", err)
	}

	allCodes := map[uint16]struct{}{
		keyLeftShift: {},
		keyLeftAlt:   {},
		keyRightAlt:  {},
		keyTab:       {},
	}
	for _, ks := range charMap {
		allCodes[ks.code] = struct{}{}
	}
	for code := range allCodes {
		if err := ioctl(fd, uiSetKeyBit, uintptr(code)); err != nil {
			f.Close()
			return nil, fmt.Errorf("UI_SET_KEYBIT(%d): %w", code, err)
		}
	}

	setup := uinputSetup{
		ID: inputID{BusType: 0x03, Vendor: 0x1234, Product: 0x5678, Version: 1},
	}
	copy(setup.Name[:], "go-uinput-virtual-keyboard")

	if err := ioctl(fd, uiDevSetup, uintptr(unsafe.Pointer(&setup))); err != nil {
		f.Close()
		return nil, fmt.Errorf("UI_DEV_SETUP: %w", err)
	}
	if err := ioctl(fd, uiDevCreate, 0); err != nil {
		f.Close()
		return nil, fmt.Errorf("UI_DEV_CREATE: %w", err)
	}

	time.Sleep(200 * time.Millisecond) // let X11/Wayland/udev notice the new device
	return &device{f: f}, nil
}

func (d *device) close() error {
	_ = ioctl(d.f.Fd(), uiDevDestroy, 0)
	return d.f.Close()
}

func (d *device) writeEvent(evType, code uint16, value int32) error {
	ev := inputEvent{Type: evType, Code: code, Value: value}
	buf := (*[unsafe.Sizeof(inputEvent{})]byte)(unsafe.Pointer(&ev))[:]
	_, err := d.f.Write(buf)
	return err
}

func (d *device) syn() error { return d.writeEvent(evSyn, synReport, 0) }

func (d *device) keyDown(code uint16) error {
	if err := d.writeEvent(evKey, code, 1); err != nil {
		return err
	}
	return d.syn()
}

func (d *device) keyUp(code uint16) error {
	if err := d.writeEvent(evKey, code, 0); err != nil {
		return err
	}
	return d.syn()
}

func (d *device) tapKey(code uint16, shift, altgr bool) error {
	if shift {
		if err := d.keyDown(keyLeftShift); err != nil {
			return err
		}
	}
	if altgr {
		if err := d.keyDown(keyRightAlt); err != nil {
			return err
		}
	}
	if err := d.keyDown(code); err != nil {
		return err
	}
	time.Sleep(5 * time.Millisecond)
	if err := d.keyUp(code); err != nil {
		return err
	}
	if altgr {
		if err := d.keyUp(keyRightAlt); err != nil {
			return err
		}
	}
	if shift {
		if err := d.keyUp(keyLeftShift); err != nil {
			return err
		}
	}
	time.Sleep(5 * time.Millisecond)
	return nil
}

func (d *device) altTab() error {
	if err := d.keyDown(keyLeftAlt); err != nil {
		return err
	}
	if err := d.keyDown(keyTab); err != nil {
		return err
	}
	time.Sleep(20 * time.Millisecond)
	if err := d.keyUp(keyTab); err != nil {
		return err
	}
	if err := d.keyUp(keyLeftAlt); err != nil {
		return err
	}
	time.Sleep(150 * time.Millisecond) // give the WM time to switch focus
	return nil
}

// SendCharacters types out chars on a virtual keyboard created via
// /dev/uinput. If prefixWithAltTab is true, it sends Alt+Tab first (e.g. to
// switch focus to another window before typing). It works under both X11
// and Wayland since events are injected at the kernel input layer, below
// the display server. Requires permission to open /dev/uinput — either
// root, or a udev rule such as:
//
//	KERNEL=="uinput", MODE="0660", GROUP="input"
//
// with your user added to the "input" group.
//
// The character-to-key mapping assumes a German ("de", QWERTZ) keyboard
// layout is active on the target system.
func SendCharacters(chars string, prefixWithAltTab bool) error {
	dev, err := newKeyboardDevice()
	if err != nil {
		return err
	}
	defer dev.close()

	if prefixWithAltTab {
		if err := dev.altTab(); err != nil {
			return fmt.Errorf("alt-tab: %w", err)
		}
	}

	for _, r := range chars {
		ks, ok := charMap[r]
		if !ok {
			return fmt.Errorf("no key mapping for character %q (rune %U)", r, r)
		}
		if err := dev.tapKey(ks.code, ks.shift, ks.altgr); err != nil {
			return fmt.Errorf("sending character %q: %w", r, err)
		}
	}

	return nil
}
