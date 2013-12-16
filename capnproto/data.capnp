using Go = import "./go.capnp";

# needed to know what package should be used for generated code
$Go.package("capnproto");

# Needed to know how to import types in the capnp file and whether two
# capnp files are in the same package
$Go.import("github.com/xiangli-cmu/serialization-bench/capnproto");

@0x832bcc6686a26d56;

struct TestST {
  n @0   :Int32;
  s @1   :Text;
  a @2   :Data;
}
