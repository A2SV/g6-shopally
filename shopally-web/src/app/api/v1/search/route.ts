//src/app/api/products/route.ts
import { getLanguage } from "@/lib/redux/languageBridge";
import { ProductResponse } from "@/types/types";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function POST(
  req: NextRequest
): Promise<NextResponse<ProductResponse>> {
  try {
    const body = await req.json();
    const { query, priceMaxETB, minRating } = body;

    // ✅ Get deviceId from headers
    const deviceId = req.headers.get("x-device-id");

    if (!query || !deviceId) {
      return NextResponse.json(
        { error: "Missing required fields", data: null },
        { status: 400 }
      );
    }

    const langCode = getLanguage() || "en-US";

    const backendRes = await fetch(`${API_BASE}/api/products/search`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-Device-ID": deviceId,
        "Accept-Language": langCode, // TODO: make dynamic later
      },
      body: JSON.stringify({ query, priceMaxETB, minRating }),
    });

    const data = await backendRes.json();

    if (!backendRes.ok) {
      return NextResponse.json(
        {
          error: data?.status || "Failed to search products",
          data: null,
        },
        { status: backendRes.status }
      );
    }

    return NextResponse.json({ data, error: null }, { status: 200 });
  } catch (error) {
    console.error("POST /api/products/search error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}
