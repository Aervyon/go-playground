package utils

import (
	"runtime"

	"github.com/alexedwards/argon2id"
	sigar "github.com/cloudfoundry/gosigar"
)

var minCPUs = 4

// We want to leave a thread left over for the OS to use.
// Hopefully this will prevent the OS from killing the process.
func hashingCPUs() uint8 {
	if runtime.NumCPU() > minCPUs {
		return uint8(runtime.NumCPU() - 1)
	}
	return uint8(runtime.NumCPU())
}

func determineMemory() uint32 {
	mem := sigar.Mem{}
	mem.Get()
	// return the total memory in KB
	totalGB := mem.Total / 1024 / 1024 / 1000

	const base = 64 * 1024 // 64MB in KB
	if totalGB <= 1 {
		return uint32(base)
	}
	return uint32(totalGB * base / 2) // 32MB per GB
}

// Baseline of 3 iterations
// If the system has less than 20 CPUs and 1/3 of that number is less than 6, use 6 iterations.
// Otherwise, use 1/3 of the CPUs.
func determineIterations() uint32 {
	cpus := hashingCPUs()
	if cpus <= 1 {
		return 3
	}

	if cpus <= 6 {
		return uint32(cpus)
	}

	if cpus <= 20 && cpus/3 <= 6 {
		return 6
	}

	return uint32(cpus / 3)
}

// HashParams is my custom set parameters for hashing passwords.
// Based on guidelines from OWASP, and the argon2id Node package.
// Automatically calculates memory, iterations and cpus based on the system.
var HashParams = &argon2id.Params{
	Memory:      determineMemory(),
	Iterations:  determineIterations(),
	SaltLength:  16,
	KeyLength:   32,
	Parallelism: hashingCPUs(),
}
