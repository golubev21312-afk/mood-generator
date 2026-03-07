import { useEffect, useState } from "react"
import { useParams, useLocation, Link } from "react-router-dom"
import { getMoodById } from "../api/mood"
import { useBackgroundColor } from "../hooks/useBackgroundColor"
import ColorPalette from "../components/ColorPalette"
import QuoteCard from "../components/QuoteCard"
import TrackCard from "../components/TrackCard"

const moodEmoji = {
  радостный: "😄", грустный: "😢", тревожный: "😰",
  спокойный: "😌", злой: "😠", вдохновлённый: "🚀", усталый: "😴",
}

export default function Result() {
  const { id } = useParams()
  const location = useLocation()
  const [data, setData] = useState(location.state || null)
  const [loading, setLoading] = useState(!location.state)
  const [visible, setVisible] = useState(false)

  useBackgroundColor(data?.palette)

  useEffect(() => {
    if (!data) {
      getMoodById(id)
        .then(setData)
        .finally(() => setLoading(false))
    }
  }, [id])

  useEffect(() => {
    if (data) {
      const timer = setTimeout(() => setVisible(true), 100)
      return () => clearTimeout(timer)
    }
  }, [data])

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-900">
        <p className="text-white/60 text-lg">Загрузка...</p>
      </div>
    )
  }

  if (!data) {
    return (
      <div className="min-h-screen flex flex-col items-center justify-center bg-gray-900 gap-4">
        <p className="text-white/60 text-lg">Результат не найден</p>
        <Link to="/" className="text-white/80 underline">На главную</Link>
      </div>
    )
  }

  const copyLink = () => navigator.clipboard.writeText(window.location.href)

  return (
    <main className="min-h-screen p-6 md:p-12">
      <div className={`max-w-2xl mx-auto space-y-8 transition-all duration-700 ${visible ? "opacity-100" : "opacity-0"}`}>

        <div className="text-center">
          <span className="text-6xl">{moodEmoji[data.mood_label] || "🎭"}</span>
          <h1 className="text-3xl font-bold mt-4 capitalize"
              style={{ color: "rgba(0,0,0,0.75)" }}>
            {data.mood_label}
          </h1>
          <p className="text-sm mt-2" style={{ color: "rgba(0,0,0,0.5)" }}>Энергия: {data.energy}/10</p>
        </div>

        <ColorPalette palette={data.palette || []} />

        <QuoteCard quote={data.quote} author={data.quote_author} />

        {data.tracks?.length > 0 && (
          <div>
            <h2 className="text-sm uppercase tracking-widest mb-4"
                style={{ color: "rgba(0,0,0,0.5)" }}>
              Плейлист
            </h2>
            <div className="space-y-3">
              {data.tracks.map((track, i) => (
                <TrackCard key={i} track={track} />
              ))}
            </div>
          </div>
        )}

        <div className="flex gap-4 justify-center pt-4">
          <button
            onClick={copyLink}
            className="px-6 py-3 rounded-xl text-sm transition-all"
            style={{ background: "rgba(0,0,0,0.15)", color: "rgba(0,0,0,0.7)" }}
          >
            Поделиться ссылкой
          </button>
          <Link
            to="/"
            className="px-6 py-3 rounded-xl text-sm transition-all"
            style={{ background: "rgba(0,0,0,0.15)", color: "rgba(0,0,0,0.7)" }}
          >
            Новый запрос
          </Link>
        </div>
      </div>
    </main>
  )
}
