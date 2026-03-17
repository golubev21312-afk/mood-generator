import axios from "axios"

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "http://localhost:8081/api",
})

export const analyzeMood = (input) =>
  api.post("/mood", { input }).then((r) => r.data)

export const getHistory = () =>
  api.get("/history").then((r) => r.data)

export const getMoodById = (id) =>
  api.get(`/mood/${id}`).then((r) => r.data)
