var trackerMap = new Map();
// bucketsMap contains all buckets the scene is using
var bucketBoundingBoxesMap = new Map();

var Provider = {
    // Should be used as long as the data received from the server
    push: function (containers) {
        let lengthOfContainers = containers.length;

        for (let i = 0; i < lengthOfContainers; i++) {
            let tLen = containers[i].Trackers.length;
            for (let k = 0; k < tLen; k++) {
                // For tracker's map
                if (!trackerMap.has(containers[i].Trackers[k].UUID)) {
                    trackerMap.set(containers[i].Trackers[k].UUID, {
                        buckets: containers[i].Trackers[k].Buckets,
                        len: containers[i].Trackers[k].Buckets.length,
                        isRunning: false
                    });
                }

                // For bucket's map
                for (let j = 0; j < containers[i].Trackers[k].Buckets.length; j++) {
                    let bucketObj = containers[i].Trackers[k].Buckets[j];
                    if (!bucketBoundingBoxesMap.has(bucketObj.Hash)){
                        bucketBoundingBoxesMap.set(bucketObj.Hash, {
                            ref: 0,
                            bounding: {x: 0, y: 0, width: 0, height: 0},
                            isShowing: false,
                            bucket: bucketObj
                        })
                    }
                }
            }
        }
    },

    getTrackerMap: function () {
        return trackerMap;
    },

    getBucketBoundingBoxesMap: function () {
        return bucketBoundingBoxesMap;
    }
}