package sample

import (
	"math/rand"
	"time"

	"github.com/abhishekvaid/grpc/pcbook/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_AZERTY
	case 2:
		return pb.Keyboard_QWERTY
	default:
		return pb.Keyboard_QWERTZ
	}
}

func randomCPUBrand() string {
	return randomFromStrings("AMD", "Intel", "Apple M1")
}

func randomGPUBrand() string {
	return randomFromStrings("Nvidia", "AMD Radeon")
}

func randomCPUName(brand string) string {
	switch brand {
	case "AMD":
		return randomFromStrings("Ryzen 5800", "Ryzen 4800")
	case "Intel":
		return randomFromStrings("10750H", "XEON", "Celeron")
	default:
		return randomFromStrings("M1", "M1 Pro", "M1 Max")
	}
}

func randomLaptopBrand() string {
	return randomFromStrings("Dell", "Apple", "Hp", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Dell":
		return randomFromStrings("Vostro", "XPS", "Inspiron", "Alienware")
	case "Apple":
		return randomFromStrings("Macbook Air", "Macbook Pro")
	case "Lenovo":
		return randomFromStrings("Legion", "Y540", "ThinkPad")
	default:
		return randomFromStrings("Spectre", "DV Series", "Omen")
	}
}

func randomGPUName(brand string) string {
	switch brand {
	case "AMD Radeon":
		return randomFromStrings("6700", "6800", "6900 XT")
	default:
		return randomFromStrings("3060", "3070", "3080", "3090", "1660Ti")
	}
}

func randomScreenResolution() *pb.Screen_Resolution {

	height := randomInt32(1080, 1080*4)
	width := height * 16 / 9

	return &pb.Screen_Resolution{
		Width:  width,
		Height: height,
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if randomBool() {
		return pb.Screen_IPS
	} else {
		return pb.Screen_OLED
	}
}

// func randomMemory *pb.Memory {

// }

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomFromStrings(s ...string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	return s[rand.Intn(n)]
}

func randomInt32(min, max int) uint32 {
	return (uint32)(min + rand.Intn(max-min+1))
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomId() string {
	return uuid.New().String()
}
