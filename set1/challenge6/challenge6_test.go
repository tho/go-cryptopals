package challenge6

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
	"unicode"

	"github.com/tho/go-cryptopals/set1/challenge5"
)

func Example() {
	const (
		minKeySize = 2
		maxKeySize = 40
	)

	f, err := os.Open("6.txt")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(f)
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(string(b))
	key, _ := FindRepeatingKeyXORKey(ciphertext, minKeySize, maxKeySize)
	plaintext := challenge5.ComputeRepeatingKeyXOR(ciphertext, key)

	fmt.Printf("%s\n", trimSpaceRight(plaintext))
	// Output:
	// I'm back and I'm ringin' the bell
	// A rockin' on the mike while the fly girls yell
	// In ecstasy in the back of me
	// Well that's my DJ Deshay cuttin' all them Z's
	// Hittin' hard and the girlies goin' crazy
	// Vanilla's on the mike, man I'm not lazy.
	//
	// I'm lettin' my drug kick in
	// It controls my mouth and I begin
	// To just let it flow, let my concepts go
	// My posse's to the side yellin', Go Vanilla Go!
	//
	// Smooth 'cause that's the way I will be
	// And if you don't give a damn, then
	// Why you starin' at me
	// So get off 'cause I control the stage
	// There's no dissin' allowed
	// I'm in my own phase
	// The girlies sa y they love me and that is ok
	// And I can dance better than any kid n' play
	//
	// Stage 2 -- Yea the one ya' wanna listen to
	// It's off my head so let the beat play through
	// So I can funk it up and make it sound good
	// 1-2-3 Yo -- Knock on some wood
	// For good luck, I like my rhymes atrocious
	// Supercalafragilisticexpialidocious
	// I'm an effect and that you can bet
	// I can take a fly girl and make her wet.
	//
	// I'm like Samson -- Samson to Delilah
	// There's no denyin', You can try to hang
	// But you'll keep tryin' to get my style
	// Over and over, practice makes perfect
	// But not if you're a loafer.
	//
	// You'll get nowhere, no place, no time, no girls
	// Soon -- Oh my God, homebody, you probably eat
	// Spaghetti with a spoon! Come on and say it!
	//
	// VIP. Vanilla Ice yep, yep, I'm comin' hard like a rhino
	// Intoxicating so you stagger like a wino
	// So punks stop trying and girl stop cryin'
	// Vanilla Ice is sellin' and you people are buyin'
	// 'Cause why the freaks are jockin' like Crazy Glue
	// Movin' and groovin' trying to sing along
	// All through the ghetto groovin' this here song
	// Now you're amazed by the VIP posse.
	//
	// Steppin' so hard like a German Nazi
	// Startled by the bases hittin' ground
	// There's no trippin' on mine, I'm just gettin' down
	// Sparkamatic, I'm hangin' tight like a fanatic
	// You trapped me once and I thought that
	// You might have it
	// So step down and lend me your ear
	// '89 in my time! You, '90 is my year.
	//
	// You're weakenin' fast, YO! and I can tell it
	// Your body's gettin' hot, so, so I can smell it
	// So don't be mad and don't be sad
	// 'Cause the lyrics belong to ICE, You can call me Dad
	// You're pitchin' a fit, so step back and endure
	// Let the witch doctor, Ice, do the dance to cure
	// So come up close and don't be square
	// You wanna battle me -- Anytime, anywhere
	//
	// You thought that I was weak, Boy, you're dead wrong
	// So come on, everybody and sing this song
	//
	// Say -- Play that funky music Say, go white boy, go white boy go
	// play that funky music Go white boy, go white boy, go
	// Lay down and boogie and play that funky music till you die.
	//
	// Play that funky music Come on, Come on, let me hear
	// Play that funky music white boy you say it, say it
	// Play that funky music A little louder now
	// Play that funky music, white boy Come on, Come on, Come on
	// Play that funky music
}

func trimSpaceRight(s []byte) []byte {
	var b bytes.Buffer

	scanner := bufio.NewScanner(bytes.NewReader(s))
	for scanner.Scan() {
		line := bytes.TrimRightFunc(scanner.Bytes(), unicode.IsSpace)
		b.Write(line)
		b.WriteByte('\n')
	}
	// NOTE: ignoring potential errors from input.Err()

	return b.Bytes()
}

