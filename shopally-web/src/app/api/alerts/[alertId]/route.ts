//src/app/api/alerts/[alertId]/route.ts
import { getLanguage } from "@/lib/redux/languageBridge";
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

    const langCode = getLanguage() || "en-US";

    const backendRes = await fetch(`${API_BASE}/api/alerts/${id}`, {
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

    return NextResponse.json({ data, error: null }, { status: 200 });
  } catch (error) {
    console.error("DELETE /api/alerts/[id] error:", error);
    return NextResponse.json(
      { error: "Something went wrong", data: null },
      { status: 500 }
    );
  }
}

//MOCK Implementation of Alerts API

// //src/app/api/alerts/[alertId]/route.ts
// import { NextResponse } from "next/server";

// let alerts: string[] = [];

// export async function DELETE(
//   req: Request,
//   { params }: { params: { alertId: string } }
// ) {
//   const { alertId } = params;
//   if (!alerts.includes(alertId)) {
//     return NextResponse.json({ error: "Not Found" }, { status: 404 });
//   }

//   alerts = alerts.filter((id) => id !== alertId);
//   return NextResponse.json({
//     data: { status: "Alert deleted successfully" },
//   });
// }
