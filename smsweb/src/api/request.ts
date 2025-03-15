import axios from 'axios'

const requests = axios.create({
    baseURL: "http://localhost:801",
    timeout: 10000,
})

export  default  requests



