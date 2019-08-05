// 云函数入口文件
const cloud = require('wx-server-sdk')
const request = require('request-promise')

cloud.init()

// 云函数入口函数
exports.main = async (event, context) => {
  return request.get('http://123.207.7.254:8080/api/v1/blog/search')
}