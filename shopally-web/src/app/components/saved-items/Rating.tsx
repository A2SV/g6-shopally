// shopally-web/src/app/components/saved-items/Rating.tsx
import { Star } from "lucide-react";

export default function Rating({
  value,
  count,
}: {
  value: number;
  count: number;
}) {
  return (
    <div className="flex items-center gap-1">
      {Array.from({ length: 5 }).map((_, i) => (
        <Star
          key={i}
          className={`w-3 h-3 ${
            i < Math.round(value)
              ? "fill-yellow-400 stroke-yellow-400"
              : "stroke-gray-300"
          }`}
        />
      ))}
      <span className="ml-1 text-sm text-gray-600">{value}</span>
      <span className="ml-1 text-xs text-gray-400">({count})</span>
    </div>
  );
}
