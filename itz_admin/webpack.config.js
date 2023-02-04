
const path = require('path');
const htmlWebpackPlugin = require('html-webpack-plugin');
const webpack  = require('webpack');

const config = {
	entry: './src/index.js',
	output: {
		path: path.resolve(__dirname, 'build'),
		filename: 'boundle.js'
	},
	devServer: {
		host: "127.0.0.1",
		port: 5002,
		hot: true,
		static:{
			directory: path.join(__dirname, 'public')
		},
		hot: true, // this
		historyApiFallback: true
	},
	resolve: {
		alias: {
			svelte: path.resolve('node_modules', 'svelte')
		},
		extensions: ['*', '.mjs', '.js', '.svelte'],
		mainFields: ['svelte', 'browser', 'module', 'main'],
	},
	module: {
		rules: [
			{
				test:  /\.js?$/,
				exclude: /node_modules/,
				use: {
					loader: 'babel-loader'
				}
			},
			{
				test: /\.svelte$/,
				use: {
					loader: 'svelte-loader'
				}
				
			},
			{
				test: /\.svg$/,
				exclude: /node_modules/,
				use: {
					loader: 'svg-inline-loader',
					options: {
					  removeSVGTagAttrs: true
					}
				}
			}
		]
	},
	plugins: [
		new htmlWebpackPlugin({
			inject: true,
			template: './public/index.html',
			filename: './index.html'
		})
	]
}


module.exports = (env, argv) => {
	const isProd = argv.mode === 'production';
	const build_config = {
		AUTH_SERVER: process.env.AUTH_SERVER,
		USERS_SERVER: process.env.USERS_SERVER
	}

	config.plugins.push(
		new webpack.DefinePlugin({
			"AUTH_SERVER": JSON.stringify(build_config.AUTH_SERVER),
			"USERS_SERVER": JSON.stringify(build_config.USERS_SERVER)
		})
	);

	return config
} 