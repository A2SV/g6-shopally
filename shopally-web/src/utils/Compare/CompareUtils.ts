// shopally-web/src/utils/Compare/CompareUtils.ts
import type {
  ComparisonItem,
  ComparisonResponse,
} from "../../types/Compare/Comparison";

export async function fetchComparison(): Promise<ComparisonItem[]> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_BASE_URL ?? ""}/data/comparison.json`,
    {
      cache: "no-store",
    }
  );

  if (!res.ok) {
    const errorText = await res.text();
    throw new Error(`Failed to fetch comparison: ${res.status} - ${errorText}`);
  }

  const data: ComparisonResponse = await res.json();
  console.log("Fetched comparison:", data);

  // unwrap and return just the array
  return data.data?.comparison ?? [];
}
