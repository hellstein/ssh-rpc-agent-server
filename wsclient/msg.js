const protobuf = require("protobufjs");
const COMMON_TYPE = 0;
const WINDOWSIZE_TYPE = 1;
const protofile = "channelmsg.proto";
/*
protobuf.load("channelmsg.proto", function(err, root) {
    if (err)
        throw err;

    // Obtain a message type
    const MsgSchema = root.lookupType("jobmgr.CommonMsg");

    // Exemplary payload
    let payload = { id: 0, content: "" };

    // Verify the payload if necessary (i.e. when possibly incomplete or invalid)
    let errMsg = MsgSchema.verify(payload);
    if (errMsg)
        throw Error(errMsg);

    // Create a new message
    let msg = MsgSchema.create(payload); // or use .fromObject if conversion is necessary
    console.log(msg)
    // Encode a message to an Uint8Array (browser) or Buffer (node)
    let buffer = MsgSchema.encode(msg).finish();
    // ... do something with buffer
    console.log(buffer)

    // Decode an Uint8Array (browser) or Buffer (node) to a message
    let message = MsgSchema.decode(buffer);
    // ... do something with message
    console.log(message)
    // If the application uses length-delimited buffers, there is also encodeDelimited and decodeDelimited.

    // Maybe convert the message back to a plain object
    var object = AwesomeMessage.toObject(message, {
        longs: String,
        enums: String,
        bytes: String,
        // see ConversionOptions
    });
});
*/

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
