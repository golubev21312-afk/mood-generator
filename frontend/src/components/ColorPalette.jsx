import { useState } from "react"

function Swatch({ color, onCopy }) {
  const [copied, setCopied] = useState(false)

  const handleClick = () => {
    navigator.clipboard.writeText(color.hex)
    setCopied(true)
    setTimeout(() => setCopied(false), 1500)
    onCopy?.()
  }

  return (
    <div
      onClick={handleClick}
      className="flex-1 min-w-[100px] rounded-2xl overflow-hidden cursor-pointer
                 hover:scale-105 transition-transform duration-200 shadow-lg"
    >
      <div className="h-32" style={{ backgroundColor: color.hex }} />
      <div className="p-3 text-center" style={{ background: "rgba(0,0,0,0.18)" }}>
        <p className="font-mono text-sm font-bold" style={{ color: "rgba(255,255,255,0.95)" }}>
          {copied ? "Скопировано!" : color.hex}
        </p>
        <p className="text-xs mt-1" style={{ color: "rgba(255,255,255,0.7)" }}>{color.name}</p>
        <p className="text-xs" style={{ color: "rgba(255,255,255,0.5)" }}>{color.role}</p>
      </div>
    </div>
  )
}

export default function ColorPalette({ palette, onCopy }) {
  return (
    <div className="w-full animate-slide-up">
      <h2 className="text-sm uppercase tracking-widest mb-4"
          style={{ color: "rgba(0,0,0,0.5)" }}>
        Цветовая палитра
      </h2>
      <div className="flex gap-3 flex-wrap">
        {palette.map((color, i) => (
          <Swatch key={i} color={color} onCopy={onCopy} />
        ))}
      </div>
    </div>
  )
}
