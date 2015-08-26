namespace java addservice
namespace php addservice

struct Response {
	1: i32 code,
	2: string desc,
	3: string data
}

struct Request {
	1: string id
}

service AddressService {
	Response getAllAddress(1:Request req),
}
