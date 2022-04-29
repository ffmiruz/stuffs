package slices

import (
    "runtime"
    "testing"
)

func TestExtractElements(t *testing.T) {
    tests := []struct {
        name     string
        original []BigData
        expected []BigData
    }{
        {
            name: "everything looks good on the surface",
            original: []BigData{
                {
                    data: [99999]int64{1},
                },
                {
                    data: [99999]int64{2},
                },
                {
                    data: [99999]int64{3},
                },
                {
                    data: [99999]int64{4},
                },
                {
                    data: [99999]int64{5},
                },
                {
                    data: [99999]int64{6},
                },
            },
            expected: []BigData{
                {
                    data: [99999]int64{1},
                },
                {
                    data: [99999]int64{2},
                },
                {
                    data: [99999]int64{3},
                },
            },
        },
    }

    for _, test := range tests {
        test := test
        t.Run(test.name, func(t *testing.T) {
            // cutting halfway (lets not think about odd numbers)
            desiredLen := len(test.original) / 2

            result := ExtractElements(test.original, desiredLen)

            if len(result) != desiredLen {
                t.Fail()
            }

            for i := 0; i < desiredLen; i++ {
                if test.expected[i].data != test.original[i].data {
                    t.Fail()
                }
            }
            t.Error("oh no")
            PrintMemUsage(t)
        })
    }
}

var res []BigData

func BenchmarkExtractElements(b *testing.B) {
    original := []BigData{
        {
            data: [99999]int64{1},
        },
        {
            data: [99999]int64{2},
        },
        {
            data: [99999]int64{3},
        },
        {
            data: [99999]int64{4},
        },
    }
    desiredLen := len(original) / 2

    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        res = ExtractElements(original, desiredLen)
    }
}

func PrintMemUsage(t *testing.T) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    // For info on each, see: https://golang.org/pkg/runtime/#MemStats
    t.Logf("Alloc = %v MiB", bToMb(m.Alloc))
    t.Logf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
    t.Logf("\tSys = %v MiB", bToMb(m.Sys))
    t.Logf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
