// middleware.ts
import { withAuth } from "next-auth/middleware";
import { NextResponse } from "next/server";

export default withAuth(
  function middleware(req) {
    console.log("middleware called on:", req.nextUrl.pathname);

    const res = NextResponse.next();

    // ✅ Check if deviceId cookie exists
    const deviceId = req.cookies.get("deviceId")?.value;
    if (!deviceId) {
      const newId = crypto.randomUUID();
      res.cookies.set("deviceId", newId, {
        path: "/",
        maxAge: 60 * 60 * 24 * 365,
      });
      console.log("Generated new deviceId:", newId);
    } else {
      console.log("Existing deviceId:", deviceId);
    }

    return res;
  },
  {
    callbacks: {
      authorized: ({ token }) => !!token,
    },
  }
);

export const config = {
  matcher: [
    "/",
    "/home",
    "/comparison",
    "/saved-items",
    "/profile",
    "/how-it-works",
  ],
};
