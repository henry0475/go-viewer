var collectData = function () {
    httpGetRequest("src/app/data/test/set1.json", function (res){
        let parsedRes = JSON.parse(res);
        // TODO: 临时处理，到时要去重排序等
        globalVariables.collectedData = parsedRes;

        //
        Provider.bucket(globalVariables.collectedData.Containers);
        Provider.tracker(globalVariables.collectedData.Containers);
    })
};