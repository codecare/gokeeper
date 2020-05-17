package shell

import (
	"errors"
	"fmt"
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

// var importantChars = "!#$%&*+,-.0123456789:;<=>?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@"
var mostGermanChars = "1234567890ß!\"§$%&/()=?\tqwertzuiopü+asdfghjklöä#<yxcvbnm,.-QWERTZUIOPÜ*ASDFGHJKLÖÄ'>YXCVBNM;:_@~|¹²³[]\\{} "

var runeMapping = map[rune] keyboardInput {
	'a': {keybd_event.VK_A, keyboardSwitches{false, false}},
	'b': {keybd_event.VK_B, keyboardSwitches{false, false}},
	'c': {keybd_event.VK_C, keyboardSwitches{false, false}},
	'd': {keybd_event.VK_D, keyboardSwitches{false, false}},
	'e': {keybd_event.VK_E, keyboardSwitches{false, false}},
	'f': {keybd_event.VK_F, keyboardSwitches{false, false}},
	'g': {keybd_event.VK_G, keyboardSwitches{false, false}},
	'h': {keybd_event.VK_H, keyboardSwitches{false, false}},
	'i': {keybd_event.VK_I, keyboardSwitches{false, false}},
	'j': {keybd_event.VK_J, keyboardSwitches{false, false}},
	'k': {keybd_event.VK_K, keyboardSwitches{false, false}},
	'l': {keybd_event.VK_L, keyboardSwitches{false, false}},
	'm': {keybd_event.VK_M, keyboardSwitches{false, false}},
	'n': {keybd_event.VK_N, keyboardSwitches{false, false}},
	'o': {keybd_event.VK_O, keyboardSwitches{false, false}},
	'p': {keybd_event.VK_P, keyboardSwitches{false, false}},
	'q': {keybd_event.VK_Q, keyboardSwitches{false, false}},
	'r': {keybd_event.VK_R, keyboardSwitches{false, false}},
	's': {keybd_event.VK_S, keyboardSwitches{false, false}},
	't': {keybd_event.VK_T, keyboardSwitches{false, false}},
	'u': {keybd_event.VK_U, keyboardSwitches{false, false}},
	'v': {keybd_event.VK_V, keyboardSwitches{false, false}},
	'w': {keybd_event.VK_W, keyboardSwitches{false, false}},
	'x': {keybd_event.VK_X, keyboardSwitches{false, false}},
	'y': {keybd_event.VK_Z, keyboardSwitches{false, false}},
	'z': {keybd_event.VK_Y, keyboardSwitches{false, false}},

	'A': {keybd_event.VK_A, keyboardSwitches{true, false}},
	'B': {keybd_event.VK_B, keyboardSwitches{true, false}},
	'C': {keybd_event.VK_C, keyboardSwitches{true, false}},
	'D': {keybd_event.VK_D, keyboardSwitches{true, false}},
	'E': {keybd_event.VK_E, keyboardSwitches{true, false}},
	'F': {keybd_event.VK_F, keyboardSwitches{true, false}},
	'G': {keybd_event.VK_G, keyboardSwitches{true, false}},
	'H': {keybd_event.VK_H, keyboardSwitches{true, false}},
	'I': {keybd_event.VK_I, keyboardSwitches{true, false}},
	'J': {keybd_event.VK_J, keyboardSwitches{true, false}},
	'K': {keybd_event.VK_K, keyboardSwitches{true, false}},
	'L': {keybd_event.VK_L, keyboardSwitches{true, false}},
	'M': {keybd_event.VK_M, keyboardSwitches{true, false}},
	'N': {keybd_event.VK_N, keyboardSwitches{true, false}},
	'O': {keybd_event.VK_O, keyboardSwitches{true, false}},
	'P': {keybd_event.VK_P, keyboardSwitches{true, false}},
	'Q': {keybd_event.VK_Q, keyboardSwitches{true, false}},
	'R': {keybd_event.VK_R, keyboardSwitches{true, false}},
	'S': {keybd_event.VK_S, keyboardSwitches{true, false}},
	'T': {keybd_event.VK_T, keyboardSwitches{true, false}},
	'U': {keybd_event.VK_U, keyboardSwitches{true, false}},
	'V': {keybd_event.VK_V, keyboardSwitches{true, false}},
	'W': {keybd_event.VK_W, keyboardSwitches{true, false}},
	'X': {keybd_event.VK_X, keyboardSwitches{true, false}},
	'Y': {keybd_event.VK_Z, keyboardSwitches{true, false}},
	'Z': {keybd_event.VK_Y, keyboardSwitches{true, false}},

	'!': {keybd_event.VK_1, keyboardSwitches{true, false}},
	'#': {keybd_event.VK_SP8, keyboardSwitches{false, false}},
	'$': {keybd_event.VK_4, keyboardSwitches{true, false}},
	'%': {keybd_event.VK_5, keyboardSwitches{true, false}},
	'&': {keybd_event.VK_6, keyboardSwitches{true, false}},
	'*': {keybd_event.VK_SP5, keyboardSwitches{true, false}},
	'+': {keybd_event.VK_SP5, keyboardSwitches{false, false}},
	',': {keybd_event.VK_SP9, keyboardSwitches{false, false}},
	'-': {keybd_event.VK_SP11, keyboardSwitches{false, false}},
	'.': {keybd_event.VK_SP10, keyboardSwitches{false, false}},

	'0': {keybd_event.VK_0, keyboardSwitches{false, false}},
	'1': {keybd_event.VK_1, keyboardSwitches{false, false}},
	'2': {keybd_event.VK_2, keyboardSwitches{false, false}},
	'3': {keybd_event.VK_3, keyboardSwitches{false, false}},
	'4': {keybd_event.VK_4, keyboardSwitches{false, false}},
	'5': {keybd_event.VK_5, keyboardSwitches{false, false}},
	'6': {keybd_event.VK_6, keyboardSwitches{false, false}},
	'7': {keybd_event.VK_7, keyboardSwitches{false, false}},
	'8': {keybd_event.VK_8, keyboardSwitches{false, false}},
	'9': {keybd_event.VK_9, keyboardSwitches{false, false}},

	':': {keybd_event.VK_SP10, keyboardSwitches{true, false}},
	';': {keybd_event.VK_SP9, keyboardSwitches{true, false}},
	'<': {keybd_event.VK_SP12, keyboardSwitches{false, false}},
	'=': {keybd_event.VK_0, keyboardSwitches{true, false}},
	'>': {keybd_event.VK_SP12, keyboardSwitches{true, false}},
	'?': {keybd_event.VK_SP2, keyboardSwitches{true, false}},
	'ß': {keybd_event.VK_SP2, keyboardSwitches{false, false}},
	'@': {keybd_event.VK_Q, keyboardSwitches{false, true}},
	'"': {keybd_event.VK_2, keyboardSwitches{true, false}},
	'§': {keybd_event.VK_3, keyboardSwitches{true, false}},
	'/': {keybd_event.VK_7, keyboardSwitches{true, false}},

	'(': {keybd_event.VK_8, keyboardSwitches{true, false}},
	')': {keybd_event.VK_9, keyboardSwitches{true, false}},

	'[': {keybd_event.VK_8, keyboardSwitches{false, true}},
	']': {keybd_event.VK_9, keyboardSwitches{false, true}},

	'{': {keybd_event.VK_7, keyboardSwitches{false, true}},
	'}': {keybd_event.VK_0, keyboardSwitches{false, true}},

	'\n': {keybd_event.VK_ENTER, keyboardSwitches{false, false}},
	'\t': {keybd_event.VK_TAB, keyboardSwitches{false, false}},

	'ü': {keybd_event.VK_SP4, keyboardSwitches{false, false}},
	'ö': {keybd_event.VK_SP6, keyboardSwitches{false, false}},
	'ä': {keybd_event.VK_SP7, keyboardSwitches{false, false}},

	'Ü': {keybd_event.VK_SP4, keyboardSwitches{true, false}},
	'Ö': {keybd_event.VK_SP6, keyboardSwitches{true, false}},
	'Ä': {keybd_event.VK_SP7, keyboardSwitches{true, false}},

	'\'': {keybd_event.VK_SP8, keyboardSwitches{true, false}},
	'_':  {keybd_event.VK_SP11, keyboardSwitches{true, false}},

	'~': {keybd_event.VK_SP5, keyboardSwitches{false, true}},

	'|':  {keybd_event.VK_SP12, keyboardSwitches{false, true}},
	'\\': {keybd_event.VK_SP2, keyboardSwitches{false, true}},
	' ':  {keybd_event.VK_SPACE, keyboardSwitches{false, false}},

	'¹': {keybd_event.VK_1, keyboardSwitches{false, true}},
	'²': {keybd_event.VK_2, keyboardSwitches{false, true}},
	'³': {keybd_event.VK_3, keyboardSwitches{false, true}},
}

func SendCharacters(chars string, prefixWithAltTab bool) error {

	keyboardInputs := make([]keyboardInput, len(chars))

	for _, singleRune := range chars {
		keyboardInput, found := runeMapping[singleRune]
		if found == false {
			// return errors.New(fmt.Sprintf("0no keyboard mapping for %s %d", string(singleRune), singleRune))
			fmt.Printf("no keyboard mapping for %s %d\n", string(singleRune), singleRune)
		} else {
			keyboardInputs = append(keyboardInputs, keyboardInput)
		}
	}

	kb, err := keybd_event.NewKeyBonding()
	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(1000 * time.Millisecond)
	}

	if err != nil { return err }

	if prefixWithAltTab {
		err = sendAltTab(kb)
		if err != nil { return err }
	}

	for _, singleInput := range keyboardInputs {
		kb.Clear()
		kb.HasSHIFT(singleInput.keyboardSwitches.hasShift)
		kb.HasALTGR(singleInput.keyboardSwitches.hasAltGr)
		kb.AddKey(singleInput.keycode)
		err = kb.Launching()
		time.Sleep(2 * time.Millisecond)
		if err != nil { return err }
	}

	time.Sleep(1000 * time.Millisecond)

	return nil
}

func sendAltTab(kb keybd_event.KeyBonding) error {
	kb.Clear()

	kb.HasALT(true)
	kb.AddKey(keybd_event.VK_TAB)

	err := kb.Launching()
	time.Sleep(500 * time.Millisecond)
	return err
}

func VerifyImportantCharmapping() error {

	err := SendCharacters(mostGermanChars + "\n", false)
	if err != nil { return err }

	input := ReadInput()

	if mostGermanChars != input {
		message := fmt.Sprintf("expected: \n\t'%s'\nbut received \n\t'%s'", mostGermanChars, input)
		return errors.New(message)
	}
	fmt.Printf("received correct input: %s\n", input)

	return nil
}

type keyboardInput struct {
	keycode int
	keyboardSwitches
}

type keyboardSwitches struct {
	hasShift bool
	hasAltGr bool
}



