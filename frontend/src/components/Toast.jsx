import { useState } from "react"

export function useToast() {
  const [message, setMessage] = useState("")
  const show = (msg) => {
    setMessage(msg)
    setTimeout(() => setMessage(""), 2000)
  }
  return { message, show }
}

export function Toast({ message }) {
  if (!message) return null
  return (
    <div className="fixed bottom-8 left-1/2 -translate-x-1/2 z-50
                    px-6 py-3 rounded-xl bg-white text-gray-900
                    text-sm font-medium shadow-xl animate-slide-up">
      {message}
    </div>
  )
}
