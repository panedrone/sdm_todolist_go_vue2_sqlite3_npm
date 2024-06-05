import fire from "./event_bus";

const JSON_HEADERS = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
};

export function fetchJson(uri, onSuccess) {
    fetch(uri)
        .then(async (resp) => {
            if (resp.status === 200) {
                let json = await resp.json()
                onSuccess(json)
                return
            }
            await _showStatusError(resp);
        })
        .catch((reason) => {
            _showException(reason)
        })
}

export function fetchJsonArray(uri, onSuccess){
    fetch(uri)
        .then(async (resp) => {
            if (resp.status === 200) {
                let arr = await _responseToArray(resp)
                onSuccess(arr)
                return
            }
            await _showStatusError(resp);
        })
        .catch((reason) => {
            _showException(reason)
        })
}

export function fetchText(uri, onSuccess) {
    fetch(uri)
        .then(async (resp) => {
            if (resp.status === 200) {
                let text = await resp.text()
                onSuccess(text)
                return
            }
            await _showStatusError(resp);
        })
        .catch((reason) => {
            _showException(reason)
        })
}

export function postJson201(uri, json, onSuccess) {
    fetch(uri, {
        method: 'post',
        headers: JSON_HEADERS,
        body: json
    })
        .then(async (resp) => {
            if (resp.status === 201) {
                onSuccess()
                return
            }
            await _showStatusError(resp)
        })
        .catch((reason) => {
            _showException(reason)
        })
}

export function putJson200(uri, json, onSuccess) {
    fetch(uri, {
        method: 'put',
        headers: JSON_HEADERS,
        body: json
    })
        .then(async (resp) => {
            if (resp.status === 200) {
                onSuccess();
                return
            }
            await _showStatusError(resp)
        })
        .catch((reason) => {
            _showException(reason)
        })
}

export function putJson(uri, json, onResp) {
    fetch(uri, {
        method: 'put',
        headers: JSON_HEADERS,
        body: json
    })
        .then(async (resp) => {
            onResp(resp)
        })
        .catch((reason) => {
            _showException(reason)
        })
}

export function delete204(uri, onSuccess) {
    fetch(uri, {
        method: 'delete'
    })
        .then(async (resp) => {
            if (resp.status === 204) {
                onSuccess();
                return
            }
            await _showStatusError(resp)
        })
        .catch((reason) => {
            _showException(reason)
        })
}

// https://stackoverflow.com/questions/17267329/converting-unicode-character-to-string-format

export function unicodeToChar(text) {
    if (!text) {
        return ""
    }
    text = text.toString()
    return text.replace(/\\u[\dA-F]{4}/gi,
        function (match) {
            return String.fromCharCode(parseInt(match.replace(/\\u/g, ''), 16));
        });
}

async function _responseToArray(resp) {
    let contentType = resp.headers.get('content-type')
    if (!contentType) {
        console.log("no 'content-type' ==> ")
        console.log(resp)
        let msg = resp.status.toString() + " ==> (no 'content-type') ==>" + resp.toString()
        _showServerError(msg)
        return []
    }
    if (contentType.includes("application/json")) {
        let json = await resp.json()
        if (!json) {
            return []
        }
        if (Array.isArray(json)) {
            // console.log(contentType)
            // console.log(json)
            return json
        }
        let msg = resp.status.toString() + " ==> (not an Array) ==> "
        console.log(msg)
        console.log(json)
        _showServerError(msg + JSON.stringify(json))
        return []
    }
    let text = await resp.text()
    _showServerError(resp.status.toString() + " ==> (" + contentType + ") ==> " + text)
    return []
}

function _showServerError(msg) {
    msg = unicodeToChar(msg);
    msg = msg.replace(/\\"/g, '"');
    // console.log(msg)
    fire.showServerError(msg)
}

async function _showStatusError(resp) {
    let msg = await resp.text()
    if (!msg) {
        msg = "(no message)"
    }
    _showServerError(resp.status.toString() + " ==> " + msg);
}

function _showException(reason) {
    console.log(".catch((reason) => {")
    console.log(reason)
    _showServerError(reason.toString())
}
