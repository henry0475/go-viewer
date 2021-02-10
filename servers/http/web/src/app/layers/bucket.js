var BucketLayer = cc.Layer.extend({
    space: 100,
    countOfBuckets: 0, // How many buckets are showing in this layer
    originPosition: {x: 100, y: 0},
    nextBucketPosition: {x: 100, y: 0},

    init: function () {
        if (this._super()) {

            let winSize = cc.director.getWinSize();
            let x = this.originPosition.x;
            let y = this.originPosition.y;
            let that = this;

            Provider.getBucketBoundingBoxesMap().forEach(function (bucketObj, bucketHash){
                if (!bucketObj.isShowing) {
                    let boundingBox = that.bucketCreator(bucketObj, x, y);
                    that.countOfBuckets += 1;

                    x += boundingBox.width + that.space;

                    if (x + boundingBox.width > winSize.width) {
                        x = that.originPosition.x;
                        y -= boundingBox.height + that.space;
                    }

                    that.nextBucketPosition = {x: x, y: y};
                }
            })
        }
    },

    bucketCreator: function (bucketObj, x, y) {
        let bucketSprite = cc.Sprite.create(gameResources.images.bucket);
        bucketSprite.setScale(0.9);
        bucketSprite.setPosition(x, y);
        let bucketBoundingBox = bucketSprite.getBoundingBox();

        let nameLabel = cc.LabelTTF.create(bucketObj.Name, "Arial", 20);
        nameLabel.setPosition(bucketBoundingBox.width/2, bucketBoundingBox.height + 40);
        bucketSprite.addChild(nameLabel, 3);

        let hashLabel = cc.LabelTTF.create(bucketObj.Hash, "Arial", 8)
        hashLabel.setPosition(bucketBoundingBox.width/2, -10)
        hashLabel.setHorizontalAlignment(cc.TEXT_ALIGNMENT_CENTER);
        bucketSprite.addChild(hashLabel, 3);

        this.addChild(bucketSprite, 3);

        return bucketSprite;
    },

    refresh: function () {
        let that = this;
        let x = this.nextBucketPosition.x;
        let y = this.nextBucketPosition.y;
        let winSize = cc.director.getWinSize();

        Provider.getBucketBoundingBoxesMap().forEach(function (obj, bucketHash){
            if (!obj.isShowing) {
                // Has new bucket
                let bucketSprite = that.bucketCreator(obj.bucket, x, y + 150);
                that.countOfBuckets += 1;
                obj.isShowing = true;
                obj.bounding = bucketSprite.getBoundingBox();

                x += obj.bounding.width + that.space;

                if (x + obj.bounding.width > winSize.width) {
                    x = that.originPosition.x;
                    y -= obj.bounding.height + that.space;
                }

                that.nextBucketPosition = {x: x, y: y};
            }
        })
    }
});