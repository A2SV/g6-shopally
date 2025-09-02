//src/app/api/alerts/[alertId]/route.ts
import { NextResponse } from "next/server";

let alerts: string[] = [];

export async function DELETE(
  req: Request,
  { params }: { params: { alertId: string } }
) {
  const { alertId } = params;
  if (!alerts.includes(alertId)) {
    return NextResponse.json({ error: "Not Found" }, { status: 404 });
  }

  alerts = alerts.filter((id) => id !== alertId);
  return NextResponse.json({
    data: { status: "Alert deleted successfully" },
  });
}
