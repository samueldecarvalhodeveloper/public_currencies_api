package concerns

import (
	"load_balancing_proxy_service/constants"

	"google.golang.org/protobuf/proto"
)

func GetNotFoundErrorMessageResponse() []byte {
	var notFoundErrorMessageResponseInstance, _ = proto.Marshal(&NotFoundErrorMessageResponseEntity{
		Message: constants.NOT_FOUND_ERROR_RESPONSE_MESSAGE,
	})

	return notFoundErrorMessageResponseInstance
}
