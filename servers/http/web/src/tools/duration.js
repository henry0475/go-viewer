var getDurationInBuckets = function (itemsInBucket) {
    let duration = 0;

    for (let i = 0; i < itemsInBucket.length; i++) {
        duration += itemsInBucket[i].Duration;
    }

    duration = duration / 1e9;
    if (duration < 1) {
        duration = 1.5;
    } else {
        duration += 2;
    }

    return duration;
};