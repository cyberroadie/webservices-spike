var path=require('path')
var packageConfig = require('./package')

module.exports = {
    entry: "./main.jsx",
    output: {
        path: path.join(__dirname, '/public'),
        filename: "bundle.js"
    },
    module: {
        loaders: [
            { test: /\.css$/, loader: "style!css" },
            { 
              test: /\.jsx?$/,         // Match both .js and .jsx files
              exclude: /node_modules/, 
              loader: "babel", 
              query:
                {
                    presets:['react']
                }
            }
        ]
    },
};
