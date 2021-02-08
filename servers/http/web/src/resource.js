var gameResources = {
    images: {
        bucket: 'res/images/bucket.png',
        tracker: 'res/images/tracker.png'
    },
    test: {
        helloWorld: "res/HelloWorld.png"
    }
}

var getGameResources = function () {
    let resArray = [];
    for (let resourceKey in gameResources) {
        for (let k in gameResources[resourceKey]) {
            if (gameResources[resourceKey].hasOwnProperty(k)) {
                resArray.push(gameResources[resourceKey][k]);
            }
        }
    }
    return resArray;
};