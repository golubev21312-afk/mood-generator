export default function TrackCard({ track }) {
  const href = `https://www.youtube.com/results?search_query=${encodeURIComponent(track.artist + " " + track.title)}`

  return (
    <a
      href={href}
      target="_blank"
      rel="noopener noreferrer"
      className="flex items-center gap-4 p-4 rounded-xl transition-all duration-200 group animate-fade-in"
      style={{ background: "rgba(0,0,0,0.12)", border: "1px solid rgba(0,0,0,0.08)" }}
    >
      {track.cover ? (
        <img
          src={track.cover}
          alt={track.title}
          className="w-14 h-14 rounded-lg object-cover flex-shrink-0 shadow-md"
        />
      ) : (
        <div className="w-14 h-14 rounded-lg flex-shrink-0"
             style={{ background: "rgba(0,0,0,0.15)" }} />
      )}
      <div className="min-w-0">
        <p className="font-medium truncate" style={{ color: "rgba(0,0,0,0.8)" }}>
          {track.title}
        </p>
        <p className="text-sm truncate" style={{ color: "rgba(0,0,0,0.5)" }}>{track.artist}</p>
      </div>
      <div className="ml-auto flex-shrink-0 transition-colors" style={{ color: "rgba(0,0,0,0.35)" }}>
        ↗
      </div>
    </a>
  )
}
