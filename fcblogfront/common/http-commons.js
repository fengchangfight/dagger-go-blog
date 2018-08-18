import axios from 'axios'

export const AXIOS = axios.create({
  baseURL: 'http://localhost:???',
  withCredentials: true
})
