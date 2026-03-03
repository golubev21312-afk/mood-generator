import React from "react"
import ReactDOM from "react-dom/client"
import { BrowserRouter, Routes, Route } from "react-router-dom"
import "./index.css"
import Home from "./pages/Home"
import Result from "./pages/Result"

ReactDOM.createRoot(document.getElementById("root")).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/mood/:id" element={<Result />} />
    </Routes>
  </BrowserRouter>
)
