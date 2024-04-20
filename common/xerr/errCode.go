package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006
const REDIS_ERROR uint32 = 100007
const UNACCESSABLE_ERROR uint32 = 100008

// 用户模块
const NO_ACCESS_TO_RESOURCE uint32 = 200001
const INSERT_ALREADY_EXSIT uint32 = 200002
const UserNotInGroup uint32 = 200003
const USER_NOT_FOUND uint32 = 200004
const MUST_CHOOSE_HEIR uint32 = 200005

// chat 模块
const WEBSOCKET_CONN_ERR uint32 = 300001
