import (
    "pkg1"
    "pkg2"
)

options (
    go_package="github.com/basecomplextech/spec/tests/pkg4"
)

service Service {
    // Subservice doc comment
    subservice(id bin128) (sub Subservice);

    // Method doc comment.
    method(msg string);

    // Method1 doc comment.
    method1(a int64, b string, c bool) (a1 int64, b1 string, c1 bool);

    args(
        a00 bool,
        a01 byte,

        a10 int16,
        a11 int32,
        a12 int64,
        
        a20 uint16,
        a21 uint32,
        a22 uint64,

        a30 float32,
        a31 float64,

        a40 bin64,
        a41 bin128,
        a42 bin256,

        a50 string,
        a51 bytes,
        a52 message,

        a60 pkg1.Enum,
        a61 pkg1.Struct,
        a62 pkg1.Submessage,
        a63 pkg2.Submessage,

        a70 []int64,
        a71 []string,
        a72 []pkg1.Struct,
        a73 []pkg1.Submessage,
        a74 []pkg2.Submessage,

        a80 any
    ) (ok bool);

    result0() (
        a00 bool,
        a01 byte,

        a10 int16,
        a11 int32,
        a12 int64
    );

    result1() (
        a20 uint16,
        a21 uint32,
        a22 uint64,

        a30 float32,
        a31 float64
    );

    result2() (
        a40 bin64,
        a41 bin128,
        a42 bin256
    );

    result3() (
        a50 string,
        a51 bytes,
        a52 message
    );

    result4() (
        a60 pkg1.Enum,
        a61 pkg1.Struct,
        a62 pkg1.Submessage,
        a63 pkg2.Submessage
    );

    result5() (
        a70 []int64,
        a71 []string,
        a72 []pkg1.Struct,
        a73 []pkg1.Submessage
    );

    result6() (
        a74 []pkg2.Submessage,
        a80 any
    );
}

subservice Subservice {
    hello(msg string) (msg1 string);
}