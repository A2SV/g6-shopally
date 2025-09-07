// middleware.ts
import { withAuth } from "next-auth/middleware";
import { NextResponse } from "next/server";

export default withAuth(
  function middleware(req) {
    console.log("middleware called on:", req.nextUrl.pathname);
    return NextResponse.next();
  },
  {
    callbacks: {
      authorized: ({ token }) => !!token,
    },
  }
);

export const config = {
  matcher: [
    // "/((?!api/auth|_next/static|_next/image|favicon.ico|SignUp|SignIn|data|profile|comparison|saved-items|home|WebsiteLogo/Frame.png).*)",
    "/((?!api/auth|_next/static|_next/image|favicon.ico|images|WebsiteLogo).*)",
    // Only protect specific existing routes
    // "/",
    // "/home",
    // "/comparison", 
    // "/saved-items",
    // "/profile",
    // "/how-it-works",
  ],
};
