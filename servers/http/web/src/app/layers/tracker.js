var TrackerLayer = cc.Layer.extend({

    lastTrackerUUID: 0, // 只有其uuid大于这个数值的，才会被执行，避免重复

    init: function (){
        if (this._super()){
            var that = this;
            Provider.getTrackerMap().forEach(function(val, uuid) {
                that.trackerCreator(val, uuid);
                that.lastTrackerUUID = uuid;
            })
        }
    },

    trackerCreator: function (trackerObj, uuid) {
        let trackerSprite = cc.Sprite.create(gameResources.images.tracker);
        trackerSprite.setScale(0.2);
        let bucketBox = Provider.getBucketBoundingBoxesMap().get(trackerObj.buckets[0].Hash);
        trackerSprite.setPosition(bucketBox.bounding.x + bucketBox.bounding.width/2, bucketBox.bounding.y + bucketBox.bounding.height/2);
        // trackerSprite.setAnchorPoint(0, 0);
        trackerSprite.addComponent(new trackerComponent());

        let actions = this.actionsCreator(trackerObj);

        // At the end
        actions.push(new cc.callFunc(this.removeFromParent, trackerSprite));
        actions.push(new cc.callFunc(() => {
            Provider.getTrackerMap().delete(uuid);
        }, trackerSprite));

        trackerSprite.runAction(new cc.Sequence(actions));
        this.addChild(trackerSprite, 3);
        trackerObj.isRunning = true;

        return trackerSprite.getBoundingBox();
    },

    actionsCreator: function (trackerObj) {
        let actions = [];

        if (trackerObj.len < 2) {
            actions.push(new cc.RotateBy(2, 180));
            return actions;
        }

        for (let i = 0; i < trackerObj.len; i++){
            // get bucket position
            let bucketBoundingBox = Provider.getBucketBoundingBoxesMap().get(trackerObj.buckets[i].Hash);
            actions.push(
                cc.moveTo(getDurationInBuckets(trackerObj.buckets[i].Items), cc.p(bucketBoundingBox.bounding.x + bucketBoundingBox.bounding.width/2, bucketBoundingBox.bounding.y + bucketBoundingBox.bounding.height/2))
            )
        }
        return actions;
    },

    refresh: function () {
        let that = this;
        Provider.getTrackerMap().forEach(function(val, uuid) {
            if (!val.isRunning && that.lastTrackerUUID > uuid) {
                that.trackerCreator(val, uuid);
                that.lastTrackerUUID = uuid;
            }
        })
    }
});