func TestComputeHammingDistane(t *testing.T) {
	tests := map[string]struct {
		s1      []byte
		s2      []byte
		want    int
		wantErr bool
	}{
		"example": {
			s1:      []byte("this is a test"),
			s2:      []byte("wokka wokka!!!"),
			want:    37,
			wantErr: false,
		},
		"nil s1": {
			s1:      nil,
			s2:      []byte{},
			want:    0,
			wantErr: false,
		},
		"nil s2": {
			s1:      []byte{},
			s2:      nil,
			want:    0,
			wantErr: false,
		},
		"nil s1 s2": {
			s1:      nil,
			s2:      nil,
			want:    0,
			wantErr: false,
		},
		"blank s1 s2": {
			s1:      []byte{},
			s2:      []byte{},
			want:    0,
			wantErr: false,
		},
		"unequal length buffers": {
			s1:      []byte("abc"),
			s2:      []byte("abcdef"),
			want:    0,
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := computeHammingDistance(tt.s1, tt.s2)
			if (err != nil) != tt.wantErr {
				t.Fatalf("computeHammingDistance(%q, %q) err: %v, want: %v", tt.s1, tt.s2, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("computeHammingDistance(%q, %q) = %d, want: %d", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	tests := map[string]struct {
		m    [][]byte
		want [][]byte
	}{
		"nil m": {
			m:    nil,
			want: nil,
		},
		"blank m": {
			m:    [][]byte{},
			want: [][]byte{},
		},
		"equal cols": {
			m: [][]byte{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
			want: [][]byte{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
				{4, 4, 4},
			},
		},
		"unequal cols": {
			m: [][]byte{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{1, 2, 3},
			},
			want: [][]byte{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
				{4, 4},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := append(tt.m[:0:0], tt.m...)

			got := transpose(m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose(%v) = %v, want: %v", tt.m, got, tt.want)
			}
			if !reflect.DeepEqual(m, tt.m) {
				t.Errorf("transpose(%v) changed input map to %v", tt.m, m)
			}
		})
	}
}

func TestSplitSize(t *testing.T) {
	tests := map[string]struct {
		s    []byte
		size int
		want [][]byte
	}{
		"nil s": {
			s:    nil,
			size: 1,
			want: [][]byte{
				[]byte(nil),
			},
		},
		"blank s": {
			s:    []byte{},
			size: 1,
			want: [][]byte{
				{},
			},
		},
		"empty s": {
			s:    []byte(""),
			size: 1,
			want: [][]byte{
				[]byte(""),
			},
		},
		"size -1": {
			s:    []byte("test"),
			size: -1,
			want: nil,
		},
		"size 0": {
			s:    []byte("test"),
			size: 0,
			want: nil,
		},
		"size 1": {
			s:    []byte("test"),
			size: 1,
			want: [][]byte{
				[]byte("t"),
				[]byte("e"),
				[]byte("s"),
				[]byte("t"),
			},
		},
		"size 2": {
			s:    []byte("test"),
			size: 2,
			want: [][]byte{
				[]byte("te"),
				[]byte("st"),
			},
		},
		"size 3 (not a multiple)": {
			s:    []byte("test"),
			size: 3,
			want: [][]byte{
				[]byte("tes"),
				[]byte("t"),
			},
		},
		"size equal length": {
			s:    []byte("test"),
			size: 4,
			want: [][]byte{
				[]byte("test"),
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := SplitSize(tt.s, tt.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSize(%q, %d) = %q; want %q", tt.s, tt.size, got, tt.want)
			}
		})
	}
}

func TestSplitSizeN(t *testing.T) {
	tests := map[string]struct {
		s    []byte
		size int
		n    int
		want [][]byte
	}{
		"nil s": {
			s:    nil,
			size: 1,
			n:    -1,
			want: [][]byte{
				[]byte(nil),
			},
		},
		"blank s": {
			s:    []byte{},
			size: 1,
			n:    -1,
			want: [][]byte{
				{},
			},
		},
		"empty s": {
			s:    []byte(""),
			size: 1,
			n:    -1,
			want: [][]byte{
				[]byte(""),
			},
		},
		"n -1": {
			s:    []byte("test"),
			size: 3,
			n:    -1,
			want: [][]byte{
				[]byte("tes"),
				[]byte("t"),
			},
		},
		"n 0": {
			s:    []byte("test"),
			size: 1,
			n:    0,
			want: nil,
		},
		"n < num full subslices": {
			s:    []byte("test"),
			size: 1,
			n:    2,
			want: [][]byte{
				[]byte("t"),
				[]byte("est"),
			},
		},
		"n = num full subslices": {
			s:    []byte("test"),
			size: 1,
			n:    4,
			want: [][]byte{
				[]byte("t"),
				[]byte("e"),
				[]byte("s"),
				[]byte("t"),
			},
		},
		"n > num full subslices": {
			s:    []byte("test"),
			size: 1,
			n:    5,
			want: [][]byte{
				[]byte("t"),
				[]byte("e"),
				[]byte("s"),
				[]byte("t"),
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := SplitSizeN(tt.s, tt.size, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSizeN(%q, %d, %d) = %q; want %q", tt.s, tt.size, tt.n, got, tt.want)
			}
		})
	}
}
