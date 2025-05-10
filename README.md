## How to run benchmarks
```
go test -bench=.
```

## Impact of memory address alignment on performance

#### CPU Cache architecture

```mermaid
block-beta
    columns 1
    block
        columns 2
        core1("CPU Core 1") core2("CPU Core 2")
        space:2
        block
            L1i1("L1 code cache") L1d1("L1 data cache")
        end
        block
            L1i2("L1 code cache") L1d2("L1 data cache")
        end
        space:2
        L21("L2 cache") L22("L2 cache")
        space:2
        L3("L3 shared cache"):2
    end
    Mem{{"\n&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Memory&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\n\n"}}:2
    
L3 --> Mem
L21 --> L3
L22 --> L3

L1i1 --> L21
L1d1 --> L21
L1i2 --> L22
L1d2 --> L22

core1 --> L1i1
core1 --> L1d1
core2 --> L1i2
core2 --> L1d2
```
#### 4 GB of memory divided by cache lines/blocks of 64 bytes each

```mermaid
block-beta
    columns 1

    block
        columns 2
        block4["[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]"] addr4("4294967232 - 4294967295")
        space:2
        block3["[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]"] addr3("128 - 181")
        block2["[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]"] addr2("64 - 127")
        block1["[][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][][]"] addr1("0 - 63")
    end
```
