const path = require('path');

module.exports = function({ env }) {
    return {
        webpack: {
            // 别名
            alias: {
                '~':  path.resolve('src')
            }
        },
        devServer: {
            port: 9999, // 端口配置
            proxy: {
                '/lpv': {
                    target: 'http://api.linkpreview.net',
                    ws: false, // websocket
                    changeOrigin: true, //是否跨域
                    secure: false,  // 如果是https接口，需要配置这个参数
                    pathRewrite: {
                        '^/lpv': ''
                    }
                },
                '/api': {
                    target: 'http://127.0.0.1:1323/',
                    ws: false, // websocket
                    changeOrigin: true, //是否跨域
                    secure: false,  // 如果是https接口，需要配置这个参数
                    pathRewrite: {
                        '^/api': '/api1'
                    }
                }

            }
        }

    }
}
