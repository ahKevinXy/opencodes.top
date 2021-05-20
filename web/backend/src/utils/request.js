import axios from 'axios'
import context from '@/main.js'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 99999
})
let activeAxios = 0
let timer

const showLoading = () => {
  activeAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (activeAxios > 0) {
      context.$bus.emit('showLoadding')
    }
  }, 400)
}
const closeLoading = () => {
  activeAxios--
  if (activeAxios <= 0) {
    clearTimeout(timer)
    context.$bus.emit('closeLoading')
  }
}

service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const token = store.getters['user/token']
    const user = store.getters['user/userInfo']
    config.data = JSON.stringify(config.data)
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': token,
      'x-user-id': user.ID
    }
    return config
  },
  error => {
    closeLoading()
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)
