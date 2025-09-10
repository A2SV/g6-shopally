// src/app/api/v1/compare/route.ts
import { ComparisonResponse } from "@/types/Compare/Comparison";
import { Product } from "@/types/types";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

// ✅ Helper for an empty comparison response
const emptyComparisonResponse: ComparisonResponse = {
  data: {
    products: [],
    overallComparison: {
      bestValueProduct: "",
      bestValueLink: "",
      bestValuePrice: { etb: 0, usd: 0, fxTimestamp: new Date().toISOString() },
      keyHighlights: [],
      summary: "",
    },
  },
  error: null,
};

export async function POST(
  req: NextRequest
): Promise<NextResponse<ComparisonResponse>> {
  try {
    const body = await req.json();
    const { products } = body;

    // Get deviceId and Language from cookies
    const cookieStore = await cookies();
    const deviceId = cookieStore.get("deviceId")?.value;
    const langCode = cookieStore.get("lang")?.value || "en";

    if (
      !products ||
      !Array.isArray(products) ||
      products.length < 2 ||
      products.length > 4
    ) {
      return NextResponse.json(
        {
          error: "Products array must have 2 to 4 items",
          data: emptyComparisonResponse.data,
        },
        { status: 400 }
      );
    }

    if (!deviceId) {
      return NextResponse.json(
        { error: "Device ID is required", data: emptyComparisonResponse.data },
        { status: 400 }
      );
    }

    const backendRes = await fetch(`${API_BASE}/api/v1/compare`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-Device-ID": deviceId,
        "Accept-Language": langCode,
      },
      body: JSON.stringify({
        products: products.map((p: Product) => ({
          id: p.id,
          title: p.title,
          imageUrl: p.imageUrl,
          aiMatchPercentage: p.aiMatchPercentage,
          price: p.price,
          productRating: p.productRating,
          sellerScore: p.sellerScore,
          deliveryEstimate: p.deliveryEstimate,
          summaryBullets: p.summaryBullets,
          deeplinkUrl: p.deeplinkUrl,
        })),
      }),
    });

    const data = await backendRes.json();

    // Log raw backend response
    console.log("🔍 Backend compare API response:", {
      status: backendRes.status,
      ok: backendRes.ok,
      data,
    });

    if (!backendRes.ok) {
      return NextResponse.json(
        {
          error: data?.error || "Failed to compare products",
          data: emptyComparisonResponse.data,
        },
        { status: backendRes.status }
      );
    }

    // Pass backend response directly, ensuring type matches
    const response: ComparisonResponse = {
      data: data?.data ?? emptyComparisonResponse.data,
      error: data?.error ?? null,
    };

    return NextResponse.json(response, { status: 200 });
  } catch (error) {
    console.error("POST /api/v1/compare error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: emptyComparisonResponse.data },
      { status: 500 }
    );
  }
}
