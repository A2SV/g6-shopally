//app/api/auth/[...nextauth]/options.ts
import type { NextAuthOptions } from "next-auth";
import GoogleProvider from "next-auth/providers/google";

export const options: NextAuthOptions = {
  providers: [
    GoogleProvider({
      profile(profile) {
        return {
          id: profile.sub, // NextAuth requires an `id`
          name: profile.name,
          email: profile.email,
          image: profile.picture,
        };
      },
      clientId: process.env.GOOGLE_ID!,
      clientSecret: process.env.GOOGLE_SECRET!,
    }),
  ],

  callbacks: {
    async signIn() {
      return true;
    },
  },
};
