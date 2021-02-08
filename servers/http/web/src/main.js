window.onload = function(){

    cc.game.onStart = function(){
        cc.view.adjustViewPort(true);
        cc.view.setDesignResolutionSize(1280, 800, cc.ResolutionPolicy.SHOW_ALL);
        cc.view.resizeWithBrowserSize(true);

        //load resources
        cc.LoaderScene.preload(getGameResources(), function () {
            cc.director.runScene(new PrepareScene());
        }, this);
    };

    cc.game.run("view");
};