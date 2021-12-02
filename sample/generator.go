package sample

import (
	"github.com/abhishekvaid/grpc/pcbook/pb"
	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {

	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numCores := randomInt32(2, 8)
	numThreads := randomInt32(4, 16)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	memory := &pb.Memory{
		Value: uint64(randomInt32(16, 72)),
		Unit:  pb.Memory_MEGABYTE,
	}

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   numCores,
		NumberThreads: numThreads,
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
		Memory:        memory,
	}
	return cpu
}

func NewGPU() *pb.GPU {

	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.5)
	memory := &pb.Memory{
		Value: uint64(randomInt32(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	cpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return cpu
}

func NewRam() *pb.Memory {
	return &pb.Memory{
		Value: uint64(randomInt32(8, 32)),
		Unit:  pb.Memory_GIGABYTE,
	}
}

func NewSSD() *pb.Storage {
	memory := &pb.Memory{
		Value: uint64(randomInt32(128, 2048)),
		Unit:  pb.Memory_GIGABYTE,
	}

	storage := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: memory,
	}

	return storage
}

func NewHDD() *pb.Storage {

	memory := &pb.Memory{
		Value: uint64(randomInt32(1, 8)),
		Unit:  pb.Memory_TERABYTE,
	}

	storage := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: memory,
	}

	return storage
}

func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   float32(randomFloat64(13, 17)),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

func NewLaptop() *pb.Laptop {

	id := randomId()
	brand := randomLaptopBrand()
	name := randomCPUName(brand)

	return &pb.Laptop{
		Id:       id,
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 4.8),
		},
		UpdatedAt: ptypes.TimestampNow(),
	}

}
