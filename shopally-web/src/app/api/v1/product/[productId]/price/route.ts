//src/app/api/v1/product/[productId]/price/route.ts

console.log("✅ Product price route loaded");

import { UpdateProductResponse } from "@/types/SavedItems/SavedItems";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE!;

export async function GET(
  req: NextRequest,
  context: { params: Promise<{ productId: string }> }
): Promise<NextResponse<UpdateProductResponse>> {
  const { productId } = await context.params;

  try {
    const cookieStore = await cookies();
    const deviceId = cookieStore.get("deviceId")?.value;
    const langCode = cookieStore.get("lang")?.value || "en";

    if (!deviceId) {
      return NextResponse.json(
        {
          error: {
            error:
              "Device ID not available yet. Please wait for notifications setup.",
          },
          data: null,
        },
        { status: 409 }
      );
    }

    const backendRes = await fetch(
      `${API_BASE}/api/v1/product/${productId}/price`,
      {
        method: "GET",
        headers: {
          "X-Device-ID": deviceId,
          "Accept-Language": langCode,
          "Content-Type": "application/json",
        },
      }
    );

    const data = await backendRes.json();

    if (!backendRes.ok) {
      return NextResponse.json(
        {
          error: { error: data?.status || "Failed to update price" },
          data: null,
        },
        { status: backendRes.status }
      );
    }

    console.log("Price updated successfully:", data);

    return NextResponse.json(
      { data: data.data, error: { error: null } }, // unwrap properly
      { status: 200 }
    );
  } catch (error) {
    console.error("GET /api/v1/product/[productId]/price error:", error);
    return NextResponse.json(
      { error: { error: "Something went wrong" }, data: null },
      { status: 500 }
    );
  }
}
