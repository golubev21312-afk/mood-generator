export default function LoadingScreen() {
  return (
    <div className="fixed inset-0 flex flex-col items-center justify-center bg-gray-900 z-50">
      <div className="relative w-24 h-24 mb-8">
        <div className="absolute inset-0 rounded-full border-4 border-white/20 animate-ping" />
        <div className="absolute inset-0 rounded-full border-4 border-white/60 animate-pulse-slow" />
        <div className="absolute inset-4 rounded-full bg-white/10 backdrop-blur-sm" />
      </div>
      <p className="text-white/80 text-xl font-light animate-pulse">
        Читаю настроение...
      </p>
      <p className="text-white/40 text-sm mt-2">
        Groq анализирует твои слова
      </p>
    </div>
  )
}
