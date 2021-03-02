package ipc

const (
	ReqAttemptDecryption = iota + 1
	ResDecryptionFailed
	ResJsonDeserializationFailed
	ReqLoadFile
	ResRequireConfigFile
	ReqCloseConnection
	ReqStatus
	ResReadyToServe
	ResNeedDecryption
	ResSuccess
	ResError
)

type Code int

type Message struct {
	Code    Code
	Payload []byte
}

var msgCodeToString = map[uint16]string{
	ReqAttemptDecryption:         `REQ_ATTEMPT_DECRYPTION`,
	ResDecryptionFailed:          `RES_DECRYPTION_FAILED`,
	ResJsonDeserializationFailed: `RES_JSON_DESERIALIZATION_FAILED`,
	ReqLoadFile:                  `REQ_LOAD_FILE`,
	ResRequireConfigFile:         `RES_REQUIRE_CONFIG_FILE`,
	ReqCloseConnection:           `REQ_CLOSE_CONNECTION`,
	ReqStatus:                    `REQ_STATUS`,
	ResReadyToServe:              `RES_READY_TO_SERVE`,
	ResNeedDecryption:            `RES_NEED_DECRYPTION`,
	ResSuccess:                   `RES_SUCCESS`,
	ResError:                     `RES_ERROR`,
}
