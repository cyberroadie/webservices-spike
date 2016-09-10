var path=require('path')
var packageConfig = require('./package')

module.exports = {
    entry: "./main.js",
    output: {
        path: path.join(__dirname, '/public'),
        filename: "bundle.js"
    },
    module: {
        loaders: [
            { test: /\.css$/, loader: "style!css" },
            { loader: "babel-loader", query: packageConfig.babel},
        ]
    }
};
