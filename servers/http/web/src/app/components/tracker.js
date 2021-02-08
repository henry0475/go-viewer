var trackerComponent = cc.Component.extend({
    onEnter: function () {
        console.log("enter");
        // let owner = this.getOwner();
        // owner.layerComponent = this;
    },
    onExit: function (){
        console.log("exit")
    },
    update: function (){}
})