//src/app/api/v1/alerts/[alertId]/route.ts
import { AlertCreateResponse } from "@/types/SavedItems/AlertCreateResponse";
import { getDeviceIdServer } from "@/utils/deviceId.server";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const API_BASE = process.env.API_BASE;

export async function DELETE(
  req: NextRequest,
  context: { params: Promise<{ alertId: string }> }
): Promise<NextResponse<AlertCreateResponse>> {
  const { alertId } = await context.params;

  try {
    // Get deviceId and Language from cookies
    const cookieStore = await cookies();
    const deviceId = (await getDeviceIdServer()) ?? "";
    const langCode = cookieStore.get("lang")?.value || "en";

    if (!deviceId) {
      return NextResponse.json(
        { error: "Device ID is required", data: null },
        { status: 400 }
      );
    }

    const backendRes = await fetch(`${API_BASE}/api/v1/alerts/${alertId}`, {
      method: "DELETE",
      headers: {
        "X-Device-ID": deviceId,
        "Accept-Language": langCode, // TODO: make dynamic later
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

    console.log("Alert deleted successfully:", data);

    return NextResponse.json({ data, error: null }, { status: 200 });
  } catch (error) {
    console.error("DELETE /api/v1/alerts/[id] error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}
