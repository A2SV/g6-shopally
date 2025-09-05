// utils/deviceId.ts

export function getOrCreateDeviceId() {
  if (typeof window !== "undefined") {
    let deviceId = localStorage.getItem("deviceId");
    if (!deviceId) {
      deviceId = crypto.randomUUID(); // built-in
      localStorage.setItem("deviceId", deviceId);
    }
    return deviceId;
  }
  return "server-device";
}
