var httpGetRequest = function (url, callbackFunc) {
    let request = cc.loader.getXMLHttpRequest();
    request.open("GET", url, true);
    request.setRequestHeader("Content-Type","application/json;charset=UTF-8");
    request.onreadystatechange = function () {
        if (request.readyState === 4) {
            // get status text
            // var httpStatus = request.statusText;
            if(callbackFunc != null || callbackFunc !== ""){
                callbackFunc(request.responseText);
            }
        }
    };
    request.send();
};