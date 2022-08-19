const path = require('path');

module.exports = function ({env}) {
    return {
        devServer: {
            port: 3001, // 端口配置
            proxy: {
                '/api1': {
                    target: 'http://127.0.0.1:1323/',
                    ws: false, // websocket
                    changeOrigin: true, //是否跨域
                    secure: false,  // 如果是https接口，需要配置这个参数
                    pathRewrite: {
                        '^/api1': '/api1'
                    }
                }

            }
        }

    }
}
