# SIMD

## X86 History

MMX 1997 Pentium 2     8 x 64-bit registers (=x87), integer 8/16/32/64 ops
SSE 1999 P3            8 x 128-bit register (XMM0-7), mostly FP instructions
SSE2 2000 P4/Athlon64  
SSE3 2004 
SSSE3 2006 Core 2
SSE4a 2007 AMD K10
SSE4.1 2007
SSE4.2 2008 Core i7
AVX 2011 (Adv. Vec. Ext.) 16 x 256-bit registers (YMM0-15), 3 op (src1, src2, dst)
AVX2 2013 
AVX512 2016 Core i9, Zen4 32 x 512-bit registes
AMX 2023
AVX10.1
AVX10.2

## Types 128bit

float32x4
float64x2

int8x16
int16x8
int32x4
int64x2

uint8x16
uint16x8
uint32x4
uint64x2

mask8x16
mask16x8
mask32x4
mask64x2

## Types 256bit

float32x8
float64x4

int8x32
int16x16
int32x8
int64x4

uint8x32
uint16x16
uint32x8
uint64x4

mask8x32
mask16x16
mask32x8
mask64x4

## Types 512bit (AVX512)

float32x16
float64x8

int8x64
int16x32
int32x16
int64x8

uint8x64
uint16x32
uint32x16
uint64x8

mask8x64
mask16x32
mask32x16
mask64x8

## Type "Cast"

As Int/Uint 8x16 16x8 32x4 64x2 - from any other 8x16 16x8 32x4 64x2 incl. Float32x4/Float64x2
- eg: func (x Int16x8) AsInt8x16() Int8x16
As Float 32x4 64x2 - from any other 8x16 16x8 32x4 64x2

As Int/Uint/Float 8x32 16x16 32x8 64x4 - from a different 8x32 16x16 32x8 64x4 
As Int/Uint/Float 8x64 16x32 32x16 64x8 - from a different 8x64 16x32 32x16 64x8

## Crypto

AESDecryptLastRound 16/32/64
AESDecryptOneRound 16/32/64
AESEncryptLastRound 16/32/64
AESEncryptOneRound 16/32/64
AESInvMixColumns
AESRoundKeyGenAssist

SHA1FourRounds, SHA1Message1, SHA1Message2, SHA1NextE  Uint32x4
SHA256TwoRounds, SHA256Message1, SHA256Message2        Uint32x4
xxx where is SHA256Next?

## Org

Broadcast 1To2/1To4/1To8/1To16/1To32/1To64  Float/Int/Uint

GetElem           Float/Int/Uint 8x16/16x8/32x4/64x2
GetHi, GetLo      Float/Int/Uint 8x32/8x64/16x16/16x32/32x8/32x16/64x4/64x8
SetElem           Float/Int/Uint 8x16/16x8/32x4/64x2
SetHi, SetLo      Float/Int/Uint 8x32/8x64/16x16/16x32/32x8/32x16/64x4/64x8

Select128FromPair Float/Int/Uint 8x32/16x16/32x8/64x4

ConcatPermute Float/Int/Uint
ConcatShiftBytesRight  Uint 8x16
ConcatShiftBytesRightGrouped Uint 8x32/8x64

Permute              Float/Int/Uint
PermuteOrZero        Int/Uint  8x16
PermuteOrZeroGrouped Int/Uint  8x32/8x64

InterleaveHi        Int/Uint 16x8/32x4/64x2
InterleaveHiGrouped Int/Uint 16x16/16x32 32x8/32x16 64x4/64x8
InterleaveLo        Int/Uint 16x8/32x4/64x2
InterleaveLoGrouped Int/Uint 16x16/16x32 32x8/32x16 64x4/64x8

Compress Float/Int/Uint
Expand   Float/Int/Uint


## Bits

And Int/Uint
AndNot Int/Uint
Or     Int/Uint 8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
Xor     Int/Uint 8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8

RotateLeft     Int/Uint 32x4/32x8/32x16 64x2/64x4/64x8 - shifts in register
RotateRight    Int/Uint 32x4/32x8/32x16 64x2/64x4/64x8
RotateAllLeft  Int/Uint 32x4/32x8/32x16 64x2/64x4/64x8 - all same shift
RotateAllRight Int/Uint 32x4/32x8/32x16 64x2/64x4/64x8

ShiftLeft        Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftLeftConcat  Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftRight       Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftRightConcat Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftAllLeft        Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftAllLeftConcat  Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftAllRight       Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
ShiftAllRightConcat Int/Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8

LeadingZeros Int/Uint 32x4/32x8/32x16 64x2/64x4/64x8
OnesCount    Int/Uint 8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8


## Float Maths

ConvertToFloat32  Float64/Int32/Int64/Uint32/Uint64
ConvertToFloat64  Float32/Int32/Int64/Uint32/Uint64

Equal        Float 32x4/32x8/32x16 64x2/64x4/64x8
NotEqual     Float 32x4/32x8/32x16 64x2/64x4/64x8
Greater      Float 32x4/32x8/32x16 64x2/64x4/64x8
GreaterEqual Float 32x4/32x8/32x16 64x2/64x4/64x8
Less         Float 32x4/32x8/32x16 64x2/64x4/64x8
LessEqual    Float 32x4/32x8/32x16 64x2/64x4/64x8

Add    Float 32x4/32x8/32x16 64x2/64x4/64x8
AddSub Float
AddPairs, AddPairsGrouped Float
Sub    Float 32x4/32x8/32x16 64x2/64x4/64x8
SubPairs, SubPairsGrouped Float

