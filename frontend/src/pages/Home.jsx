import { useState } from "react"
import { useNavigate } from "react-router-dom"
import MoodInput from "../components/MoodInput"
import LoadingScreen from "../components/LoadingScreen"
import { analyzeMood } from "../api/mood"

export default function Home() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
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
      <main className="min-h-screen bg-gradient-to-br from-gray-900 to-gray-800 flex flex-col items-center justify-center p-6">
        <div className="text-center mb-12 animate-fade-in">
          <h1 className="text-5xl font-bold text-white mb-4">
            Mood Generator
          </h1>
          <p className="text-white/60 text-lg max-w-md">
            Опиши своё состояние — получи палитру, цитату и плейлист
          </p>
        </div>
        <MoodInput onSubmit={handleSubmit} loading={loading} />
        {error && (
          <p className="mt-4 text-red-300 text-sm">{error}</p>
        )}
      </main>
    </>
  )
}
