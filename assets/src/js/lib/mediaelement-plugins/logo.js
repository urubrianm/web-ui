Object.assign(MediaElementPlayer.prototype, {
    watchControlsVisible(controls, callback) {
        const observer = new MutationObserver((mutations) => {
            mutations.forEach((mutation) => {
                if (mutation.attributeName === 'class') {
                    callback(!mutation.target.classList.contains(`${this.options.classPrefix}offscreen`));
                }
            });
        });
        observer.observe(controls, {
            attributes: true,
        });
    },
    async buildlogo(player, controls, layers) {
        this.watchControlsVisible(controls, (visible) => {
          
        });
    }
});