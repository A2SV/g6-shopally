//src/utils/deviceId.client.ts
import { app } from "@/lib/firebase"; // your Firebase init
import { getMessaging, getToken } from "firebase/messaging";

const VAPID_KEY = process.env.NEXT_PUBLIC_FIREBASE_VAPID_KEY!;

export async function getDeviceIdClient(): Promise<string | null> {
  if (typeof window === "undefined") return null;

  // Check if already saved in localStorage
  const existing = localStorage.getItem("deviceId");
  if (existing) return existing;

  try {
    const messaging = getMessaging(app);
    const token = await getToken(messaging, { vapidKey: VAPID_KEY });
    if (token) {
      localStorage.setItem("deviceId", token);
      return token;
    }
    return null;
  } catch (err) {
    console.error("Failed to get FCM token:", err);
    return null;
  }
}
