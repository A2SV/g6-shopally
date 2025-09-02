//src/app/api/alerts/[alertId]/route.ts
import { NextResponse } from "next/server";

let alerts: string[] = [];

export async function POST(req: Request) {
  const { productId } = await req.json();
  if (!productId) {
    return NextResponse.json(
      { error: "productId is required" },
      { status: 400 }
    );
  }

  alerts.push(productId);
  return NextResponse.json(
    { data: { status: "Alert created successfully", alertId: productId } },
    { status: 201 }
  );
}
