package crypt

import (
	"github.com/codecare/gokeeper/internal/application"
	"encoding/hex"
	"fmt"
	"github.com/codecare/gokeeper/internal/passdata"
	"testing"
)

func Test_roundtrip(t *testing.T) {

	theMessage := "ääÄädfwrt34tökelg geheim"
	s := []byte(theMessage)
	encrypted, err := encrypt([]byte("123456"), s)
	if err != nil {
		t.Errorf("encrypt() error = %v", err)
		return
	}

	plain, err := decrypt([]byte("123456"), encrypted)
	if err != nil {
		t.Errorf("decrypt() error = %v", err)
		return
	}

	is := string(plain)
	if !(theMessage == is) {
		t.Errorf("roundtrip() tobe: %v + is: %v", theMessage, is)
	}
}

func Test_decrypt(t *testing.T) {

	var cipherText = "d6e96e1a5d8614c4460826f63dd1258dbb"
	var nonceText = "9a405f2c07b70597974446c2"
	var keySaltText = "7d1c2070b6e9febf33b2542aa9c120cf"

	cipher, err := hex.DecodeString(cipherText)
	if err != nil { panic(err)}

	nonce, err := hex.DecodeString(nonceText)
	if err != nil { panic(err)}

	keySalt, err := hex.DecodeString(keySaltText)
	if err != nil { panic(err)}

	var encrypted = Encrypted { Cipher: cipher, Nonce: nonce, KeySalt: keySalt }

	decrypted, err := decrypt([]byte("test"), encrypted )
	if err != nil { panic(err)}

	fmt.Printf("decrypted %s\n", string(decrypted))
}

func Test_encrypt(t *testing.T) {

	theMessage := "1"
	s := []byte(theMessage)
	encrypted, err := encrypt([]byte("test"), s)
	if err != nil { panic(err)}

	fmt.Printf("encrypted: nonce: %s cipher: %s keysalt: %s \n", hex.EncodeToString(encrypted.Nonce), hex.EncodeToString(encrypted.Cipher), hex.EncodeToString(encrypted.KeySalt))

	plain, err := decrypt([]byte("test"),encrypted)
	if err != nil {
		t.Errorf("decrypt() error = %v", err)
		return
	}

	is := string(plain)
	if !(theMessage == is) {
		t.Errorf("roundtrip() tobe: %v + is: %v", theMessage, is)
	}
}


func Test_decryptFromJava(t *testing.T) {

	theMessage := "{ \"algorithm\": \"PBKDF2:16383:256:SHA512:AES:GCM\", \"nonce\": \"f3e1dc1f00e120a61501d071\", \"keySalt\": \"3ee32b0036cbb12fa5ab2efb22244580\", \"encryptedText\" : \"9f209c1b125f83714cdd8e003d7ca889e52aae815d953aae79c3e666cec32c9753555f0997c183afa919152a7a282336e5aaded5361ed6798c383e403feff44e9af2c6df444acbd5f733a90e8147d48a5d2c6fc32a7b9df2ead03fba09a74d18717c88d00df813618070515d9a89e92b2d6aee10ec8a1740134a5b8dc35b84ded02c51fff4ee8f218ab127cc1a52ec37c8ea2b3b301f232ff29572c6e2d3926e2975cdb2348f8418d5a1beec68bb695908fa85757a4b5f73fefb1da0e504cae0cc64aba79da234277f875e42ed02bf55ec605cf3f95b8cb09c918fff79aad41890551a9bd21689beb0847ffecd4082129ea5b8a4cf36036a5e1380548891ff073c0a8be4ac1f16364fbe9ed1d7483f9bacf7966b236a37766032e8bc5899ec0e82971ae1e0511056c9ae4e227690291994b0a5fe9d7fa3b1a982a2413f14d249ceab15e8a7ba6f7594c498c035f1cf722d74c1fa51a3aef123f42918e8cc55964c7a18c77cdee883b0c5cb04070ceb1955892708f54a23362bab4bcc32c701da7096d3cbb2425a71ac8d9859b3b56fbc704efddee81c54e6ced1c2bc6871a56d72316ea18e77\"}"

	cryptoContainer, err := passdata.CryptoContainerFromJson([]byte(theMessage))
	if err != nil { panic(err) }

	application.Key = []byte("test")
	bytes, err := DecryptFromContainer(cryptoContainer)
	if err != nil { panic(err) }

	fmt.Printf("decrypted: %s\n", string(bytes))
}

func Test_decryptFromJava2(t *testing.T) {

	theMessage := "{\n  \"algorithm\" : \"PBKDF2:16383:256:SHA512:AES:GCM\",\n  \"nonce\" : \"5fd3fc951910fdae99855121\",\n  \"keySalt\" : \"1714408eaacdf62a029f2d2e649d9242\",\n  \"encryptedText\" : \"f1e7862a6279ce68daffc2b92b5f6674fd3798ec86b2ec95f8bf555665c7da0652ef29294b29e07f3b886a78605047a16713a069eeba1a1cb045eda401f9e38f5f398ee48e60f772bf90047479826882fcaec18d79f4e29116a00b763ff73fb1b0cb23c5959726f5de9cf6b012f0a90b9f096ca36d1f3b53307604178bb213054b4a938463b2b7bffafa881891b1f73760b03019d40473f94c9e4a41648ebf9e97ac1a717759ffdcc73680fa8af686038a07974f312529f6b8e1f6a629afec01c629a462ae34ed01c9ae02482c99381c03efc3d2a9742dcde1f2da73857e4b1ad67b502e69c37178aa935d18b1ccf924a310431fe877011455f4fa650f1c9d3a78e971cfff43628a538e7bc3aa9a4249fc84cd835d109caa46b3e2e8161f82ed2ece219ab8597d3ec8df1a564210f3bec42db61d49020fb5aed97310fbae9aafd509b64a08c5f5cb13e0f2983c313cb45e138fc83e63f053cbb585d654b75c79535ab73af35b36d8d1286f9e44994014cbe09effa6a2e14c7e7799d747b290c4526964ee1048ec0849dc062a0929a7e3ec3a40fdca37c0efd9d5f4c520dcfe7ab34f075a\"\n}\n"

	cryptoContainer, err := passdata.CryptoContainerFromJson([]byte(theMessage))
	if err != nil { panic(err) }

	application.Key = []byte("qwer")
	bytes, err := DecryptFromContainer(cryptoContainer)
	if err != nil { panic(err) }

	fmt.Printf("decrypted: %s\n", string(bytes))
}

