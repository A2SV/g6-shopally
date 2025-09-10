// utils/deviceId.server.ts
import { cookies } from "next/headers";

export async function getDeviceIdServer(): Promise<string | null> {
  const cookieStore = await cookies();
  return cookieStore.get("deviceId")?.value || null;
}
