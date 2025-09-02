//src/app/api/users/[alertId]/route.ts
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function DELETE(
  req: NextRequest,
  { params }: { params: { id: string } }
): Promise<NextResponse<AlertCreateResponse>> {
  const { id } = params;

  try {
    // ✅ Get deviceId from headers
    const deviceId = req.headers.get("x-device-id");

    if (!deviceId) {
      return NextResponse.json(
        { error: "Device ID is required", data: null },
        { status: 400 }
      );
    }

    const backendRes = await fetch(`${API_BASE}/api/alerts/${id}`, {
      method: "DELETE",
      headers: {
        "X-Device-ID": deviceId,
        "Accept-Language": "en-US", // TODO: make dynamic later
        "Content-Type": "application/json",
      },
    });

    const data = await backendRes.json();

    if (!backendRes.ok) {
      return NextResponse.json(
        {
          error: data?.status || "Failed to delete Alert",
          data: null,
        },
        { status: backendRes.status }
      );
    }

    return NextResponse.json({ data, error: null }, { status: 200 });
  } catch (error) {
    console.error("DELETE /api/alerts/[id] error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}
