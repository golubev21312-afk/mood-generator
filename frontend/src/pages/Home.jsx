import { useState } from "react"
import { useNavigate } from "react-router-dom"
import MoodInput from "../components/MoodInput"
import LoadingScreen from "../components/LoadingScreen"
import HistoryDrawer from "../components/HistoryDrawer"
import { analyzeMood } from "../api/mood"

export default function Home() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [historyOpen, setHistoryOpen] = useState(false)
  const navigate = useNavigate()

  const handleSubmit = async (input) => {
    setLoading(true)
    setError("")
    try {
      const result = await analyzeMood(input)
      navigate(`/mood/${result.request_id}`, { state: result })
    } catch {
      setError("Что-то пошло не так. Попробуй ещё раз.")
    } finally {
      setLoading(false)
    }
  }

  return (
    <>
      {loading && <LoadingScreen />}
      <HistoryDrawer open={historyOpen} onClose={() => setHistoryOpen(false)} />

      <button
        onClick={() => setHistoryOpen(true)}
        className="fixed top-6 right-6 px-4 py-2 rounded-xl bg-white/10 hover:bg-white/20 text-white text-sm transition-all z-10"
      >
        История
      </button>

      <main className="min-h-screen bg-gradient-to-br from-gray-900 to-gray-800 flex flex-col items-center justify-center p-4 md:p-8">
        <div className="text-center mb-12 animate-fade-in">
          <h1 className="text-4xl md:text-5xl font-bold text-white mb-4">
            Mood Generator
          </h1>
          <p className="text-white/60 text-base md:text-lg max-w-md">
            Опиши своё состояние — получи палитру, цитату и плейлист
          </p>
        </div>
        <MoodInput onSubmit={handleSubmit} loading={loading} />
        {error && (
          <div className="mt-4 p-4 rounded-xl bg-red-500/20 border border-red-500/30 text-red-200 text-sm max-w-2xl w-full">
            {error}
          </div>
        )}
      </main>
    </>
  )
}
