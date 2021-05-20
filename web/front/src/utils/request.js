import axios from 'axios'

const service = axios.create({
    baseUrl:"",
    timeout:10000
})

service.interceptors.request.use(
    config=>{

    },
    error => {
        console.log(error)
    }
)
service.interceptors.response.use(
    response =>{
        const code = response.data.code
        console.log(code)
        return response.data
    },
    error => {
        console.log(error)
    }
)


export default service
