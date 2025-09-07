//src/app/api/v1/alerts/route.ts
import { getLanguage } from "@/lib/redux/languageBridge";
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function POST(
  req: NextRequest
): Promise<NextResponse<AlertCreateResponse>> {
  try {
    const body = await req.json();
    const { productId, currentPriceETB, deviceId: deviceIdFromBody } = body;

    // ✅ Take deviceId from headers first, fallback to body
    const deviceIdFromHeader = req.headers.get("x-device-id");
    const deviceId = deviceIdFromHeader || deviceIdFromBody;

    if (!productId || !deviceId || !currentPriceETB) {
      return NextResponse.json(
        { error: "Missing required fields", data: null },
        { status: 400 }
      );
    }

    const langCode = getLanguage() || "en-US";

    const response = await fetch(`${API_BASE}/api/v1/alerts`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-Device-ID": deviceId,
        "Accept-Language": langCode,
      },
      body: JSON.stringify({ productId, deviceId, currentPriceETB }),
    });

    const data = await response.json();

    if (!response.ok) {
      return NextResponse.json(
        {
          error: data?.status || "Failed to create alert",
          data: null,
        },
        { status: response.status }
      );
    }

    console.log("Alert created successfully1:", data);

    return NextResponse.json({ data, error: null }, { status: 201 });
  } catch (error) {
    console.error("POST /api/v1/alerts error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}
