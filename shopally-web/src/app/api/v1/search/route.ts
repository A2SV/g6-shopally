//src/app/api/search/route.ts
import { ProductResponse } from "@/types/types";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function GET(
  req: NextRequest
): Promise<NextResponse<ProductResponse>> {
  try {
    const { searchParams } = new URL(req.url);
    const query = searchParams.get("q");
    const priceMaxETB = searchParams.get("priceMaxETB");
    const minRating = searchParams.get("minRating");

    // Get deviceId and Language from cookies
    const cookieStore = await cookies();
    const deviceId = cookieStore.get("deviceId")?.value;
    const langCode = cookieStore.get("lang")?.value || "en";

    // If deviceId is missing, inform client to wait for FCM token
    if (!deviceId) {
      return NextResponse.json(
        {
          error:
            "Device ID not available yet. Please wait for notifications setup.",
          data: null,
        },
        { status: 409 } // 409 Conflict to indicate precondition not met
      );
    }

    if (!query) {
      return NextResponse.json(
        { error: "Missing required fields", data: null },
        { status: 400 }
      );
    }

    const url = `${API_BASE}/api/v1/search?${new URLSearchParams({
      q: query,
      ...(priceMaxETB ? { priceMaxETB } : {}),
      ...(minRating ? { minRating } : {}),
    }).toString()}`;

    const backendRes = await fetch(url, {
      headers: {
        "X-Device-ID": deviceId,
        "Accept-Language": langCode,
      },
    });

    const contentType = backendRes.headers.get("content-type") || "";
    const rawText = await backendRes.text();

    console.log("🔎 Backend status:", backendRes.status);
    console.log("🔎 Backend content-type:", contentType);
    console.log("🔎 Backend raw response:", rawText.slice(0, 300));

    let data: ProductResponse;
    try {
      data = JSON.parse(rawText) as ProductResponse;
    } catch (err) {
      console.error("❌ Failed to parse JSON:", err);
      return NextResponse.json(
        { error: "Backend returned invalid JSON", data: null },
        { status: 500 }
      );
    }

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: data?.error || "Failed to search products", data: null },
        { status: backendRes.status }
      );
    }

    return NextResponse.json(data, { status: 200 });
  } catch (error) {
    console.error("GET /api/v1/search error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}
