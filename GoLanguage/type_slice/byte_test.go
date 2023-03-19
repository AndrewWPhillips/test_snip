package __

import (
	"bytes"
	"log"
	"testing"
)

func TestByteSlice(t *testing.T) {
	var fbinary []byte        // nil slice (len 0)
	log.Println(len(fbinary)) // 0
	var tmp = []byte{}        // zero-length slice (not nil)
	log.Println(len(tmp))     // 0

	log.Println(fbinary) // []
	if fbinary == nil {
		log.Println("NIL") // NIL
	}

	fbinary = nil
	log.Println(fbinary) // []
	if fbinary == nil {
		log.Println("NIL") // NIL
	}

	fbinary = make([]byte, 0, 0)
	log.Println(fbinary) // [] but not nil
	if fbinary == nil {
		log.Println("NIL") // not printed
	}

	tmp = []byte{42, 43, 44}
	log.Println(tmp) // [42 43 44]

	fbinary = make([]byte, 3)
	fbinary[0] = 42
	fbinary[1] = 43
	fbinary[2] = 44
	log.Println(fbinary) // [42 43 44]

	// log.Println(fbinary == tmp)  // slice can only be compared to nil
	log.Println(bytes.Equal(fbinary, tmp)) // true

	tmp[2] = 0
	fbinary = tmp
	log.Println(fbinary) // [42 43 0]

	tmp = nil
	fbinary = tmp
	log.Println(fbinary) // [] and nil
	if fbinary == nil {
		log.Println("NIL") // NIL
	}

	fbinary = []byte{}
	log.Println(fbinary) // [] but not nil
	if fbinary == nil {
		log.Println("NIL") // not printed
	}

	var a [3]byte
	a[1] = 99
	// fbinary = a   // Error: cannot use a (type [3]byte) as type []byte
	fbinary = a[:]       // OK
	log.Println(fbinary) // [0 99 0]

	tmp, fbinary = nil, nil
	log.Println(bytes.Equal(fbinary, tmp)) // true
	tmp = []byte{}
	log.Println(bytes.Equal(fbinary, tmp)) // true!! (empty and nil should be different)
	tmp = []byte{1}
	log.Println(bytes.Equal(fbinary, tmp)) // false
}
