import { ComparisonResponse } from "@/types/Compare/Comparison";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function POST(
  req: NextRequest,
): Promise<NextResponse<ComparisonResponse>> {
  try {
    const body = await req.json();
    const { products } = body;

    // ✅ Get deviceId and language from headers
    const deviceId = req.headers.get("x-device-id");
    const acceptLanguage = req.headers.get("accept-language") || "en-US";

    if (
      !products ||
      !Array.isArray(products) ||
      products.length < 2 ||
      products.length > 4
    ) {
      return NextResponse.json(
        {
          error: "Products array must have 2 to 4 items",
          data: { comparison: [] },
        },
        { status: 400 },
      );
    }

    if (!deviceId) {
      return NextResponse.json(
        { error: "Device ID is required", data: { comparison: [] } },
        { status: 400 },
      );
    }

    const backendRes = await fetch(`${API_BASE}/api/compare`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-Device-ID": deviceId,
        "Accept-Language": acceptLanguage,
      },
      body: JSON.stringify({ products }),
    });

    const data = await backendRes.json();

    if (!backendRes.ok) {
      return NextResponse.json(
        {
          error: data?.status || "Failed to compare products",
          data: { comparison: [] },
        },
        { status: backendRes.status },
      );
    }

    return NextResponse.json({ data, error: null }, { status: 200 });
  } catch (error) {
    console.error("POST /api/compare error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: { comparison: [] } },
      { status: 500 },
    );
  }
}
