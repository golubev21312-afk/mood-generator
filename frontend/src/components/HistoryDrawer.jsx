import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import { getHistory } from "../api/mood"

const moodEmoji = {
  радостный: "😄", грустный: "😢", тревожный: "😰",
  спокойный: "😌", злой: "😠", вдохновлённый: "🚀", усталый: "😴",
}

export default function HistoryDrawer({ open, onClose }) {
  const [items, setItems] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    if (open) {
      setLoading(true)
      getHistory()
        .then(setItems)
        .finally(() => setLoading(false))
    }
  }, [open])

  return (
    <>
      {/* Overlay */}
      <div
        onClick={onClose}
        className={`fixed inset-0 bg-black/50 backdrop-blur-sm z-40 transition-opacity duration-300 ${
          open ? "opacity-100" : "opacity-0 pointer-events-none"
        }`}
      />

      {/* Drawer */}
      <aside
        className={`
          fixed top-0 right-0 h-full w-full max-w-sm z-50
          bg-gray-900/95 backdrop-blur-md border-l border-white/10
          transform transition-transform duration-300
          ${open ? "translate-x-0" : "translate-x-full"}
          flex flex-col
        `}
      >
        <div className="flex items-center justify-between p-6 border-b border-white/10">
          <h2 className="text-white font-semibold text-lg">История</h2>
          <button
            onClick={onClose}
            className="text-white/50 hover:text-white text-2xl transition-colors"
          >
            ✕
          </button>
        </div>

        <div className="flex-1 overflow-y-auto p-4 space-y-3">
          {loading && (
            <p className="text-white/40 text-center py-8">Загрузка...</p>
          )}
          {!loading && items.length === 0 && (
            <p className="text-white/40 text-center py-8">
              История пуста
            </p>
          )}
          {items.map((item) => (
            <Link
              key={item.id}
              to={`/mood/${item.id}`}
              onClick={onClose}
              className="
                block p-4 rounded-xl bg-white/5 hover:bg-white/10
                border border-white/10 transition-all duration-200
              "
            >
              <div className="flex items-center gap-3">
                <span className="text-2xl">
                  {moodEmoji[item.mood_label] || "🎭"}
                </span>
                <div className="min-w-0">
                  <p className="text-white text-sm font-medium capitalize">
                    {item.mood_label || "—"}
                    {item.energy && (
                      <span className="text-white/40 font-normal ml-2">
                        {item.energy}/10
                      </span>
                    )}
                  </p>
                  <p className="text-white/40 text-xs truncate mt-0.5">
                    {item.user_input}
                  </p>
                </div>
              </div>
              {item.palette && (
                <div className="flex gap-1 mt-3">
                  {item.palette.slice(0, 3).map((c, i) => (
                    <div
                      key={i}
                      className="h-2 flex-1 rounded-full"
                      style={{ backgroundColor: c.hex }}
                    />
                  ))}
                </div>
              )}
              <p className="text-white/30 text-xs mt-2">
                {new Date(item.created_at).toLocaleDateString("ru-RU", {
                  day: "numeric", month: "short", hour: "2-digit", minute: "2-digit"
                })}
              </p>
            </Link>
          ))}
        </div>
      </aside>
    </>
  )
}
