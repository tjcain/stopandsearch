module.exports = {
    assetsDir: 'static',

    devServer: {
        proxy: {
            '/api/v1/': {
                target: 'http://localhost:4000'
            }
        }
    }
}