//src/app/api/alerts/route.ts
import { getLanguage } from "@/lib/redux/languageBridge";
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function POST(
  req: NextRequest
): Promise<NextResponse<AlertCreateResponse>> {
  try {
    const body = await req.json();
    const { productId, currentPriceETB } = body;

    // ✅ Take deviceId from headers instead of body
    const deviceId = req.headers.get("x-device-id");

    if (!productId || !deviceId || !currentPriceETB) {
      return NextResponse.json(
        { error: "Missing required fields", data: null },
        { status: 400 }
      );
    }


    const langCode = getLanguage() || "en-US";

    const response = await fetch(`${API_BASE}/api/alerts`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-Device-ID": deviceId,
        "Accept-Language": langCode, // TODO: make dynamic later
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

    return NextResponse.json({ data, error: null }, { status: 201 });
  } catch (error) {
    console.error("POST /api/alerts error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}

// MOCK Implementation of Alerts API

// //src/app/api/alerts/[alertId]/route.ts
// import { NextResponse } from "next/server";

// let alerts: string[] = [];

// export async function POST(req: Request) {
//   const { productId } = await req.json();
//   if (!productId) {
//     return NextResponse.json(
//       { error: "productId is required" },
//       { status: 400 }
//     );
//   }

//   alerts.push(productId);
//   return NextResponse.json(
//     { data: { status: "Alert created successfully", alertId: productId } },
//     { status: 201 }
//   );
// }