Div Float       32x4/32x8/32x16 64x2/64x4/64x8
Mul Float       32x4/32x8/32x16 64x2/64x4/64x8
MulAdd Float    32x4/32x8/32x16 64x2/64x4/64x8
MulAddSub Float 32x4/32x8/32x16 64x2/64x4/64x8
MulSubAdd Float 32x4/32x8/32x16 64x2/64x4/64x8

Ceil                     Float 32x4/32x8 64x2/64x4
CeilScaled               Float 32x4/32x8/32x16 64x2/64x4/64x8
CeilScaledResidue        Float 32x4/32x8/32x16 64x2/64x4/64x8
Floor                    Float 32x4/32x8 64x2/64x4
FloorScaled              Float 32x4/32x8/32x16 64x2/64x4/64x8
FloorScaledResidue       Float 32x4/32x8/32x16 64x2/64x4/64x8
RoundToEven              Float 32x4/32x8 64x2/64x4
RoundToEvenScaled        Float 32x4/32x8/32x16 64x2/64x4/64x8
RoundToEvenScaledResidue Float 32x4/32x8/32x16 64x2/64x4/64x8
Trunc                    Float 32x4/32x8 64x2/64x4
TruncScaled              Float 32x4/32x8/32x16 64x2/64x4/64x8
TruncScaledResidue       Float 32x4/32x8/32x16 64x2/64x4/64x8
Scale                    Float 32x4/32x8/32x16 64x2/64x4/64x8

Max/Min        Float 32x4/32x8/32x16 64x2/64x4/64x8
Sqrt           Float 32x4/32x8/32x16 64x2/64x4/64x8
Reciprocal     Float 32x4/32x8/32x16 64x2/64x4/64x8
ReciprocalSqrt Float 32x4/32x8/32x16 64x2/64x4/64x8


## Int Maths

ConvertToInt32  Float32/Float64
ConvertToInt64  Float32/Float64
ConvertToUint32 Float32/Float64
ConvertToUint64 Float32/Float64

TruncateToInt8  Int 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
TruncateToInt16 Int 32x4/32x8/32x16 64x2/64x4/64x8
TruncateToInt32 Int 64x2/64x4/64x8
TruncateToUint8  Uint 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
TruncateToUint16 Uint 32x4/32x8/32x16 64x2/64x4/64x8
TruncateToUint32 Uint 64x2/64x4/64x8

Abs     Int  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
CopySign Int 8x16/8x32 16x8/16x16 32x4/32x8
SumAbsDiff Uint8x16 Uint8x32 Uint8x64

Average Uint 8x16/8x32/8x64 16x8/16x16/16x32

Add Int/Uint  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
AddPairs, AddPairsGrouped Int/Uint
Sub Int/Uint  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
SubPairs, SubPairsGrouped Int/Uint

AddSaturated Int/Uint  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
AddPairsSaturated, AddPairsSaturatedGrouped Int/Uint
SubSaturated Int/Uint  8x16/8x32/8x64 16x8/16x16/16x32 xxx 32??
SubPairsSaturated, SubPairsSaturatedGrouped Int16

xxx why no int Div???
Mul          Int/Uint  16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
xxx why no int8/uint8???
MulEvenWiden Int/Uint 32x4/32x8
MulHigh      Int/Uint 16x8/16x16/16x32

Max/Min     Int/Uint  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8

Equal       Int/Uint  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
NotEqual    Int/Uint  8x64  16x32  32x16  64x8
xxx why 12 versions of greater for int but only 4 for uint and 4 for less int
Greater      Int  8x16/8x32/8x64 16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
Greater      Uint 8x64/16x32/32x16/64x8
GreaterEqual Int/Uint  8x64/16x32/32x16/64x8
Less      Int/Uint  8x64/16x32/32x16/64x8
LessEqual Int/Uint  8x64/16x32/32x16/64x8

ExtendLo2To Int64/Uint64
ExtendLo4To Int/Uint 32/64
ExtendLo8To Int/Uint 16/32/64

ExtendToInt16 Int8x16/Int8x32 - sign extend
ExtendToInt32 Int8x16/Int16x8/Int16x16
ExtendToInt64 Int16x8/Int32x4/In32x8
ExtendToUint16 Uint8x16/Uint8x32 - zero extend
ExtendToUint32 Uint8x16/Uint16x8/Uint16x16
ExtendToUint64 Uint16x8/Uint32x4/Uin32x8

SaturateToInt8  Int  16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
SaturateToInt16 Int  32x4/32x8/32x16 64x2/64x4/64x8
SaturateToInt16Concat        Int32x4
SaturateToInt16ConcatGrouped Int32x8/Int32x16
SaturateToUint32 Uint  64x2/64x4/64x8
SaturateToUint8  Uint  16x8/16x16/16x32 32x4/32x8/32x16 64x2/64x4/64x8
SaturateToUint16 Uint  32x4/32x8/32x16 64x2/64x4/64x8
SaturateToUint16Concat        Int32x4             xxx should this be uint?
SaturateToUint16ConcatGrouped Int32x8/Int32x16    xxx should this be uint?
SaturateToUint32 Uint  64x2/64x4/64x8


DotProductPairs          Int  16x8/16x16/16x32
DotProductPairsSaturated Uint 8x16/8x32/8x64

GaloisFieldAffineTransform, GaloisFieldAffineTransformInverse, GaloisFieldMul Uint 8x16/8x32/8x64

## Type Conversions (no op)


## Mask Operations

asMask
And
Or
To...


## What are:

VADDSUBPS  VADDSUBPD

VHADDPS  VHADDPD  - add pairs

VPAVGB VPAVGW 

VPERMI2B VPERMI2W - concat permute

VPALIGNR

VCVTPD2PSX
