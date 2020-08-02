package lib

type Wav struct{
	RHead []byte //4
	Size []byte	//4
	_Size uint32

	WHead []byte //4
	FHead []byte //4

	Bytes []byte //4
	FmtId []byte //2 
	Chnum []byte //2
	SRate []byte //4
	Speed []byte //4
	BSize []byte //2
	BRate []byte //2
//	ESize []byte //2
	DHead []byte //4
	DSize []byte //4
	DATA []byte

	_SRate uint32
	_DSize uint32
}

//func MkWav() Wav
