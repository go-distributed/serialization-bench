package gobin

type PreAccept struct {
        LeaderId int32
        Replica  int32
        Instance int32
        Ballot   int32
        Command  []byte
        Seq      int32
        Deps     [5]int32
}
