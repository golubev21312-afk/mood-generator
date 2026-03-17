export default function QuoteCard({ quote, author }) {
  return (
    <div className="w-full p-5 md:p-8 rounded-2xl animate-slide-up"
         style={{ background: "rgba(0,0,0,0.12)", backdropFilter: "blur(8px)", border: "1px solid rgba(0,0,0,0.1)" }}>
      <p className="text-xl md:text-2xl font-light leading-relaxed italic"
         style={{ color: "rgba(0,0,0,0.75)" }}>
        &ldquo;{quote}&rdquo;
      </p>
      {author && (
        <p className="text-sm mt-4 text-right" style={{ color: "rgba(0,0,0,0.45)" }}>
          — {author}
        </p>
      )}
    </div>
  )
}
