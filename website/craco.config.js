const path = require('path');
const {BundleAnalyzerPlugin} = require("webpack-bundle-analyzer");

module.exports = function ({env}) {
    console.log(env)
    return {
        webpack: {
            // 别名
            alias: {
                '~': path.resolve('src')
            },
            plugins: [
                //打包分析
                // new BundleAnalyzerPlugin(),
            ],
            configure: (webpackConfig, {env, paths}) => {
                webpackConfig.externals = {
                    react: 'React',
                    'react-dom': 'ReactDOM',
                    "antd": "antd",
                    "moment": "moment"
                }
                return webpackConfig;
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
                '/api1': {
                    target: 'http://localhost:8088',
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
