// shopally-web/src/utils/Home/productsUtils.ts
import type { Product, ProductResponse } from "@/types/types";

export async function fetchProducts(): Promise<Product[]> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_BASE_URL ?? ""}/data/products.json`,
    {
      cache: "no-store",
    },
  );

  if (!res.ok) {
    const errorText = await res.text();
    throw new Error(`Failed to fetch products: ${res.status} - ${errorText}`);
  }

  const data: ProductResponse = await res.json();
  console.log("Fetched products:", data);

  // unwrap and return just the array
  return data.data?.products ?? [];
}
