module pkg1;

import (
    "pkg2"
)

message Message {
    field_bool    bool      1;
    field_enum    Enum      2;

    field_int8    int8      10;
    field_int16   int16     11;
    field_int32   int32     12;
    field_int64   int64     13;

    field_uint8   uint8     20;
    field_uint16  uint16    21;
    field_uint32  uint32    22;
    field_uint64  uint64    23;

    field_float32 float32   30;
    field_float64 float64   31;

    field_u128    u128  40;
    field_u256    u256  41;

    field_string  string    50;
    field_bytes   bytes     51;

    list        []int64         60;
    strings     []string        61;
    messages    []SubMessage    62;
}

message Node {
    value   string  1;
    next    *Node   2;
}