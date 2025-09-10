//src/app/api/v1/alerts/route.ts
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function POST(
  req: NextRequest
): Promise<NextResponse<AlertCreateResponse>> {
  try {
    const body = await req.json();
    const { productId, currentPriceETB } = body;
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

    if (!productId || !currentPriceETB) {
      return NextResponse.json(
        { error: "Missing required fields", data: null },
        { status: 400 }
      );
    }

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
