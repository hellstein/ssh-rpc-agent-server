const COMMON_TYPE = 0;
const WINDOWSIZE_TYPE = 1;

function commonMsg(data) {
    return '1'+data;
}

function windowSizeMsg(data) {
    return '0'+data;
}

module.exports = {
    commonMsg,
    windowSizeMsg
}
