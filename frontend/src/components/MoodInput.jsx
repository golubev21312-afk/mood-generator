import { useState } from "react"

export default function MoodInput({ onSubmit, loading }) {
  const [value, setValue] = useState("")

  const handleSubmit = (e) => {
    e.preventDefault()
    if (value.trim()) onSubmit(value.trim())
  }

  return (
    <form onSubmit={handleSubmit} className="w-full max-w-2xl">
      <textarea
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Как ты себя чувствуешь прямо сейчас?.."
        rows={5}
        disabled={loading}
        className="
          w-full p-4 md:p-6 rounded-2xl text-base md:text-lg resize-none
          bg-white/10 backdrop-blur-sm border border-white/20
          text-white placeholder-white/50
          focus:outline-none focus:ring-2 focus:ring-white/40
          transition-all duration-300
        "
      />
      <button
        type="submit"
        disabled={loading || !value.trim()}
        className="
          mt-4 w-full py-4 rounded-2xl text-lg font-semibold
          bg-white/20 hover:bg-white/30 text-white
          disabled:opacity-40 disabled:cursor-not-allowed
          transition-all duration-200 active:scale-95
        "
      >
        {loading ? "Анализирую..." : "Понять своё настроение"}
      </button>
    </form>
  )
}
