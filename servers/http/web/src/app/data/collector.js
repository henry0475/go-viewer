var collectData = function () {

    let requestURL = (configs.isDebugging) ? configs.JSONsPathForTesting + "set1.json": configs.serverURL;

    httpGetRequest(requestURL, function (res){
        let parsedRes = JSON.parse(res)
        globalVariables.collectedData = parsedRes;

        Provider.push(globalVariables.collectedData.Containers);
    })
};