var ViewerScene = cc.Scene.extend({

    bucketLayer: new BucketLayer(),
    trackerLayer: new TrackerLayer(),

    onEnter:function () {
        this._super();
        var size = cc.director.getWinSize();

        this.loader();

        // bucket
        this.bucketLayer.setPosition(0, size.height - 360);
        this.bucketLayer.setAnchorPoint(0, 1);
        this.bucketLayer.init();
        this.addChild(this.bucketLayer, 1);

        this.trackerLayer.setPosition(0, size.height - 360);
        this.trackerLayer.init();
        this.addChild(this.trackerLayer, 2);

        // Background
        var background = cc.Sprite.create("res/HelloWorld.png");
        background.setPosition(size.width / 2, size.height / 2);
        background.setScale(size.width/background.getTextureRect().width, size.height/background.getTextureRect().height);
        this.addChild(background, 0);

        this.touchListener();

        var connectionLabel = cc.LabelTTF.create("Connected: www.example.com", "Arial", 20);
        connectionLabel.setPosition(150, size.height - 30);
        connectionLabel.setHorizontalAlignment(cc.TEXT_ALIGNMENT_LEFT);
        this.addChild(connectionLabel, 1);

    },

    touchListener: function () {
        this.touchListener = cc.EventListener.create({
            event: cc.EventListener.TOUCH_ONE_BY_ONE,
            swallowTouches: true,
            onTouchBegan: function(touch, event){
                let pos = touch.getLocation();
                let target = event.getCurrentTarget();
                if(cc.rectContainsPoint(target.getBoundingBox(),pos)){
                    console.log(pos.x+" "+pos.y) //点击的位置
                    //在这里写点击事件
                    return true;
                }
                return false;
            }
        });
        cc.eventManager.addListener(this.touchListener,this);
    },

    loader: function () {
        collectData();
        this.bucketLayer.refresh();
        this.trackerLayer.refresh();

        let that = this;
        setTimeout(function (){
            that.loader();
        }, 1000)

    }
});
