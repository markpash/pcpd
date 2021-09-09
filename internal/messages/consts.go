package messages

const Version = 0x02

type OpCode uint8

const (
	Announce OpCode = 0
	Map      OpCode = 1
	Peer     OpCode = 2
)

func (o OpCode) String() string {
	switch o {
	case Announce:
		return "ANNOUNCE"
	case Map:
		return "MAP"
	case Peer:
		return "PEER"
	default:
		return ""
	}
}

type Result uint8

const (
	Success               Result = 0
	UnsuppVersion         Result = 1
	NotAuthorized         Result = 2
	MalformedRequest      Result = 3
	UnsuppOpcode          Result = 4
	UnsuppOption          Result = 5
	MalformedOption       Result = 6
	NetworkFailure        Result = 7
	NoResources           Result = 8
	UnsuppProtocol        Result = 9
	UserExQuota           Result = 10
	CannotProvideExternal Result = 11
	AddressMismatch       Result = 12
	ExcessiveRemotePeers  Result = 13
)

func (r Result) String() string {
	switch r {
	case Success:
		return "SUCCESS"
	case UnsuppVersion:
		return "UNSUPP_VERSION"
	case NotAuthorized:
		return "NOT_AUTHORIZED"
	case MalformedRequest:
		return "MALFORMED_REQUEST"
	case UnsuppOpcode:
		return "UNSUPP_OPCODE"
	case UnsuppOption:
		return "UNSUPP_OPTION"
	case MalformedOption:
		return "MALFORMED_OPTION"
	case NetworkFailure:
		return "NETWORK_FAILURE"
	case NoResources:
		return "NO_RESOURCES"
	case UnsuppProtocol:
		return "UNSUPP_PROTOCOL"
	case UserExQuota:
		return "USER_EX_QUOTA"
	case CannotProvideExternal:
		return "CANNOT_PROVIDE_EXTERNAL"
	case AddressMismatch:
		return "ADDRESS_MISMATCH"
	case ExcessiveRemotePeers:
		return "EXCESSIVE_REMOTE_PEERS"
	default:
		return ""
	}
}

type Protocol uint16

const (
	TCP Protocol = 6
	UDP Protocol = 17
)

func (p Protocol) String() string {
	switch p {
	case TCP:
		return "tcp"
	case UDP:
		return "udp"
	default:
		return ""
	}
}
