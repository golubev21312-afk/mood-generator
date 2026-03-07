import { useEffect } from "react"

export function useBackgroundColor(palette) {
  useEffect(() => {
    if (!palette || palette.length === 0) return

    const primary = palette.find((c) => c.role === "основной") || palette[0]
    const bg = palette.find((c) => c.role === "фон") || palette[1]

    document.body.style.background = bg
      ? `linear-gradient(135deg, ${primary.hex} 0%, ${bg.hex} 100%)`
      : primary.hex

    return () => {
      document.body.style.background = ""
    }
  }, [palette])
}
